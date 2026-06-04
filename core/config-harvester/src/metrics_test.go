package harvester

import "testing"

func TestIngestionMetrics(t *testing.T) {
	m := &IngestionMetrics{}
	m.RecordFetchSuccess(3)
	m.RecordFetchFail()
	m.RecordParseFail()
	if m.FetchSuccess != 1 || m.FetchFail != 1 || m.ParseFail != 1 || m.ItemsIngest != 3 {
		t.Fatalf("unexpected metric values: %+v", m)
	}
}
