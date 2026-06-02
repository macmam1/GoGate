# 01) Feature Dependency Map (ریز وابستگی‌ها)

| Feature | پیش‌نیاز | خروجی مستقیم | وابستگی بعدی |
|---|---|---|---|
| Config Import/Parser | فرمت‌شناسی، validation اولیه | کانفیگ نرمال | تست سلامت |
| Subscription Manager | scheduler، diff | بروزرسانی خودکار | re-score خودکار |
| Health Probe (Quick) | TCP/TLS probe | alive/dead + RTT | رتبه‌بندی |
| Health Probe (Deep) | quick probe + timeout policy | jitter/loss/handshake score | انتخاب پروفایل |
| IP/Domain Candidate Scan | policy + throttle + allowlist | لیست candidate | verify + scoring |
| Scoring Engine | دیتای probe + تاریخچه | امتیاز کیفیت | Smart Connect |
| Geo/ISP Classification | geoip/asn db | دسته‌بندی منطقه‌ای | UI سرورها |
| Smart Connect | scoring + policy | مسیر پیشنهادی | اتصال اصلی |
| Fallback Chain | adapter چند engine | failover خودکار | پایداری جلسه |
| Engine Adapters | binary lifecycle mgmt | اتصال واقعی | telemetry |
| Session Telemetry | event schema | گزارش وضعیت | debug/export |
| Windows Installer | packaging + dependency bootstrap | نصب یکپارچه | auto-update |
| Android Packaging | VpnService wiring | APK release | staged rollout |
| User Consent Gate (Sensitive Modes) | legal copy + explicit opt-in | فعال‌سازی کنترل‌شده | audit trail |

## زنجیره منطقی اسکنر (مطابق نیاز شما)
`Scan -> Validate -> Verify -> Score -> Attach to Profile -> Connect Test -> Rank -> Show in UI`

> نتیجه: اسکنر یک قابلیت مستقل نیست؛ یک موتور تغذیه برای افزایش نرخ موفقیت در اتصال است.
