package main

import (
	"encoding/json"
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

	// serviceList := getServices("2", "0")
	// fmt.Println(serviceList)

	// not work
	// SRTicketSteps := getSRTicketSteps("02.05.2019")
	// fmt.Println(SRTicketSteps)

	getServices := getServices("2", "0")

	var groups3 Groups3

	err := json.Unmarshal([]byte(getServices), &groups3)
	if err != nil {
		panic(err)
	}
	var service []Services
	for i := 0; i < len(groups3.Services); i++ {

		service = append(service, Services{
			Name: groups3.Services[i].Name,
			// Name: services.Groups1[i].Groups2[i].Groups3[i].Name,
		})
	}

	fmt.Printf("%+v\n", service)

	// resultString := result.String()
	// ServiceList = resultString

	/* 	getServiceByID := getServiceByID("289")
	fmt.Println(getServiceByID) */

	getTicketSteps := getTicketSteps("0,5,6")

	var ticketSteps GetTickStp

	err1 := json.Unmarshal([]byte(getTicketSteps), &ticketSteps)
	if err1 != nil {
		panic(err1)
	}
	var tiketstep []TicketSteps
	for i := 0; i < len(ticketSteps.TicketSteps); i++ {

		tiketstep = append(tiketstep, TicketSteps{
			TicketStepID: ticketSteps.TicketSteps[i].TicketStepID,
			TicketNo:     ticketSteps.TicketSteps[i].TicketNo,
			CustID:       ticketSteps.TicketSteps[i].CustID,
			CustData:     ticketSteps.TicketSteps[i].CustData,
			SourceKind:   ticketSteps.TicketSteps[i].SourceKind,
			State:        ticketSteps.TicketSteps[i].State,
			ServiceID:    ticketSteps.TicketSteps[i].ServiceID,
			RegTime:      ticketSteps.TicketSteps[i].RegTime,
			CallTime:     ticketSteps.TicketSteps[i].CallTime,
			PriorityID:   ticketSteps.TicketSteps[i].PriorityID,
			QualityMark:  ticketSteps.TicketSteps[i].QualityMark,
		})
	}
	// fmt.Printf("%+v\n", tiketstep)

	/* 	getLanguages := getLanguages()
	   	fmt.Println(getLanguages) */

	// no tested
	/* getGetWorkusers := getGetWorkusers()
	fmt.Println(getGetWorkusers) */

	c.HTML(http.StatusOK, "terminal.html", gin.H{"tiketstep": tiketstep, "service": service})
	// http.Redirect(c.Writer, c.Request, "/terminal", 302)
}

func indexPagePost(c *gin.Context) {
	c.Request.ParseForm()
	serviceid := c.PostForm("serviceid")
	fio := c.PostForm("fio")
	comment := c.PostForm("comment")
	fmt.Printf("serviceid: %v; fio: %v; comment: %v", serviceid, fio, comment)

	// tiketList := getTicketSteps("3")

	Register(serviceid, fio, comment, "0", "")
	// tiketList := Register("288", "Vany3", "", "0", "")

	// {"Command":"cmd_GetTicketSteps","TicketSteps":[{"TicketStepID":"49916","TicketNo":"77","CustID":"49449","CustData":"Ширкина А.П.",
	// "SourceKind":"1","State":"0","ServiceID":"190","RegTime":"08.05.2019 14:14:40","CallTime":"01.01.2000","PriorityID":"0","QualityMark":"0"}],"ResultCode":"0"}
	// fmt.Println(tiketList)
	// c.HTML(http.StatusOK, "terminal.html", gin.H{"tiketList": tiketList})
	http.Redirect(c.Writer, c.Request, "/", 302)
	//c.JSON(http.StatusOK, tiketList)

}
