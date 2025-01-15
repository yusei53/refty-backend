package render

import (
	"encoding/json"
	"net/http"
	"refty-backend/internal/domain/infra/logging"

	"github.com/gin-gonic/gin"
)

func JSON(c *gin.Context, res interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(c.Writer).Encode(res); err != nil {
		logging.Warnf(c, "JSON json.NewEncoder %v", err)
	}
}

func OK(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(c.Writer).Encode(struct {
		OK bool `json:"ok"`
	}{true}); err != nil {
		logging.Warnf(c, "OK json.NewEncoder %v", err)
	}
}
