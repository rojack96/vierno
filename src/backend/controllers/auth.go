package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", nil)
}

func PerformLogin(c *gin.Context) {
	user := c.PostForm("username")
	pass := c.PostForm("password")

	if user == "admin" && pass == "1234" {
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
