# 13) Engine + Harvester Integration Plan

## Objective
Integrate runtime engines and config ingestion without rewriting existing network cores.

## A) Engine Adapter Strategy
- Adapter interface: `start`, `stop`, `health`, `capabilities`, `lastError`
- Process supervision with restart policy
- Per-engine manifest-driven parameters

### Planned adapter order
1. xray adapter
2. sing-box adapter
3. psiphon-compatible adapter (distribution/legal validation before bundle mode)

### Psiphon note
- Two deployment modes:
  1) external-core mode (user-provided binary/runtime)
  2) bundled mode (only if licensing/distribution terms permit)

## B) Harvester Strategy
- Source types:
  - subscription URLs
  - allowlisted Git repositories
  - local imported files
- Allowlist format:
  - `core/config-harvester/sources.allowlist.txt`
  - one source-prefix per line
- Pipeline:
  `fetch -> parse -> normalize -> dedupe -> validate -> candidate pool`
- Scheduler:
  - periodic refresh
  - manual force refresh
  - delta update support

## C) Candidate Quality Pipeline
- quick probe
- deep probe
- score + rank
- policy filter
- connect attempt
- fallback if failure

## D) Safe Advanced Modes
- restricted/experimental modes remain opt-in
- explicit user notice and audit trail
- immediate disable path

## Deliverables for this track
- adapter interface contracts
- initial manifests
- harvester source connectors (first pass)
- scoring baseline config
