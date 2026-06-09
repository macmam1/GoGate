# Android Executable Shell Scaffold

This folder contains a baseline Android shell scaffold for Stage 9 readiness.

Current files:
- `MainActivityShell.kt` state shell with bridge subscription
- `NavigationHost.kt` route management host
- `CoreBridgeModels.kt` bridge command/event models
- `CoreBridgeClient.kt` bridge client interface
- `MockCoreBridgeClient.kt` mock implementation for local testing
- `LocalRpcCoreBridgeClient.kt` real local-rpc transport client baseline (+ reconnect/backoff hardening)
- `ThemeRuntime.kt` runtime theme profile binding
- `ShellModels.kt` profile/session runtime models
- `BridgeHealthIndicator.kt` health badge view-model mapper
- `ShellPageBindings.kt` runtime-to-page binding layer
- `SharedProfileConfig.kt` shared profile format parser/serializer

Next:
- bind route host to concrete Android navigation/pages
- wire bridge health indicator widget in status card
- shared profile import/export flows in UI
