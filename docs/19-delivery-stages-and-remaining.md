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

### Stage 6 — Harvester Connectors (in progress)
- [x] source allowlist manager
- [x] robust Git connector handling (rate-limit/retries) baseline
- [x] subscription ingestion scheduler baseline
- [ ] persistence hooks for scheduler state storage integration
- [ ] ingestion metrics counters/events
- [ ] backoff tuning from runtime policy

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
**6 major stages remain** (Stages 6 to 11).
