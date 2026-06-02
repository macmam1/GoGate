# 10) GitHub Bootstrap Checklist (Automation-First)

## هدف
با کمترین کار دستی شما، بیشترین اتوماسیون برای ساخت و انتشار GoGate انجام شود.

## مرحله A — یک‌بار توسط شما (5 دقیقه)
1. وارد GitHub شوید: https://github.com/settings/tokens?type=beta
2. Fine-grained token بسازید:
   - Resource owner: `macmam1`
   - Repository access: Only selected repositories -> `GoGate` (اگر هنوز نیست، بعد از ساخت repo انتخاب کن)
   - Permissions:
     - Contents: Read and Write
     - Metadata: Read-only
     - Pull requests: Read and Write
     - Actions: Read and Write
     - Workflows: Read and Write
     - Issues: Read and Write (اختیاری)
3. Expiration: 30 days
4. Token را نگه دار (فعلاً share عمومی نکن)

## مرحله B — ساخت repo
- یا خودت بساز: `https://github.com/new` با نام `GoGate`
- یا به من token بده تا با CLI بسازم و تنظیمات کامل کنم.

## مرحله C — وقتی آماده بودی
من این کارها را اتوماتیک انجام می‌دهم:
- bootstrap monorepo
- branch protections
- issue/pr templates
- labels/milestones
- GitHub Actions pipelines
- release workflow + checksum

## مرحله D — بعد از تحویل نسخه اولیه
- token rotate/revoke
- تنظیم token جدید برای عملیات روزمره
