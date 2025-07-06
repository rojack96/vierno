package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowLogin(c *gin.Context) {
	redirect := c.Query("redirect")
	c.HTML(http.StatusOK, "login.tmpl", gin.H{"Redirect": redirect})
}

func PerformLogin(c *gin.Context) {
	user := c.PostForm("username")
	pass := c.PostForm("password")
	redirect := c.PostForm("redirect")
	if redirect == "" {
		redirect = c.Query("redirect")
	}

	if user == "admin" && pass == "1234" {
		c.SetCookie("session", "valid", 3600, "/", "", false, true)
		if redirect != "" {
			c.Redirect(http.StatusFound, redirect)
		} else {
			c.Redirect(http.StatusFound, "/dashboard")
		}
	} else {
		c.HTML(http.StatusUnauthorized, "login.tmpl", gin.H{"Error": "Invalid credentials", "Redirect": redirect})
	}
}
