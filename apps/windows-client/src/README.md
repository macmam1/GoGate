# Windows Executable Shell Scaffold

This folder contains a baseline executable shell scaffold for Stage 8.

Current files:
- `MainWindowShell.cs` state/event shell with bridge subscription
- `CoreBridgeContracts.cs` bridge command/event models
- `ICoreBridgeClient.cs` bridge client interface
- `MockCoreBridgeClient.cs` mock implementation for local testing

Next:
- replace mock bridge with real local bridge transport
- add navigation host and theme profile runtime binding
