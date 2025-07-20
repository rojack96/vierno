package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rojack96/vierno/config"
)

func AuthMiddleware(cfg *config.ViernoConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if strings.HasPrefix(auth, "Basic ") {
			payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth, "Basic "))
			if err == nil {
				parts := strings.SplitN(string(payload), ":", 2)
				if len(parts) == 2 && parts[0] == cfg.Vierno.Auth.User && parts[1] == cfg.Vierno.Auth.Password {
					c.Next()
					return
				}
			}
		}

		if cookie, err := c.Cookie("session"); err == nil && cookie == "valid" {
			c.Next()
			return
		}

		// Se la richiesta sembra provenire da un browser, fai redirect alla login
		accept := c.GetHeader("Accept")
		if strings.Contains(accept, "text/html") || strings.Contains(accept, "application/xhtml+xml") {
			c.Redirect(http.StatusFound, "/")
			c.Abort()
			return
		}

		// Altrimenti, risposta 401 per client API
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}
