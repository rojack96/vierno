package routes

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rojack96/vierno/config"
	"github.com/rojack96/vierno/controllers"
	"github.com/rojack96/vierno/controllers/file_getter"
	"github.com/rojack96/vierno/controllers/file_reader"
	"github.com/rojack96/vierno/middleware"
)

func SetupRouter(cfg *config.ViernoConfig) *gin.Engine {
	r := gin.Default()
	// assetPath := "./dist/assets"
	// indexPath := "./dist/index.html"
	// Only development, so the assets are in the parent directory
	assetPath := "../../vierno-config-server/dist/assets"
	indexPath := "../../vierno-config-server/dist/index.html"
	r.Static("/assets", assetPath)
	r.LoadHTMLFiles(indexPath)

	// todo probabilmente solo per dev
	config := setupCors()

	r.Use(cors.New(config))

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Use(middleware.ConfigRequired(cfg))

	// Login page (solo per mostrare il form, non gestisce più la sessione)
	//r.GET("/login", controllers.ShowLogin)
	r.POST("/auth", controllers.PerformLogin)

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(204) // oppure serve un vero file favicon se vuoi
	})

	// Tutte le route protette da Basic Auth
	protected := r.Group("/",
		middleware.AuthMiddleware(cfg),
	)
	{
		protected.GET(":app/:profile", file_getter.GetSimpleFile)
		protected.GET(":app", file_getter.GetFile)

		foldersGroup := protected.Group("/folders")
		{
			foldersGroup.GET("lookup", file_reader.GetFolders)
			foldersGroup.GET(":folder", file_reader.GetFolder)
		}
	}

	return r
}

func setupCors() cors.Config {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	return config
}
