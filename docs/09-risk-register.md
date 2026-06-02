# 09) Risk Register

| ریسک | شدت | احتمال | اقدام کاهشی |
|---|---|---|---|
| ناپایداری شبکه/فیلترینگ پویا | بالا | بالا | fallback چندمرحله‌ای + re-test هوشمند |
| false positive در اسکن/رتبه‌بندی | متوسط | بالا | verify چندمرحله‌ای + history-aware scoring |
| هشدار اعتبار فایل در ویندوز | متوسط | بالا | انتشار شفاف + checksum + برنامه code-signing |
| پیچیدگی UI انیمیشنی | متوسط | متوسط | بودجه performance + reduced-motion |
| وابستگی به پروژه‌های ثالث | بالا | متوسط | adapter abstraction + pin version + mirror policy |
| بدهی فنی از ادغام زیاد | بالا | متوسط | phase gating + ADR اجباری |
| ریسک حقوقی در قابلیت‌های حساس | بالا | متوسط | opt-in صریح + توضیح ریسک + غیرفعال پیش‌فرض |
