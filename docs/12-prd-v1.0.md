# 12) PRD v1.0 — GoGate

## Product Statement
GoGate is a multi-platform connectivity client focused on resilient operation in restricted and unstable network conditions, with explicit user control, transparent behavior, and performance-aware UX.

## Primary Users
1. End users in unstable/restricted network conditions
2. Power users who need profile/source control
3. Maintainers/researchers who need logs and diagnostics

## Core Outcomes (Phase 1)
- Import/update config sources reliably
- Score and rank candidates based on measurable health signals
- Connect with fallback chain for higher success rates
- Provide clear status, logs, and user controls

## Non-Goals (Phase 1)
- iOS client
- complex enterprise policy management
- custom low-level protocol core rewrite

## Functional Requirements
1. Config source management
   - add/remove subscription links
   - add/remove allowlisted Git sources
   - manual import/export
2. Candidate validation
   - schema validation
   - duplicate elimination
   - source metadata retention
3. Probe pipeline
   - quick probe (latency/basic reachability)
   - deep probe (stability window, handshake success)
4. Scoring and ranking
   - weighted score (latency, stability, success rate)
   - history-aware re-ranking
5. Session orchestration
   - connect/disconnect/retry
   - multi-step fallback profiles
6. UI capabilities
   - connection dashboard
   - server list by location/category
   - logs and diagnostics export
   - theme profile selector (Lite/Balanced/Rich)
7. Sensitive/advanced mode gating
   - opt-in consent gate
   - clear warning and audit entry

## Quality Requirements
- startup time budget
- low idle CPU and memory targets
- predictable failure recovery
- deterministic logs

## Telemetry Policy
- default: local-only diagnostics
- optional: explicit opt-in usage telemetry later

## Branding/Identity Requirements
- product name: GoGate
- brand signature in app/about/docs: MRH-DevLoop
- English default language + Persian selectable language

## Release Requirements
- GitHub Releases with checksum artifacts
- bilingual user docs
- changelog and migration notes
