// deploy.go — ModernTimezGift theme validation and deployment tool
// Usage: go run deploy.go [--push --store YOUR_STORE.myshopify.com]
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	assetRe  = regexp.MustCompile(`'([^']+)'\s*\|\s*asset_url`)
	renderRe = regexp.MustCompile(`render\s+'([^']+)'`)
)

type issue struct {
	file string
	msg  string
}

func main() {
	store := flag.String("store", "", "Shopify store domain (e.g. mystore.myshopify.com)")
	push := flag.Bool("push", false, "Push theme to Shopify after validation")
	flag.Parse()

	fmt.Println("┌─────────────────────────────────────────────┐")
	fmt.Println("│   ModernTimezGift · Theme Deploy Tool       │")
	fmt.Println("└─────────────────────────────────────────────┘")

	// Step 1: Validate
	fmt.Println("\n[1/3] Validating theme...")
	issues := validate()
	if len(issues) > 0 {
		fmt.Printf("      ✗ %d issue(s) found:\n\n", len(issues))
		for _, iss := range issues {
			fmt.Printf("        • %s\n          %s\n\n", iss.file, iss.msg)
		}
		fmt.Println("      Fix issues before deploying.")
		os.Exit(1)
	}
	fmt.Println("      ✓ All asset, snippet, and section references OK")

	// Step 2: Git
	fmt.Println("\n[2/3] Git commit...")
	gitCommit()

	// Step 3: Shopify push
	fmt.Println("\n[3/3] Shopify deploy...")
	if *push {
		if *store == "" {
			fmt.Println("      ✗ --store is required with --push")
			fmt.Println("        Example: go run deploy.go --push --store mystore.myshopify.com")
			os.Exit(1)
		}
		shopifyPush(*store)
	} else {
		fmt.Println("      Run the following to publish to Shopify:")
		fmt.Println()
		fmt.Println("        shopify theme push --store YOUR_STORE.myshopify.com")
		fmt.Println()
		fmt.Println("      Or use this tool with flags:")
		fmt.Println()
		fmt.Println("        go run deploy.go --push --store YOUR_STORE.myshopify.com")
		fmt.Println()
		fmt.Println("      To list existing themes first:")
		fmt.Println()
		fmt.Println("        shopify theme list --store YOUR_STORE.myshopify.com")
	}

	fmt.Println("\nDone.")
}

// validate checks all liquid and JSON files for broken references.
func validate() []issue {
	var issues []issue

	assets := fileSet("assets")
	snippets := fileSet("snippets")
	sections := fileSet("sections")

	_ = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		// Skip hidden dirs and node_modules
		for _, part := range strings.Split(path, string(os.PathSeparator)) {
			if strings.HasPrefix(part, ".") || part == "node_modules" {
				return nil
			}
		}
		switch filepath.Ext(path) {
		case ".liquid":
			issues = append(issues, checkLiquid(path, assets, snippets)...)
		case ".json":
			issues = append(issues, checkJSON(path, sections)...)
		}
		return nil
	})

	return issues
}

func checkLiquid(path string, assets, snippets map[string]bool) []issue {
	var issues []issue
	f, err := os.Open(path)
	if err != nil {
		return issues
	}
	defer f.Close()

	lineNum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineNum++
		text := scanner.Text()

		for _, m := range assetRe.FindAllStringSubmatch(text, -1) {
			name := m[1]
			if !assets[name] {
				issues = append(issues, issue{
					file: fmt.Sprintf("%s:%d", path, lineNum),
					msg:  fmt.Sprintf("missing asset: %s", name),
				})
			}
		}

		for _, m := range renderRe.FindAllStringSubmatch(text, -1) {
			name := m[1] + ".liquid"
			if !snippets[name] {
				issues = append(issues, issue{
					file: fmt.Sprintf("%s:%d", path, lineNum),
					msg:  fmt.Sprintf("missing snippet: %s", name),
				})
			}
		}
	}
	return issues
}

func checkJSON(path string, sections map[string]bool) []issue {
	var issues []issue
	data, err := os.ReadFile(path)
	if err != nil {
		return issues
	}

	var v map[string]interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return append(issues, issue{
			file: path,
			msg:  "invalid JSON: " + err.Error(),
		})
	}

	secs, ok := v["sections"].(map[string]interface{})
	if !ok {
		return issues
	}
	for key, sec := range secs {
		secMap, ok := sec.(map[string]interface{})
		if !ok {
			continue
		}
		t, ok := secMap["type"].(string)
		if !ok || t == "" || t == "_blocks" {
			continue
		}
		name := t + ".liquid"
		if !sections[name] {
			issues = append(issues, issue{
				file: path,
				msg:  fmt.Sprintf("section key '%s' references type '%s' — no sections/%s found", key, t, name),
			})
		}
	}
	return issues
}

func fileSet(dir string) map[string]bool {
	m := make(map[string]bool)
	entries, _ := os.ReadDir(dir)
	for _, e := range entries {
		if !e.IsDir() {
			m[e.Name()] = true
		}
	}
	return m
}

func isGitRepo() bool {
	_, err := os.Stat(".git")
	return err == nil
}

func gitCommit() {
	if !isGitRepo() {
		fmt.Print("      Initializing git repo... ")
		if err := exec.Command("git", "init").Run(); err != nil {
			fmt.Println("failed:", err)
			os.Exit(1)
		}
		fmt.Println("done")
	}

	if err := exec.Command("git", "add", "-A").Run(); err != nil {
		fmt.Println("      ✗ git add failed:", err)
		os.Exit(1)
	}

	out, _ := exec.Command("git", "status", "--porcelain").Output()
	if len(strings.TrimSpace(string(out))) == 0 {
		fmt.Println("      ✓ Nothing to commit — working tree clean")
		return
	}

	msg := "feat: fix broken links, enable announcement bar, add newsletter, schema.org, magnify.js fix"
	err := exec.Command("git", "commit", "-m", msg).Run()
	if err != nil {
		fmt.Println("      ✗ git commit failed:", err)
		os.Exit(1)
	}
	fmt.Println("      ✓ Changes committed")
}

func shopifyPush(store string) {
	fmt.Printf("      Pushing to %s...\n", store)
	cmd := exec.Command("shopify", "theme", "push", "--store", store)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("      ✗ shopify theme push failed:", err)
		fmt.Println("        Make sure Shopify CLI is installed: npm install -g @shopify/cli @shopify/theme")
		os.Exit(1)
	}
	fmt.Println("      ✓ Theme pushed live")
}
