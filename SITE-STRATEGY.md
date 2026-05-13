# AwardWinner / ModernTimez — Site Strategy & Architecture

> **Analysis Date:** 2026-05-13  
> **Platform:** Shopify (Dawn theme base)  
> **GitHub:** https://github.com/jin801/AwardWinner.git (branch: main — clean, up to date)  
> **Brand Name:** ModernTimez Gift Shop  
> **Specialty:** Custom awards, personalized gifts, engraving, recognition products

---

## 1. Brand Positioning

ModernTimez occupies the premium, personalized-gifts niche — not a general gift shop. Every purchase decision should be guided by four brand pillars:

| Pillar | What It Means |
|---|---|
| **Craftsmanship** | Engraved trophies, crystal awards, plaques — built to last |
| **Personalization** | Every product feels made for the recipient, not mass-produced |
| **Occasion-Led** | Categories built around real milestones: baptism → real estate closing → retirement |
| **Trust** | Recognition expertise since 2006; bulk ordering for orgs |

**Primary audiences:**
1. **Corporate buyers** — HR, operations, event planners buying awards, plaques, recognition gifts in bulk
2. **Real estate agents** — closing gifts, referral gifts, new-homeowner keepsakes for clients
3. **Life-event shoppers** — baptism, graduation, wedding, anniversary
4. **Organizations** — churches, schools, non-profits needing recognition items

---

## 2. Collection Architecture (Shopify Collections Needed)

### Occasion-Based Collections
| Handle | Display Name | Priority |
|---|---|---|
| `baptism-gifts` | Baptism & Christening Gifts | High |
| `wedding-gifts` | Wedding Gifts | High |
| `graduation-gifts` | Graduation Gifts | High |
| `fathers-day-gifts` | Father's Day Gifts | Seasonal |
| `real-estate-gifts` | Real Estate Client Gifts | High |
| `anniversary-gifts` | Anniversary Gifts | Medium |
| `retirement-gifts` | Retirement Gifts | Medium |
| `new-home-gifts` | New Home Gifts | Medium |

### Product-Type Collections
| Handle | Display Name | Priority |
|---|---|---|
| `corporate-awards` | Corporate Awards & Trophies | High |
| `crystal-awards` | Crystal Awards | High |
| `engraved-plaques` | Engraved Plaques | High |
| `personalized-gifts` | Personalized Gifts | High |
| `trophies` | Trophies | Medium |
| `gift-cards` | Gift Cards | Always |

### Audience Collections
| Handle | Display Name | Priority |
|---|---|---|
| `for-him` | Gifts for Him | Medium |
| `for-her` | Gifts for Her | Medium |
| `for-teams` | Team & Group Gifts | Medium |
| `bulk-orders` | Bulk & Wholesale | High |

---

## 3. Page Architecture (All Pages)

### Core Pages (Live / Built)
- `/` — Homepage (awards hero + occasion grid + featured products + testimonials)
- `/collections/all` — All Products
- `/pages/about` — About Us ✅
- `/pages/contact` — Contact / Quote Request ✅
- `/pages/faq` — FAQ ✅
- `/pages/gift-wrapping` — Gift Wrapping Service ✅
- `/pages/returns` — Returns Policy ✅
- `/pages/shipping` — Shipping Policy ✅

### New Pages Required
- `/pages/corporate` — Corporate Gifting Program ✅ (added)
- `/pages/real-estate` — Real Estate Agent Client Gifts ✅ (added)
- `/pages/personalization` — How Personalization Works ✅ (added)
- `/pages/bulk-quote` — Bulk / Wholesale Quote Form (use contact page for now)

---

## 4. Navigation Structure (Mega Menu)

