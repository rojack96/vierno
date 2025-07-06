package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rojack96/vierno/config"
)

func ConfigRequired(cfg *config.ViernoConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Puoi salvare la config nel contesto per recuperarla nei controller
		c.Set("viernoConfig", cfg)
		c.Next()
	}
}
