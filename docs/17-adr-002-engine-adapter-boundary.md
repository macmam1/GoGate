# ADR-002: Engine Adapter Boundary

## Status
Accepted

## Context
Project must integrate existing network engines without rewriting protocol cores.

## Decision
Adopt strict adapter boundary:
- engine process lifecycle management
- normalized start/stop/health contract
- capability map for orchestration

## Consequences
- easier engine swapping and A/B validation
- less technical debt from tightly-coupled integrations
- manifest-driven runtime control
