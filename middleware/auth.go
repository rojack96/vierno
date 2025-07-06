package middleware

import (
	"encoding/base64"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthRequired(username, password string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if strings.HasPrefix(auth, "Basic ") {
			payload, _ := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth, "Basic "))
			parts := strings.SplitN(string(payload), ":", 2)
			if len(parts) == 2 && parts[0] == username && parts[1] == password {
				c.Next()
				return
			}
		}
		// Se richiesta da browser (Accept HTML), mostra la pagina di login con redirect
		accept := c.GetHeader("Accept")
		if strings.Contains(accept, "text/html") {
			redirectURL := c.Request.RequestURI
			c.Redirect(http.StatusFound, "/login?redirect="+url.QueryEscape(redirectURL))
			c.Abort()
			return
		}
		// Altrimenti, prompt Basic Auth
		c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
