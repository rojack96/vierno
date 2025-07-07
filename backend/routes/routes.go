package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rojack96/vierno/config"
	"github.com/rojack96/vierno/controllers"
	"github.com/rojack96/vierno/controllers/file_getter"
	"github.com/rojack96/vierno/middleware"
)

func SetupRouter(cfg *config.ViernoConfig) *gin.Engine {
	r := gin.Default()
	r.Static("/assets", "./frontend/dist/assets") // vite mette tutto in /assets
	r.LoadHTMLFiles("./frontend/dist/index.html")

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Login page (solo per mostrare il form, non gestisce più la sessione)
	r.GET("/login", controllers.ShowLogin)
	r.POST("/login", controllers.PerformLogin)

	// Tutte le route protette da Basic Auth
	protected := r.Group("/",
		middleware.ConfigRequired(cfg),
		middleware.AuthMiddleware(),
	)
	{
		protected.GET(":app/:profile", file_getter.GetSimpleFile)
		protected.GET(":app", file_getter.GetFile)
	}

	return r
}
