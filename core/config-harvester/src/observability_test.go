package harvester

import "testing"

func TestInMemoryEventSink(t *testing.T) {
	sink := &InMemoryEventSink{}
	sink.Emit(NewFetchEvent("sourceA", 2, nil))
	if len(sink.Events) != 1 {
		t.Fatalf("expected one event")
	}
	if sink.Events[0].Name == "" {
		t.Fatalf("expected event name")
	}
}