```
HEADER MENU
├── Shop Awards
│   ├── Crystal Awards
│   ├── Engraved Plaques
│   ├── Trophies
│   └── Personalized Gifts
├── By Occasion
│   ├── Baptism & Christening
│   ├── Wedding
│   ├── Graduation
│   ├── Father's Day
│   ├── Anniversary
│   └── Retirement
├── Corporate & Business
│   ├── Corporate Awards
│   ├── Real Estate Gifts
│   ├── Bulk Orders
│   └── Request a Quote
├── Gift Ideas          (Blog / gift guide hub)
└── About
    ├── Our Story
    ├── How Personalization Works
    └── Contact Us
```

### Footer Menu Structure
```
SHOP                HELP              COMPANY           FOLLOW US
Crystal Awards      FAQ               About Us          Facebook
Plaques             Shipping          Our Promise       Instagram
Trophies            Returns           Contact           TikTok
Personalized        Gift Wrapping     Corporate Gifts
Real Estate Gifts   Track Order       Real Estate Gifts
Gift Cards          Privacy Policy
```

---

## 5. Homepage Section Order (Optimized)

| # | Section | Purpose | Color Scheme |
|---|---|---|---|
| 1 | Announcement Bar | Promotions / shipping threshold | scheme-4 (gold on dark) |
| 2 | Awards Hero | Hero with CTA | scheme-3 (dark mahogany) |
| 3 | Trust Badges | Free engraving / Fast ship / 30-day returns | scheme-2 |
| 4 | Featured Collection | Best sellers | scheme-1 |
| 5 | Shop by Occasion | 6-tile grid: Baptism / Wedding / Corporate / Real Estate / Graduation / Anniversary | scheme-2 |
| 6 | Corporate + Real Estate CTA | Split CTA for B2B audiences | scheme-3 |
| 7 | Testimonials | Social proof | scheme-1 |
| 8 | Newsletter | Email capture | scheme-5 (gold) |
| 9 | Gift Advice CTA | Personal shopping service | scheme-3 |

---

## 6. Product Page Structure (Optimized)

```
main-product
  ├── vendor
  ├── title
  ├── price + rating
  ├── variant picker
  ├── quantity selector
  ├── buy buttons (dynamic checkout)
  ├── description
  ├── [tab] Personalization Guide
  ├── [tab] Gift Wrapping Available
  ├── [tab] Shipping & Delivery
  ├── [tab] Returns & Exchanges
  └── [tab] Why This Makes a Great Gift

gift-promise (3-col icons)
  ├── Free Engraving
  ├── Gift Wrapping
  └── 30-Day Returns

related-products ("Complete the Gift")
```

---

## 7. SEO Structure

### Page Title Templates
- Homepage: `Custom Awards & Personalized Gifts | ModernTimez`
- Collection: `[Collection Name] | ModernTimez Awards & Gifts`
- Product: `[Product Name] — Engraved & Personalized | ModernTimez`

### Priority Keywords
| Keyword Cluster | Target Page |
|---|---|
| Custom awards, engraved trophies, crystal awards | Homepage, /collections/corporate-awards |
| Corporate recognition gifts, employee awards | /pages/corporate |
| Real estate closing gifts, realtor client gifts | /pages/real-estate |
| Baptism gifts personalized, christening keepsakes | /collections/baptism-gifts |
| Personalized wedding gifts, engraved wedding | /collections/wedding-gifts |
| Bulk trophy orders, wholesale awards | /pages/corporate |

### Schema.org (Status)
- Organization ✅
- Product ✅
- BlogPosting ✅
- WebSite with SearchAction ✅
- FAQPage — add to /pages/faq (future improvement)

---

## 8. GitHub Readiness Checklist

