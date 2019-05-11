package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexPage(c *gin.Context) {
	isloggedin := isloggedin(c)
	if isloggedin {
		c.HTML(http.StatusOK, "operator.html", gin.H{"is_logged_in": isloggedin})
	} else {
		http.Redirect(c.Writer, c.Request, "/login", 302)
	}
}
