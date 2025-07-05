package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rojack96/vierno/controllers"
	"github.com/rojack96/vierno/controllers/file_reader"
	//"github.com/rojack96/vierno/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/login", controllers.ShowLogin)
	r.POST("/login", controllers.PerformLogin)
	r.GET("/:app", file_reader.GetFile)
	r.GET("/:app/:profile", file_reader.GetSimpleFile)
	// r.GET("/config/:filename", config.GetFile)

	/*auth := r.Group("/", middleware.AuthRequired())
	{
		auth.GET("/fs", controllers.ListFiles)
		auth.GET("/fs/:filename", controllers.ShowEditFile)
		auth.POST("/fs/:filename", controllers.SaveFile)
	}*/

	return r
}
