package engineadapters

import (
    "bytes"
    "errors"
    "os"
    "os/exec"
    "sync"
)

type ProcessSpec struct {
    Name    string
    Args    []string
    WorkDir string
    Env     []string
}

type Supervisor struct {
    mu         sync.Mutex
    cmd        *exec.Cmd
    spec       ProcessSpec
    stderr     bytes.Buffer
    stdout     bytes.Buffer
    lastExit   error
    lastStatus string
}

func NewSupervisor(spec ProcessSpec) *Supervisor {
    return &Supervisor{spec: spec, lastStatus: "idle"}
}

func (s *Supervisor) UpdateSpec(spec ProcessSpec) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.spec = spec
}

func (s *Supervisor) Start() error {
    s.mu.Lock()
    defer s.mu.Unlock()

    if s.cmd != nil && s.cmd.Process != nil {
        return errors.New("process already running")
    }
    if s.spec.Name == "" {
        return errors.New("empty binary name")
    }

    s.stdout.Reset()
    s.stderr.Reset()
    s.lastExit = nil
    s.lastStatus = "starting"

    cmd := exec.Command(s.spec.Name, s.spec.Args...)
    cmd.Stdout = &s.stdout
    cmd.Stderr = &s.stderr
    if s.spec.WorkDir != "" {
        cmd.Dir = s.spec.WorkDir
    }
    if len(s.spec.Env) > 0 {
        cmd.Env = append(os.Environ(), s.spec.Env...)
    }

    if err := cmd.Start(); err != nil {
        s.lastStatus = "failed"
        s.lastExit = err
        return err
    }
    s.cmd = cmd
    s.lastStatus = "running"

    go func(localCmd *exec.Cmd) {
        err := localCmd.Wait()
        s.mu.Lock()
        defer s.mu.Unlock()
        s.lastExit = err
        if err != nil {
            s.lastStatus = "crashed"
        } else {
            s.lastStatus = "exited"
        }
        s.cmd = nil
    }(cmd)

    return nil
}

func (s *Supervisor) Stop() error {
    s.mu.Lock()
    defer s.mu.Unlock()
    if s.cmd == nil || s.cmd.Process == nil {
        return nil
    }
    err := s.cmd.Process.Kill()
    s.lastExit = err
    s.lastStatus = "stopped"
    s.cmd = nil
    return err
}

func (s *Supervisor) Running() bool {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.cmd != nil && s.cmd.Process != nil
}

func (s *Supervisor) LastExitError() error {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.lastExit
}

func (s *Supervisor) LastStatus() string {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.lastStatus
}

func (s *Supervisor) LastStderr() string {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.stderr.String()
}
