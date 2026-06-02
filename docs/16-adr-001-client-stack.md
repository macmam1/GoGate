# ADR-001: Client Stack Strategy

## Status
Proposed (needs owner approval)

## Context
GoGate requires:
- Windows + Android in phase 1
- high-quality animated UI
- performance modes and reduced motion
- fast iteration speed

## Decision (proposed)
Use a shared UI stack for phase 1 clients to reduce duplicate effort.

### Candidate options considered
1. Flutter
2. Separate native stacks (Windows + Android)
3. Electron + Android native split

### Preferred option
Flutter for app shell + platform-specific adapter host bridges.

## Rationale
- strong animation tooling
- consistent UX across platforms
- easier bilingual UI parity
- performance tuning tools available

## Consequences
- requires platform channel layer for engine adapter integration
- packaging and native interop setup needed per platform
