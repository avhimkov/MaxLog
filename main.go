package main

import (
	"log"
	"net/http"

	"github.com/asdine/storm"

	"github.com/gin-gonic/gin"
)

//middlleware gin config
/*
 func permHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set up a middleware handler for Gin, with a custom "permission denied" message.
		// permissionHandler := func(c *gin.Context) {
		// Check if the user has the right admin/user rights
		if perm.Rejected(c.Writer, c.Request) {
			// Deny the request, don't call other middleware handlers
			c.AbortWithStatus(http.StatusForbidden)
			fmt.Fprint(c.Writer, "Permission denied!")
			return
		}
		// Call the next middleware handler
		c.Next()
	}
}
*/

/*
func isloggedin(c *gin.Context) bool {
	usercook, _ := userstate.UsernameCookie(c.Request)
	isloggedin := userstate.IsLoggedIn(usercook)
	return isloggedin
}
*/

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

	/*
		g.Use(static.Serve("/web", static.LocalFile("/web", false)))
		v1 := router.Group("api/v1")
		{
			v1.GET("/instructions", GetInstructions)
		}
	*/

	// g.Use(permissionHandler)

	return g
}

var db = DB()

//middlleware db
/*
var perm, err = perminit("db/bolt.db")
*/
//middlleware var
/*
var userstate = perm.UserState()
var permissionHandler = permHandler()
*/

//open databas
func DB() *storm.DB {
	db, err := storm.Open("db/data.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {

	defer db.Close()

	// Default user /administrator admin
	/*
		userstate.AddUser("admin", "admin", "admin@mail.ru")
		userstate.MarkConfirmed("admin")
		userstate.SetAdminStatus("admin")
	*/

	//init struct in boltdb
	/*
		errdbp := db.Init(&Person{})
		if errdbp != nil {
			log.Fatal(errdbp)
		}
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
