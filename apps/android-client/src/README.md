# Android Executable Shell Scaffold

This folder contains a baseline Android shell scaffold for Stage 9 readiness.

Current files:
- `MainActivityShell.kt` state shell with bridge subscription
- `CoreBridgeModels.kt` bridge command/event models
- `CoreBridgeClient.kt` bridge client interface
- `MockCoreBridgeClient.kt` mock implementation for local testing

Next:
- replace mock bridge with real local bridge transport
- add session details and logs UI wiring
