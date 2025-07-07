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
		// Salvare sessione o token nel cookie
		c.SetCookie("session", "valid", 3600, "/", "", false, true)
		c.Redirect(http.StatusFound, "/dashboard")
	} else {
		c.HTML(http.StatusUnauthorized, "login.tmpl", gin.H{"Error": "Invalid credentials"})
	}

}
