package harvester

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	Name      string
	Timestamp time.Time
	Fields    map[string]string
}

type EventSink interface {
	Emit(Event)
}

type InMemoryEventSink struct {
	mu     sync.Mutex
	Events []Event
}

func (s *InMemoryEventSink) Emit(e Event) {
	if s == nil {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Events = append(s.Events, e)
}

func NewFetchEvent(source string, itemCount int, err error) Event {
	fields := map[string]string{
		"source": source,
		"items":  fmt.Sprintf("%d", itemCount),
	}
	if err != nil {
		fields["result"] = "error"
		fields["error"] = err.Error()
	} else {
		fields["result"] = "ok"
	}
	return Event{Name: "harvester.fetch", Timestamp: time.Now(), Fields: fields}
}
