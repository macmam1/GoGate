# 04) GitHub & Release Operations

## آیا می‌توانم 100% فرایند ساخت تا انتشار را انجام دهم؟
**تقریباً بله**؛ با دسترسی مناسب می‌توانم از طراحی تا ساخت ریپو، CI/CD و انتشار Release را انجام دهم.

## کارهایی که حتماً شما باید انجام بدهی
1. ساخت/مالکیت حساب GitHub و فعال‌بودن 2FA
2. ساخت Fine-grained PAT با حداقل دسترسی لازم
3. تصمیم درباره Public/Private و نام نهایی ریپو
4. برای امضای ویندوز (اگر خواستی): تهیه گواهی و دسترسی امن به secret

## کارهایی که من اتوماتیک انجام می‌دهم
- ساخت ساختار monorepo
- تنظیم استانداردهای repo (templates, labels, CODEOWNERS, SECURITY)
- CI pipeline برای build/test/release
- تولید artifact + checksum + release notes
- مدیریت نسخه‌ها و changelog

## Release Channels
- `nightly` (internal)
- `beta` (community)
- `stable` (public)

## خروجی هر Release
- Windows installer + portable
- Android APK (فاز اول خارج Play)
- Checksums (SHA256)
- راهنمای نصب/رفع خطا
