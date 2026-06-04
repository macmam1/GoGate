package scoring

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

type HistoryStore interface {
	Record(profileID string, success bool) error
	RecentSuccess(profileID string) (float64, bool)
}

type profileHistory struct {
	Successes int `json:"successes"`
	Failures  int `json:"failures"`
}

type MemoryHistoryStore struct {
	mu   sync.Mutex
	data map[string]profileHistory
}

func NewMemoryHistoryStore() *MemoryHistoryStore {
	return &MemoryHistoryStore{data: map[string]profileHistory{}}
}

func (s *MemoryHistoryStore) Record(profileID string, success bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	v := s.data[profileID]
	if success {
		v.Successes++
	} else {
		v.Failures++
	}
	s.data[profileID] = v
	return nil
}

func (s *MemoryHistoryStore) RecentSuccess(profileID string) (float64, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	v, ok := s.data[profileID]
	if !ok {
		return 0, false
	}
	total := v.Successes + v.Failures
	if total == 0 {
		return 0, false
	}
	return float64(v.Successes) / float64(total), true
}

type FileHistoryStore struct {
	Path string
	mu   sync.Mutex
}

func (s *FileHistoryStore) Record(profileID string, success bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	m, err := s.loadUnsafe()
	if err != nil {
		return err
	}
	v := m[profileID]
	if success {
		v.Successes++
	} else {
		v.Failures++
	}
	m[profileID] = v
	return s.saveUnsafe(m)
}

func (s *FileHistoryStore) RecentSuccess(profileID string) (float64, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	m, err := s.loadUnsafe()
	if err != nil {
		return 0, false
	}
	v, ok := m[profileID]
	if !ok {
		return 0, false
	}
	total := v.Successes + v.Failures
	if total == 0 {
		return 0, false
	}
	return float64(v.Successes) / float64(total), true
}

func (s *FileHistoryStore) loadUnsafe() (map[string]profileHistory, error) {
	data, err := os.ReadFile(s.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return map[string]profileHistory{}, nil
		}
		return nil, err
	}
	if len(data) == 0 {
		return map[string]profileHistory{}, nil
	}
	out := map[string]profileHistory{}
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *FileHistoryStore) saveUnsafe(m map[string]profileHistory) error {
	if err := os.MkdirAll(filepath.Dir(s.Path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.Path, data, 0o600)
}
