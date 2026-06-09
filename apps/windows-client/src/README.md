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

Next:
- add shell navigation host -> real page/view wiring
- expose bridge health indicator in UI status card
