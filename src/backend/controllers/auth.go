package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rojack96/vierno/config"
)

func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", nil)
}

func PerformLogin(c *gin.Context) {
	user := c.PostForm("username")
	pass := c.PostForm("password")

	value, exists := c.Get("viernoConfig")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "configurazione mancante"})
		return
	}

	cfg := value.(*config.ViernoConfig)

	if user == cfg.Vierno.Auth.User && pass == cfg.Vierno.Auth.Password {
		// Imposta cookie valido per 1 ora
		c.SetCookie("session", "valid", 3600, "/", "", false, true)

		c.JSON(http.StatusOK, gin.H{
			"message": "Login effettuato con successo",
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Credenziali non valide",
		})
	}
}
