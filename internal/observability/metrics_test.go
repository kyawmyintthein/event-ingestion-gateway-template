package observability

import (
	"testing"
	"time"

	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
)

func TestMetricsExercisePlaceholderIsSafeToCall(t *testing.T) {
	(&Metrics{}).Record(domain.BatchResult{}, time.Millisecond)
}
