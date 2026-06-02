# ADR-003: Harvester Source Policy

## Status
Accepted

## Context
Config ingestion from public sources can introduce noisy or unsafe entries.

## Decision
Use allowlist-first harvester policy:
- only approved source connectors
- parse + normalize + validate before candidate admission
- dedupe and metadata retention
- periodic refresh with bounded concurrency

## Consequences
- safer default behavior
- more predictable quality pipeline
- source governance required
