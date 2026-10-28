package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
	"github.com/kyawmyintthein/event-ingestion-gateway/internal/ingestion"
)

type Handler struct {
	Ingestor     ingestion.Ingestor
	MaxBodyBytes int64
	MaxEvents    int
	BatchMode    ingestion.BatchMode
	DeliveryMode ingestion.DeliveryMode
}

func (h *Handler) PostEvents(c *gin.Context) {
	// TODO(student): limit and parse a top-level JSON array, dispatch, map safe results/status.
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not_implemented", "message": domain.ErrNotImplemented.Error()})
}

func (h *Handler) Healthz(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) }

func (h *Handler) Readyz(c *gin.Context) {
	// TODO(student, advanced): check Kafka reachability with a bounded context.
	c.JSON(http.StatusOK, gin.H{"status": "ready"})
}

func (h *Handler) Register(r gin.IRouter) {
	r.POST("/v1/events", h.PostEvents)
	r.GET("/healthz", h.Healthz)
	r.GET("/readyz", h.Readyz)
}
