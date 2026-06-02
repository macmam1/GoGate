package engineadapters

import (
	"fmt"
	"time"

	"github.com/macmam1/GoGate/core/internal/contracts"
)

type XrayAdapter struct {
	sup     *Supervisor
	lastErr error
}

func NewXrayAdapter(binaryPath string, args []string) *XrayAdapter {
	return &XrayAdapter{sup: NewSupervisor(ProcessSpec{Name: binaryPath, Args: args})}
}

func (a *XrayAdapter) Capabilities() AdapterCapabilities {
	return AdapterCapabilities{EngineID: "xray", Protocols: []string{"vmess", "vless", "trojan", "shadowsocks"}, SupportsRouteOverride: true}
}

func (a *XrayAdapter) Start(req contracts.SessionStartRequest) (SessionHandle, error) {
	if err := a.sup.Start(); err != nil {
		a.lastErr = err
		return SessionHandle{}, err
	}
	return SessionHandle{ID: fmt.Sprintf("xray-%d", time.Now().UnixNano()), EngineID: "xray"}, nil
}

func (a *XrayAdapter) Health(handle SessionHandle) (HealthSnapshot, error) {
	if !a.sup.Running() {
		return HealthSnapshot{Status: "stopped"}, nil
	}
	return HealthSnapshot{Status: "running", LatencyMs: 0, HandshakeOK: true}, nil
}

func (a *XrayAdapter) Stop(handle SessionHandle) error {
	err := a.sup.Stop()
	if err != nil {
		a.lastErr = err
	}
	return err
}

func (a *XrayAdapter) LastError() error { return a.lastErr }
