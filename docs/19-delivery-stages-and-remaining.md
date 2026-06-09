# 19) Delivery Stages & Remaining Work

## Current snapshot (as of 2026-06-02)
Completed:
1. Repository bootstrap, governance, CI scaffolding
2. Bilingual README + donation setup
3. PRD/ADR/roadmap artifacts
4. Core runtime skeleton (contracts, orchestrator, scoring, harvester, adapter boundary)

## Remaining stages (high-level)
### Stage 5 — Engine Runtime Integration (completed baseline)
- [x] xray process adapter wiring (initial)
- [x] sing-box process adapter wiring (initial)
- [x] capability negotiation and base error mapping (initial)
- [x] runtime stderr semantic parsing
- [x] integration test fixtures for adapter lifecycle

### Stage 6 — Harvester Connectors (completed baseline)
- [x] source allowlist manager
- [x] robust Git connector handling (rate-limit/retries) baseline
- [x] subscription ingestion scheduler baseline
- [x] persistence hooks for scheduler state storage integration (baseline)
- [x] ingestion metrics counters/events (baseline)
- [x] backoff tuning from runtime policy (baseline)
- [x] production telemetry wiring to app-level observability (baseline)
- [x] persistent scheduler state store adapter (baseline)

### Stage 7 — Probe + Scoring + Fallback Integration (completed)
- [x] quick/deep probe worker contracts and stub implementation
- [x] scoring aggregation helper from probe signals
- [x] fallback policy execution integration hook in orchestrator ranking flow
- [x] historical score blending hook in scoring layer
- [x] live-session fallback execution binding
- [x] persisted history blending source wiring

### Stage 8 — Windows App MVP (UI + Core Bridge) (in progress)
- [x] shell/navigation/state scaffold specs
- [x] shell -> core bridge API contract draft
- [x] shell executable scaffold implementation (baseline)
- [x] state sync hook (baseline via bridge subscription)
- [x] state sync with real orchestrator bridge transport (baseline local-rpc client)
- [x] theme profiles runtime binding (baseline)
- [x] bridge reconnect/error hardening (baseline)
- [x] shell navigation host wiring (baseline)
- [x] shell navigation host -> page binding layer wiring (baseline)
- [x] bridge health indicator component wiring (baseline)
- [ ] concrete UI page/view wiring
- [ ] concrete bridge-health status-card widget wiring

### Stage 9 — Android App MVP (in progress)
- [x] shell/navigation/state scaffold specs
- [x] shell executable scaffold implementation (baseline)
- [x] state sync hook (baseline via bridge subscription)
- [x] local-rpc bridge client baseline
- [x] theme runtime profile binding baseline
- [x] bridge reconnect/error hardening (baseline)
- [x] shell navigation host wiring (baseline)
- [x] profile selection/connect flows (baseline shell wiring)
- [x] session status and log view access (baseline shell wiring)
- [x] shared configuration compatibility (baseline shared schema/parser)
- [x] UI page binding layer for profile + logs screens (baseline)
- [ ] concrete Android page/view wiring

### Stage 10 — Packaging & Release Hardening
- installer pipeline
- build metadata + checksums
- release validation checklist

### Stage 11 — Public Beta Readiness
- docs polish (EN/FA)
- diagnostics and support workflow
- bug triage and stabilization

## Estimated remaining stage count
**4 major stages remain** (Stages 8 to 11).
