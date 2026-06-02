# Config Harvester Module

Responsibilities:
- Ingest subscription links and configured Git sources
- Normalize profile formats into internal schema
- Deduplicate entries and preserve source metadata
- Schedule incremental refreshes

Security notes:
- Allowlist-only source ingestion
- Validation before admission to candidate pool
