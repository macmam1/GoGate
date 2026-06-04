# 20) Stage 6 Production Hardening Plan

## Scope
Turn baseline harvester components into production-grade services.

## Work packages
1. Policy-backed retry tuning
   - expose retry config via runtime policy
   - per-source retry profile support
2. Scheduler persistence
   - hook interface wired to state store
   - remember last successful sync timestamp
3. Ingestion telemetry
   - counters: fetch_success, fetch_fail, parse_fail, dedupe_drop
   - per-source error classification
4. Connector resilience
   - distinguish auth/rate-limit/transient/server errors
   - clear operator-facing diagnostics messages

## Exit criteria
- deterministic retries with bounded backoff
- recoverable behavior across transient outages
- observable ingestion pipeline health
