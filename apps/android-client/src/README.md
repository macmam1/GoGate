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

Next:
- wire route host to actual Android navigation/pages
- bind profile list/session details/log widgets to shell runtime methods
- shared configuration compatibility wiring
