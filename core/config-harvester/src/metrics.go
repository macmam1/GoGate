package harvester

import "sync/atomic"

type IngestionMetrics struct {
	FetchSuccess int64
	FetchFail    int64
	ParseFail    int64
	ItemsIngest  int64
}

func (m *IngestionMetrics) RecordFetchSuccess(items int) {
	if m == nil {
		return
	}
	atomic.AddInt64(&m.FetchSuccess, 1)
	atomic.AddInt64(&m.ItemsIngest, int64(items))
}

func (m *IngestionMetrics) RecordFetchFail() {
	if m == nil {
		return
	}
	atomic.AddInt64(&m.FetchFail, 1)
}

func (m *IngestionMetrics) RecordParseFail() {
	if m == nil {
		return
	}
	atomic.AddInt64(&m.ParseFail, 1)
}
