# 02) High-Level Architecture

## Monorepo Layout (پیشنهادی)
- `apps/windows-client`
- `apps/android-client`
- `apps/macos-client` (phase 3)
- `apps/linux-client` (phase 3)
- `core/orchestrator`
- `core/config-pipeline`
- `core/scoring-engine`
- `core/telemetry`
- `engines/adapters/{xray,singbox,wireguard,openvpn}`
- `packages/design-system`
- `packages/animated-ui-assets`
- `infra/ci`
- `docs/`

## لایه‌ها
1. Ingestion Layer
2. Validation & Trust Layer
3. Probe & Scoring Layer
4. Orchestration & Fallback Layer
5. Execution Layer (Engine Adapters)
6. UX + Observability

## اصل مهم مهندسی
- **No core rewrite**: هسته‌های متن‌باز استاندارد استفاده می‌شوند.
- Adapter abstraction برای جلوگیری از قفل شدن روی یک engine.
- Feature flags برای ویژگی‌های Experimental.

## دیتافلو
- ورودی کانفیگ/ساب -> normalize -> health probes -> score -> profile select -> connect -> monitor -> fallback if needed
