package engineadapters

import (
	"errors"
	"os/exec"
	"sync"
)

type ProcessSpec struct {
	Name string
	Args []string
}

type Supervisor struct {
	mu   sync.Mutex
	cmd  *exec.Cmd
	spec ProcessSpec
}

func NewSupervisor(spec ProcessSpec) *Supervisor {
	return &Supervisor{spec: spec}
}

func (s *Supervisor) Start() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.cmd != nil && s.cmd.Process != nil {
		return errors.New("process already running")
	}
	cmd := exec.Command(s.spec.Name, s.spec.Args...)
	if err := cmd.Start(); err != nil {
		return err
	}
	s.cmd = cmd
	return nil
}

func (s *Supervisor) Stop() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.cmd == nil || s.cmd.Process == nil {
		return nil
	}
	err := s.cmd.Process.Kill()
	s.cmd = nil
	return err
}

func (s *Supervisor) Running() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.cmd != nil && s.cmd.Process != nil
}
