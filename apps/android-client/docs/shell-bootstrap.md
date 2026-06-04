# Android Shell Bootstrap (Stage 8/9)

## Scope
Bootstrap Android app shell with shared UX model and core-runtime-compatible state transitions.

## Initial screens
1. Home
2. Server List
3. Connection Details
4. Logs (compact)
5. Settings

## Integration contracts
- profile import + selection
- connect/disconnect commands
- session state rendering
- diagnostics export action

## Mobile-specific constraints
- battery-safe polling/event model
- background behavior limits
- reduced animation fallback for low-end devices

## Immediate tasks
- app shell navigation scaffold
- shared state model mapping
- status card + quick connect widget
