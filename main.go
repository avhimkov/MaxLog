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
	g.LoadHTMLGlob("templates/*.html")

	// g := gin.New()

	// Logging middleware
	g.Use(gin.Logger())
	// Recovery middleware
	g.Use(gin.Recovery())
	// g.Use(static.Serve("/assets", static.LocalFile("./assets/favicon.ico", true)))
	// g.Use(favicon.New("./assets/favicon.ico"))

	// g.Use(static.Serve("/web", static.LocalFile("/web", false)))
	// v1 := router.Group("api/v1")
	// {
	// 	v1.GET("/instructions", GetInstructions)
	// }

	// g.Use(permissionHandler)

	return g
}

func main() {
	/* 	// SMS
	   	fileSMS, _ := os.Open("configsms.json")
	   	decoderSMS := json.NewDecoder(fileSMS)
	   	configurationSMS := ConfigSMS{}
	   	errSMS := decoderSMS.Decode(&configurationSMS)
	   	if errSMS != nil {
	   		log.Panic(errSMS)
	   	}

	   	// SMS
	   	login := configurationSMS.Login
	   	pass := configurationSMS.Password

	   	register("289", "name", "", "0", "6")

	   	// sms
	   	tele2("79827468271", login, pass, "MFC", "Hello")

	   	// push
	   	SendGCMToClient("Hello from GCM", "<CLIENT TOKEN>")
	*/

	g := SetupRouter()
	g.GET("/", indexPage)

	// 404 page
	g.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "404.html", gin.H{})
	})

	// Start serving the application
	g.Run(":3000")

}
