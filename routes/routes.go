package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rojack96/vierno/config"
	"github.com/rojack96/vierno/controllers"
	"github.com/rojack96/vierno/controllers/file_getter"
	"github.com/rojack96/vierno/middleware"
)

func SetupRouter(cfg *config.ViernoConfig) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	// Login page (solo per mostrare il form, non gestisce più la sessione)
	r.GET("/login", controllers.ShowLogin)

	// Tutte le route protette da Basic Auth
	protected := r.Group("/",
		middleware.ConfigRequired(cfg),
		middleware.AuthRequired("admin", "1234"),
	)
	{
		protected.GET(":app/:profile", file_getter.GetSimpleFile)
		protected.GET(":app", file_getter.GetFile)
	}

	return r
}
