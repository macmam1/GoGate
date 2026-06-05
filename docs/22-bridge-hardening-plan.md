# 22) Bridge Hardening Plan (Stage 8)

## Goal
Make shell bridge transports resilient when core runtime is unavailable or unstable.

## Work items
1. Exponential reconnect backoff with cap
2. Circuit-breaker style temporary cooldown after repeated failures
3. Command timeout classification and user-facing reason mapping
4. Event poll dedupe + sequence guards
5. Bridge health indicator for UI status card
6. Graceful fallback to mock/offline mode for developer testing

## Exit criteria
- no tight retry loops under bridge outage
- user-visible state transitions remain deterministic
- logs include bridge failure reasons and recoveries
