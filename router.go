package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexPageGet(c *gin.Context) {

	/*
		// sms
		tele2("79827468271", login, pass, "MFC", "Hello")

		// push
		SendGCMToClient("Hello from GCM", "<CLIENT TOKEN>")
	*/
	// for _, i := range tiketList {
	// 	cheked = append(cheked, userstate.IsAdmin(i))
	// }
	// fio := c.PostForm("fio")
	// comment := c.PostForm("comment")
	// serviceid := c.PostForm("serviceid")
	// register(serviceid, fio, comment, "0", "")
	// {"Command":"cmd_Register","Number":"1","CustID":"44167","RegDateTime":"02.05.2019 00:08:51","QNT":"0","WaitTime":"-","ResultCode":"0"}

	serviceList := getServices("2", "0")

	/* for _, i := range listusers {
		cheked = append(cheked, userstate.IsAdmin(i))
	} */
	fmt.Println(serviceList)
	c.HTML(http.StatusOK, "terminal.html", gin.H{})
	// http.Redirect(c.Writer, c.Request, "/terminal", 302)
}

func indexPagePost(c *gin.Context) {

	// tiketList := getTicketSteps("3")
	tiketList := Register("288", "Vany2", "", "0", "")

	// {"Command":"cmd_GetTicketSteps","TicketSteps":[{"TicketStepID":"49916","TicketNo":"77","CustID":"49449","CustData":"Ширкина А.П.",
	// "SourceKind":"1","State":"0","ServiceID":"190","RegTime":"08.05.2019 14:14:40","CallTime":"01.01.2000","PriorityID":"0","QualityMark":"0"}],"ResultCode":"0"}
	fmt.Println(tiketList)
	// c.HTML(http.StatusOK, "terminal.html", gin.H{"tiketList": tiketList})
	c.HTML(http.StatusOK, "terminal.html", gin.H{})
	//c.JSON(http.StatusOK, tiketList)

}
