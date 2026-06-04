package harvester

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type SchedulerRunState struct {
	LastRunAt  time.Time `json:"lastRunAt"`
	LastStatus string    `json:"lastStatus"`
	LastError  string    `json:"lastError,omitempty"`
}

type SchedulerStateStore interface {
	Load(context.Context) (SchedulerRunState, error)
	Save(context.Context, SchedulerRunState) error
}

type FileSchedulerStateStore struct {
	Path string
	mu   sync.Mutex
}

func (s *FileSchedulerStateStore) Load(_ context.Context) (SchedulerRunState, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := os.ReadFile(s.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return SchedulerRunState{}, nil
		}
		return SchedulerRunState{}, err
	}
	var state SchedulerRunState
	if err := json.Unmarshal(data, &state); err != nil {
		return SchedulerRunState{}, err
	}
	return state, nil
}

func (s *FileSchedulerStateStore) Save(_ context.Context, state SchedulerRunState) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := os.MkdirAll(filepath.Dir(s.Path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.Path, data, 0o600)
}

func PersistentSchedulerHooks(store SchedulerStateStore, sink EventSink) SchedulerHooks {
	return SchedulerHooks{
		BeforeRun: func(at time.Time) {
			if sink != nil {
				sink.Emit(Event{Name: "harvester.scheduler.before", Timestamp: at, Fields: map[string]string{"result": "started"}})
			}
		},
		AfterRun: func(at time.Time, err error) {
			state := SchedulerRunState{LastRunAt: at}
			if err != nil {
				state.LastStatus = "error"
				state.LastError = err.Error()
			} else {
				state.LastStatus = "ok"
			}
			if store != nil {
				_ = store.Save(context.Background(), state)
			}
			if sink != nil {
				fields := map[string]string{"status": state.LastStatus}
				if state.LastError != "" {
					fields["error"] = state.LastError
				}
				sink.Emit(Event{Name: "harvester.scheduler.after", Timestamp: at, Fields: fields})
			}
		},
	}
}
