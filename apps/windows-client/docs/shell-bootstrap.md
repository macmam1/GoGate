# Windows Shell Bootstrap (Stage 8)

## Scope
Bootstrap a desktop shell that binds to orchestrator/core runtime and exposes primary user flows.

## Initial screens
1. Home (connection status + connect/disconnect)
2. Settings (connection, advanced mode gate, themes)
3. IP Scanner (results, apply-to-profile)
4. Logs (structured diagnostics + export)
5. About (version/build/signature MRH-DevLoop)

## Integration contracts
- Core bridge endpoint: local runtime bridge (to be finalized in ADR-004)
- Event streams consumed:
  - session state changes
  - fetch/ingestion events
  - scoring/ranking updates

## Performance profile controls
- Lite / Balanced / Rich
- Reduced Motion toggle

## Immediate tasks
- shell navigation scaffold
- state store interface
- mock data wiring for UI test pass
