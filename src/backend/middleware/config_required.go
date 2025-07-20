package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rojack96/vierno/config"
)

func ConfigRequired(cfg *config.ViernoConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("viernoConfig", cfg)
		c.Next()
	}
}

func DevMode(devMode bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("devMode", devMode)
		c.Next()
	}
}
