package http

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPostEventsReturnsNotImplementedUntilExercise(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	(&Handler{}).Register(router)
	response := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "/v1/events", nil)
	router.ServeHTTP(response, request)
	if response.Code != 501 {
		t.Fatalf("status = %d, want 501", response.Code)
	}
}
