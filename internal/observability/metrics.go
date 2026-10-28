package observability

import (
	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
	"time"
)

type Metrics struct{}

func (m *Metrics) Record(result domain.BatchResult, duration time.Duration) {
	// TODO(student): record request, event, failure, and latency metrics.
}
