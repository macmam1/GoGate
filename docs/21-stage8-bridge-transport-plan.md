# 21) Stage 8 Bridge Transport Plan

## Objective
Replace mock bridge clients in shells with real local transport for runtime commands/events.

## Proposed transport
- Desktop: local named pipe / localhost unix-domain equivalent
- Android: local IPC/service bridge abstraction (app-safe)

## Command flow
1. Shell sends command (`connect`, `disconnect`, `rank_candidates`, `fetch_logs`)
2. Core bridge validates payload and executes
3. Response returned to shell
4. async events emitted to shell subscribers (`session_state_changed`, `ranking_updated`, etc.)

## Implementation tasks
- define serialized payload schema (JSON)
- bridge client transport adapters per platform
- reconnect/backoff behavior when core bridge unavailable
- security checks for local bridge endpoint ownership

## Exit criteria
- shells operate without mock clients
- state/event sync driven by real core bridge
- recover gracefully on bridge disconnect
