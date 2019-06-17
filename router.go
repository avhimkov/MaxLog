package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service struct {
	ID                   string `json:"ID"`
	ShowElement          string `json:"ShowElement"`
	SRShowElement        string `json:"SRShowElement"` //SR_ShowElement
	Visible              string `json:"Visible"`
	Type                 string `json:"Type"`
	OrderNum             string `json:"OrderNum"`
	AllowWPorWUSelect    string `json:"AllowWPorWUSelect"`
	OnlyForSR            string `json:"OnlyForSR"`
	QueueID              string `json:"QueueID"`
	Name                 string `json:"Name"`
	ParentID             string `json:"ParentID"`
	State                string `json:"State"`
	NeedPriorityID       string `json:"NeedPriorityID"`
	NeedRate             string `json:"NeedRate"`
	OfferToOtherBranches string `json:"OfferToOtherBranches"`
}

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

	// serviceList := getServices("2", "0")
	/* for _, i := range listusers {
		cheked = append(cheked, userstate.IsAdmin(i))
	} */
	// fmt.Println(serviceList)

	// not work
	// SRTicketSteps := getSRTicketSteps("02.05.2019")
	// fmt.Println(SRTicketSteps)

	getServices := getServices("2", "0")

	// json.Unmarshal([]byte(received_JSON), &arbitrary_json)

	list := []Service{}

	for _, entry := range getServices {
		f := Service{
			ID:                   entry.ID(),
			ShowElement:          entry.ShowElement(),
			SRShowElement:        entry.SRShowElement(),
			Visible:              entry.Visible(),
			Type:                 entry.Type(),
			OrderNum:             entry.OrderNum(),
			AllowWPorWUSelect:    entry.AllowWPorWUSelect(),
			OnlyForSR:            entry.OnlyForSR(),
			QueueID:              entry.QueueID(),
			Name:                 entry.Name(),
			ParentID:             entry.ParentID(),
			State:                entry.State(),
			NeedPriorityID:       entry.NeedPriorityID(),
			NeedRate:             entry.NeedRate(),
			OfferToOtherBranches: entry.OfferToOtherBranches(),
		}
		list = append(list, f)
	}

	// output, err := json.Marshal(list)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(string(output))
	//https://gist.github.com/kousik93/6d95c4c4d37d8c731d7b
	// http://qaru.site/questions/2161532/golang-marshal-osfileinfo-into-json

	// fmt.Println(getServices)

	/* 	type ServiceList []Service
	   	result := gjson.Get(getServices, "Services")
	   	for _, name := range result.Array() {
	   		println(name.String())
	   		// ServiceList = name.String()
	   	} */

	// resultString := result.String()
	// ServiceList = resultString

	// result.ForEach(func(key, value gjson.Result) bool {
	// 	println(value.String())
	// 	return true // keep iterating
	// })

	// fmt.Println(value.String())

	/* 	workplaces := getWorkplaces()
	   	fmt.Println(workplaces)

	   	value := gjson.Get(workplaces, "Workplaces.#.Name")
	   	fmt.Println(value.String())
	*/
	/* getServiceByID := getServiceByID("289")
	fmt.Println(getServiceByID) */

	/* 	getTicketSteps := getTicketSteps("0,5,6")
	   	fmt.Println(getTicketSteps) */

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
