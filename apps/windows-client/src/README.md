# Windows Executable Shell Scaffold

This folder contains a baseline executable shell scaffold for Stage 8.

Current files:
- `MainWindowShell.cs` state/event shell with bridge subscription
- `NavigationHost.cs` route management host
- `CoreBridgeContracts.cs` bridge command/event models
- `ICoreBridgeClient.cs` bridge client interface
- `MockCoreBridgeClient.cs` mock implementation for local testing
- `LocalRpcCoreBridgeClient.cs` real local-rpc transport client baseline (+ reconnect/backoff hardening)
- `ThemeRuntime.cs` runtime theme profile binding
- `ShellModels.cs` profile/session runtime models
- `BridgeHealthIndicator.cs` health badge view-model mapper
- `ShellPageBindings.cs` runtime-to-page binding layer
- `SharedProfileConfig.cs` shared profile format parser/serializer

Next:
- bind shell navigation host/pages to concrete UI views
- wire bridge health indicator widget in status card
- shared profile import/export flows in UI
