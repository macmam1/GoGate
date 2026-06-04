# 19) Delivery Stages & Remaining Work

## Current snapshot (as of 2026-06-02)
Completed:
1. Repository bootstrap, governance, CI scaffolding
2. Bilingual README + donation setup
3. PRD/ADR/roadmap artifacts
4. Core runtime skeleton (contracts, orchestrator, scoring, harvester, adapter boundary)

## Remaining stages (high-level)
### Stage 5 — Engine Runtime Integration (in progress)
- [x] xray process adapter wiring (initial)
- [x] sing-box process adapter wiring (initial)
- [x] capability negotiation and base error mapping (initial)
- [ ] runtime stderr semantic parsing
- [ ] integration test fixtures for adapter lifecycle

### Stage 6 — Harvester Connectors (Productionized)
- source allowlist manager
- robust Git connector handling (rate-limit/retries)
- subscription ingestion scheduler

### Stage 7 — Probe + Scoring + Fallback Integration
- quick/deep probe workers
- scoring aggregation with history
- fallback policy execution in live sessions

### Stage 8 — Windows App MVP (UI + Core Bridge)
- shell screens (Home/Settings/IP Scanner/Logs/About)
- state sync with orchestrator
- theme profiles (Lite/Balanced/Rich + reduced motion)

### Stage 9 — Android App MVP
- profile selection/connect flows
- session status and log view
- shared configuration compatibility

### Stage 10 — Packaging & Release Hardening
- installer pipeline
- build metadata + checksums
- release validation checklist

### Stage 11 — Public Beta Readiness
- docs polish (EN/FA)
- diagnostics and support workflow
- bug triage and stabilization

## Estimated remaining stage count
**6 full major stages remain + completion items for Stage 5**.
