package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexPageGet(c *gin.Context) {

	fio := c.PostForm("fio")
	comment := c.PostForm("comment")
	serviceid := c.PostForm("serviceid")
	register(serviceid, fio, comment, "0", "")
	// {"Command":"cmd_Register","Number":"1","CustID":"44167","RegDateTime":"02.05.2019 00:08:51","QNT":"0","WaitTime":"-","ResultCode":"0"}

	http.Redirect(c.Writer, c.Request, "/terminal", 302)
}

func indexPagePost(c *gin.Context) {

	tiketList := getTicketSteps("0,5,6")
	// {"Command":"cmd_GetTicketSteps","TicketSteps":[{"TicketStepID":"49916","TicketNo":"77","CustID":"49449","CustData":"Ширкина А.П.",
	// "SourceKind":"1","State":"0","ServiceID":"190","RegTime":"08.05.2019 14:14:40","CallTime":"01.01.2000","PriorityID":"0","QualityMark":"0"}],"ResultCode":"0"}

	c.HTML(http.StatusOK, "terminal.html", gin.H{"tiketList": tiketList})
}
