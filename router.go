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
	/* for _, i := range listusers {
		cheked = append(cheked, userstate.IsAdmin(i))
	} */
	// fmt.Println(serviceList)

	// not work
	// SRTicketSteps := getSRTicketSteps("02.05.2019")
	// fmt.Println(SRTicketSteps)

	// getServices := getServices("2", "0")

	/* 	var result map[string]interface{}
	   	json.Unmarshal([]byte(getServices), &result)
	   	birds := result["Services"].(map[string]interface{})

	   	for key, value := range birds {
	   		// Each value is an interface{} type, that is type asserted as a string
	   		fmt.Println(key, value.(string))
	   	} */

	// resultString := result.String()
	// ServiceList = resultString

	/* 	workplaces := getWorkplaces()
	   	fmt.Println(workplaces)

	   	value := gjson.Get(workplaces, "Workplaces.#.Name")
	   	fmt.Println(value.String())
	*/

	/* 	getServiceByID := getServiceByID("289")
	fmt.Println(getServiceByID) */

	type Users struct {
		TicketStepID string
		CallTime     string
		RegTime      string
	}

	getTicketSteps := getTicketSteps("0,5,6")

	var ticketSteps GetTickStp

	err := json.Unmarshal([]byte(getTicketSteps), &ticketSteps)
	if err != nil {
		panic(err)
	}
	var steps []Users
	for i := 0; i < len(ticketSteps.TicketSteps); i++ {

		// fmt.Printf("%+v\n", ticketSteps.TicketSteps[i])

		steps = append(steps, Users{
			TicketStepID: ticketSteps.TicketSteps[i].TicketStepID,
			CallTime:     ticketSteps.TicketSteps[i].CallTime,
			RegTime:      ticketSteps.TicketSteps[i].RegTime,
		})

	}
	fmt.Printf("%+v\n", steps)

	/* 	getLanguages := getLanguages()
	   	fmt.Println(getLanguages) */

	// no tested
	/* getGetWorkusers := getGetWorkusers()
	fmt.Println(getGetWorkusers) */

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

func dumpJSON(v interface{}, kn string) {
	iterMap := func(x map[string]interface{}, root string) {
		var knf string
		if root == "root" {
			knf = "%q:%q"
		} else {
			knf = "%s:%q"
		}
		for k, v := range x {
			dumpJSON(v, fmt.Sprintf(knf, root, k))
		}
	}

	iterSlice := func(x []interface{}, root string) {
		var knf string
		if root == "root" {
			knf = "%q:[%d]"
		} else {
			knf = "%s:[%d]"
		}
		for k, v := range x {
			dumpJSON(v, fmt.Sprintf(knf, root, k))
		}
	}

	switch vv := v.(type) {
	case string:
		fmt.Printf("%s => (string) %q\n", kn, vv)
	case bool:
		fmt.Printf("%s => (bool) %v\n", kn, vv)
	case float64:
		fmt.Printf("%s => (float64) %f\n", kn, vv)
	case map[string]interface{}:
		fmt.Printf("%s => (map[string]interface{}) ...\n", kn)
		iterMap(vv, kn)
	case []interface{}:
		fmt.Printf("%s => ([]interface{}) ...\n", kn)
		iterSlice(vv, kn)
	default:
		fmt.Printf("%s => (unknown?) ...\n", kn)
	}
}
