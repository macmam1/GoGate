# Fallback Flow (v1)

1. Select best candidate from scoring pool
2. Connect attempt with timeout budget
3. If fail -> classify reason (timeout/handshake/auth/network)
4. Apply next fallback candidate within same policy profile
5. If exhausted -> optional cross-profile fallback (if enabled)
6. Emit final state + diagnostics hint
