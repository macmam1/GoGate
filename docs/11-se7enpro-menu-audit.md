# 11) Se7en Pro Menu Audit (from provided screenshots)

## مشاهده‌شده در منوها
- Sidebar: Home, Settings, IP Scanner, Logs, About
- Settings > Appearance: Theme selector
- Settings > Connection:
  - Egress Region (Auto/best available)
  - Set Windows system proxy automatically when connected
  - Disable timeouts (for unstable networks)
- Settings > Local Proxy Ports:
  - SOCKS5 port
  - HTTP port
  - Allow connections from LAN
  - LAN client auth (username/password optional)
- Settings > Upstream Proxy:
  - Type: HTTP / SOCKS5(h) / SOCKS4a
  - Host/IP, Port, optional credentials
- Settings > App Behavior:
  - Auto-connect on startup
  - Start with Windows
  - Window close behavior
- Settings > Advanced Tunneling:
  - Protocol mode: Auto / Direct / CDN Fronting
  - Beast Mode (parallel aggressive strategy)
  - Auto-find IP & SNI toggle
  - Save found IPs/SNIs
  - Custom Edge IPs
  - SNI overrides

## ویژگی‌هایی که باید در GoGate تضمین شوند
1. parity کامل منوی Settings بالا (با UX بهتر)
2. IP Scanner با حالت Quick/Deep + verify چندمرحله‌ای
3. Logs ساختاریافته + export diagnostics
4. About با برند MRH-DevLoop + نسخه/بیلد + hash

## پیاده‌سازی ایمن/قانونی برای بخش حساس
- به‌جای منطق تهاجمی:
  - Restricted Network Mode (opt-in)
  - CDN Route Optimization (تنظیم مسیرهای مجاز)
  - Smart Fallback Profiles
  - Explicit consent + kill-switch + audit trail

## موارد افزوده رقابتی
- Explainable routing (چرا این مسیر انتخاب شد)
- Theme performance profiles: Lite/Balanced/Rich
- Reduced Motion mode
- Connection recovery score
