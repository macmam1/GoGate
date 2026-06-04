package engineadapters

import (
	"fmt"
	"strings"
	"time"

	"github.com/macmam1/GoGate/core/internal/contracts"
)

type SingboxAdapter struct {
	binaryPath string
	sup        *Supervisor
	lastErr    error
}

func NewSingboxAdapter(binaryPath string, baseArgs []string) *SingboxAdapter {
	spec := ProcessSpec{Name: binaryPath, Args: append([]string{}, baseArgs...)}
	return &SingboxAdapter{binaryPath: binaryPath, sup: NewSupervisor(spec)}
}

func (a *SingboxAdapter) Capabilities() AdapterCapabilities {
	return AdapterCapabilities{EngineID: "sing-box", Protocols: []string{"vmess", "vless", "trojan", "hysteria2", "tuic"}, SupportsRouteOverride: true}
}

func (a *SingboxAdapter) Start(req contracts.SessionStartRequest) (SessionHandle, error) {
	args, err := buildSingboxArgs(req)
	if err != nil {
		a.lastErr = MapEngineError("sing-box", err, "")
		return SessionHandle{}, a.lastErr
	}

	a.sup.UpdateSpec(ProcessSpec{
		Name:    a.binaryPath,
		Args:    args,
		WorkDir: req.WorkingDir,
		Env:     mapEnv(req.Env),
	})

	if err := a.sup.Start(); err != nil {
		a.lastErr = MapEngineError("sing-box", err, a.sup.LastStderr())
		return SessionHandle{}, a.lastErr
	}
	return SessionHandle{ID: fmt.Sprintf("singbox-%d", time.Now().UnixNano()), EngineID: "sing-box"}, nil
}

func (a *SingboxAdapter) Health(handle SessionHandle) (HealthSnapshot, error) {
	if a.sup.Running() {
		return HealthSnapshot{Status: "running", LatencyMs: 0, HandshakeOK: true}, nil
	}
	if exitErr := a.sup.LastExitError(); exitErr != nil {
		return HealthSnapshot{Status: "crashed", LatencyMs: 0, HandshakeOK: false}, MapEngineError("sing-box", exitErr, a.sup.LastStderr())
	}
	return HealthSnapshot{Status: a.sup.LastStatus(), LatencyMs: 0, HandshakeOK: false}, nil
}

func (a *SingboxAdapter) Stop(handle SessionHandle) error {
	err := a.sup.Stop()
	if err != nil {
		a.lastErr = MapEngineError("sing-box", err, a.sup.LastStderr())
		return a.lastErr
	}
	return nil
}

func (a *SingboxAdapter) LastError() error { return a.lastErr }

func buildSingboxArgs(req contracts.SessionStartRequest) ([]string, error) {
	cfg := strings.TrimSpace(req.RuntimeConfigPath)
	if cfg == "" {
		return nil, fmt.Errorf("missing runtime config path")
	}
	args := []string{"run", "-c", cfg}
	if req.AdvancedModeEnable && req.Hints.EdgeIP != "" {
		args = append(args, "--route-address-set", req.Hints.EdgeIP)
	}
	args = append(args, req.AdditionalArgs...)
	return args, nil
}
