package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConfigSMS struct {
	Login    string
	Password string
}

func SetupRouter() *gin.Engine {

	// Set Gin to production mode
	//gin.SetMode(gin.ReleaseMode)

	g := gin.Default()

	g.Static("/web", "./web")
	g.Static("/assets", "./assets")
	g.Static("/node_modules", "./node_modules")
	g.StaticFile("/favicon.ico", "./assets/favicon.ico")
	g.LoadHTMLGlob("templates/*.html")

	// g := gin.New()

	// Logging middleware
	g.Use(gin.Logger())
	// Recovery middleware
	g.Use(gin.Recovery())

	// g.Use(static.Serve("/web", static.LocalFile("/web", false)))
	// v1 := router.Group("api/v1")
	// {
	// 	v1.GET("/instructions", GetInstructions)
	// }

	// g.Use(permissionHandler)

	return g
}

func main() {
	/*
		register("289", "name", "", "0", "6")

		// sms
		tele2("79827468271", login, pass, "MFC", "Hello")

		// push
		SendGCMToClient("Hello from GCM", "<CLIENT TOKEN>")
	*/

	g := SetupRouter()
	g.GET("/", indexPageGet)
	g.POST("/", indexPagePost)

	// 404 page
	g.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "404.html", gin.H{})
	})

	// Start serving the application
	g.Run(":3000")

}