| Check | Status | Notes |
|---|---|---|
| Branch synced with origin/main | PASS | Up to date |
| No secrets in tracked files | PASS | No API keys or tokens found |
| .gitignore present | PASS | |
| Layout files complete | PASS | theme.liquid, password.liquid |
| Core sections present | PASS | All Dawn sections + custom awards-hero |
| Customer account templates | PASS | All 7 pages |
| Schema.org structured data | PASS | 4 types implemented |
| Settings data configured | PASS | Colors, fonts, social links |
| .claude/ uncommitted | PASS | settings.local.json stays local |
| Real estate page template | FIXED | Added this session |
| Corporate page template | FIXED | Added this session |
| Personalization page | FIXED | Added this session |
| Real estate in occasion grid | FIXED | Added this session |
| About page brand positioning | FIXED | Updated this session |
| FAQ personalization section | FIXED | Added this session |

### Git Commands to Push After Review
```bash
git add templates/page.corporate.json
git add templates/page.real-estate.json
git add templates/page.personalization.json
git add templates/index.json
git add templates/page.about.json
git add templates/page.faq.json
git add SITE-STRATEGY.md
git commit -m "feat: corporate, real-estate, personalization pages + occasion grid + brand positioning"
git push origin main
```

---

## 9. Shopify Admin Action Items (Cannot Be Done Via Code)

### Collections to Create
1. `baptism-gifts` — Baptism & Christening Gifts
2. `wedding-gifts` — Wedding Gifts
3. `graduation-gifts` — Graduation Gifts
4. `real-estate-gifts` — Real Estate Client Gifts
5. `corporate-awards` — Corporate Awards & Trophies
6. `crystal-awards` — Crystal Awards
7. `engraved-plaques` — Engraved Plaques
8. `personalized-gifts` — Personalized Gifts
9. `fathers-day-gifts` — Father's Day Gifts

### Pages to Create (after pushing templates)
1. `/pages/corporate` — assign template: `page.corporate`
2. `/pages/real-estate` — assign template: `page.real-estate`
3. `/pages/personalization` — assign template: `page.personalization`

### Navigation
1. Main menu — restructure with mega menu (3 columns)
2. Footer menu — 4-column layout

### Announcement Bar Rotation
1. "Free personalization on orders over $75"
2. "Bulk orders for corporate & events — Request a quote"
3. "Ships in 1–2 business days"

### Theme Settings
- Enable mega menu for desktop
- Upload logo.svg
- Add Instagram + TikTok social links (currently only Facebook)

---

## 10. Launch Phasing

### Phase 1 — Go-Live Ready
- [x] Homepage optimized (6-occasion grid, updated CTAs)
- [x] Corporate gifting page
- [x] Real estate agent page
- [x] Personalization guide page
- [x] About page repositioned
- [x] FAQ with personalization section
- [ ] Create collections in Shopify admin
- [ ] Upload product photos (minimum 10 hero products)
- [ ] Set up payment gateway

### Phase 2 — Growth (Month 1)
- [ ] Blog: "Best Real Estate Closing Gifts" (SEO)
- [ ] Blog: "Corporate Award Ideas for Every Budget"
- [ ] Blog: "How to Choose a Baptism Gift"
- [ ] Google Merchant Center product feed
- [ ] Meta/Instagram product catalog
- [ ] Customer review app (Judge.me or Okendo)

### Phase 3 — Scale (Month 2–3)
- [ ] Loyalty/rewards program (Smile.io)
- [ ] Personalization configurator app (Kickflip or Infinite Options)
- [ ] Wholesale portal for real estate agents
- [ ] Email flows: welcome, abandoned cart, post-purchase
- [ ] Live chat (Tidio or Gorgias)

---

## 11. Design System

| Token | Value | Use |
|---|---|---|
| `#FAF7F0` | Warm cream | Primary background |
| `#2C1A0E` | Dark espresso | Body text |
| `#9A7B2F` | Antique gold | CTA buttons, accents |
| `#3D1F0D` | Deep mahogany | Dark hero sections |
| `#D4B86A` | Soft gold | Announcement bar, callouts |
| `#F5F0E8` | Ivory | Secondary backgrounds |
| Playfair Display | Serif | All headings |
| DM Sans | Sans-serif | Body, UI text |

---

*Generated 2026-05-13. Update as the store evolves.*
