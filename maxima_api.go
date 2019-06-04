package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var gatewayURL = maxima()

type ConfigMax struct {
	IP   string
	Port string
}

// чтение конфигурационного файла
func maxima() string {
	// MAXIMA
	fileMax, _ := os.Open("./config/configMax.json")
	decoderMax := json.NewDecoder(fileMax)
	configurationMax := ConfigMax{}
	errMax := decoderMax.Decode(&configurationMax)
	if errMax != nil {
		log.Panic(errMax)
	}

	uri := "http://" + configurationMax.IP + ":5612/QMSSR"

	fmt.Println(uri)
	return uri
}

type reg struct {
	Command     string `json:"Command"`
	Number      string `json:"Number"`
	CustID      string `json:"CustID"`
	RegDateTime string `json:"RegDateTime"`
	QNT         string `json:"QNT"`
	WaitTime    string `json:"WaitTime"`
}

// Регистрация талона
func Register(serid, custdata, note, priorid, toid string) string /* *reg */ {
	// command=cmd_Register&ServiceID=289&CustData=dfgdgdfgdfgdfgdfg&Note=&PriorityID=0
	// response
	// {"Command":"cmd_Register","Number":"1","CustID":"44167","RegDateTime":"02.05.2019 00:08:51","QNT":"0","WaitTime":"-","ResultCode":"0"}

	urlStr := gatewayURL

	// Params
	v := url.Values{}
	v.Set("command", "cmd_Register")
	v.Set("ServiceID", serid)    //ID-услуги, на которую будет производиться регистрация
	v.Set("CustData", custdata)  //Информация о посетителе
	v.Set("Note", note)          // "" -- Примечание оператора
	v.Set("PriorityID", priorid) // "0" -- Приоритет посетителя
	v.Set("RegToID", toid)       // ID рабочего места или роли, к кому направить посетителя (по умолчанию(-1) - к любому)
	//v.Set("callTime", "0")                 // время вызова
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)

	respons := "'Command':'cmd_Register'"

	return respons

}

// Регистрация нескольких талонов
// multiRegister() - если в первой выбранной услуге указываться RegToId, то у всех последующих
// выбранных услугах, у которых RegToId не указан будет тот же RegToId, что и у первой услуги.
func multiRegisters(servcount, langid, custdata, note string) {
	// command=cmd_Registers
	// response

	urlStr := gatewayURL

	// Params
	v := url.Values{}
	v.Set("command", "cmd_Registers")
	v.Set("ServiceCount", servcount) // "289"
	v.Set("LangID", langid)          // "0"
	v.Set("CustData", custdata)      // ""
	v.Set("Note", note)              // ""
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)
}

// sr_Register
type srRegister struct {
	Command     string `json:"Command"`
	CheckResult string `json:"CheckResult"`
	Number      string `json:"Number"`
	CustID      string `json:"CustID"`
	RegDateTime string `json:"RegDateTime"`
	ResultCode  string `json:"ResultCode"`
}

// Регистрация на предварительную запись
func srReg(servid, custdata, note, priorid, calltime string) {
	// command=cmd_SR_Register&ServiceID=289&CustData=hjgjghj&Note=&PriorityID=0&CallTime=30.05.2019+8%3A00%3A00
	// response
	// {"Command":"cmd_SR_Register","CheckResult":"0","Number":"6","CustID":"49462","RegDateTime":"10.05.2019 12:00:00","ResultCode":"0"}

	urlStr := gatewayURL

	// Params
	v := url.Values{}
	v.Set("command", "cmd_SR_Register")
	v.Set("ServiceID", servid)   //"180"
	v.Set("CustData", custdata)  //name "Турчак"
	v.Set("Note", note)          //""
	v.Set("PriorityID", priorid) //"0"
	v.Set("CallTime", calltime)  //"10.05.2019 12:00:00"
	// SR_PIN: 1
	// RegToID: params.userId == -1 ? undefined : params.userId
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)
}

type workplace struct {
	Command    string `json:"Command"`
	Workplaces struct {
		ID      string `json:"ID"`
		Name    string `json:"Name"`
		PlaceNo string `json:"PlaceNo"`
		IDCon   string `json:"IDCon"`
		Active  string `json:"Active"`
		State   string `json:"State"`
	}
	ResultCode string `json:"ResultCode"`
}

// Получение рабочие места
func getWorkplaces() {
	// command=cmd_GetWorkplaces

	// gatewayURL := maxima()
	urlStr := gatewayURL

	// Params
	v := url.Values{}
	v.Set("command", "cmd_GetWorkplaces")
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)

}

type service struct {
	Command string `json:"Command"`
	Groups  struct {
		ID             string `json:"ID"`
		Visible        string `json:"Visible"`
		Type           string `json:"Type"`
		Level          string `json:"Level"`
		ShowElement    string `json:"ShowElement"`
		SRShowElement  string `json:"SR_ShowElement"` //SR_ShowElement
		OrderNum       string `json:"OrderNum"`
		Name           string `json:"Name"`
		NeedPriorityID string `json:"NeedPriorityID"`
		ParentID       string `json:"ParentID"`
	}
	ResultCode string `json:"ResultCode"`
}

// Получение список услуг
func getServices(typeofmenu, langid string) {
	// command=cmd_getservices&TypeOfMenu=2&LanguageID=0
	// response in file getService.json
	urlStr := gatewayURL

	// Params
	v := url.Values{}
	v.Set("command", "cmd_getservices")
	v.Set("TypeOfMenu", typeofmenu) //2
	v.Set("LanguageID", langid)     //0
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)
}

type serviceid struct {
	Command string `json:"Command"`
	Service struct {
		Name string `json:"Name"`
	}
	HandlingWP struct {
		ID                string `json:"ID"`
		Name              string `json:"Name"`
		Active            string `json:"Active"`
		BlockedForRegTurn string `json:"BlockedForRegTurn"`
		BlockedForRegSR   string `json:"BlockedForRegSR"`
	}
	Schedule struct {
		StartDate string `json:"StartDate"`
		EndDate   string `json:"EndDate"`
		Days      struct {
			Day       string `json:"Day"`
			StartTime string `json:"StartTime"`
			EndTime   string `json:"EndTime"`
		}
	}
	ResultCode string `json:"ResultCode"`
}

// Получение инфорамии об услуге по ID
func getServiceByID(servid string) {
	// command=cmd_GetServiceByID&GetHandlingWPorWU=1&ServiceID=289&GetWaitingCountPerWPOrWU=1
	// response in file getServiceByID.json

	urlStr := gatewayURL

	// Params
	v := url.Values{}
	v.Set("command", "cmd_GetServiceByID")
	v.Set("GetHandlingWPorWU", "1")
	v.Set("ServiceID", servid)
	v.Set("GetWaitingCountPerWPOrWU", "1")
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)
}

// Получение конфигурации
func getConfig() {
	// command=cmd_getconfig
	urlStr := gatewayURL

	// Params
	v := url.Values{}
	v.Set("command", "cmd_getconfig")
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)
}

type getTickStp struct {
	Command     string `json:"Command"`
	TicketSteps struct {
		TicketStepID string `json:"TicketStepID"`
		TicketNo     string `json:"TicketNo"`
		CustID       string `json:"CustID"`
		CustData     string `json:"CustData"`
		SourceKind   string `json:"SourceKind"`
		State        string `json:"State"`
		RegTime      string `json:"RegTime"`
		CallTime     string `json:"CallTime"`
		PriorityID   string `json:"PriorityID"`
		QualityMark  string `json:"QualityMark"`
	}
	ResultCode string
}

// Получение списка талонов
func getTicketSteps(state string) string {
	// command=cmd_GetTicketSteps&State=0%2C5%2C6
	// response
	// {"Command":"cmd_GetTicketSteps","TicketSteps":[{"TicketStepID":"49916","TicketNo":"77","CustID":"49449","CustData":"Ширкина А.П.",
	// "SourceKind":"1","State":"0","ServiceID":"190","RegTime":"08.05.2019 14:14:40","CallTime":"01.01.2000","PriorityID":"0","QualityMark":"0"}],"ResultCode":"0"}

	urlStr := gatewayURL

	// Params
	v := url.Values{}
	v.Set("command", "cmd_GetTicketSteps")
	v.Set("State", state) //"0,5,6", "0,5,6,3" 0 = , 3 = неявка заявителя (удаление талона из очереди), 5 = талон отложен , 6 = ожидающие
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)

	respons := "'Command':'cmd_GetTicketSteps'"

	return respons
}

type getSRTickStp struct {
	Command       string `json:"Command"`
	SRTicketSteps string `json:"SRTicketSteps"`
	ResultCode    string `json:"ResultCode"`
}

// Получение списка талонов по предварительной записи
func getSRTicketSteps(srdata string) {
	// command=cmd_GetSRTicketSteps&SRDate=02.05.2019
	// response
	// {"Command":"cmd_GetSRTicketSteps","SRTicketSteps":[],"ResultCode":"0"}

	//	if (endDate) {
	// 		data["SRDateTo"] = endDate
	// 	}

	urlStr := gatewayURL

	// Params
	v := url.Values{}
	v.Set("command", "cmd_GetSRTicketSteps")
	v.Set("SRDate", srdata)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)
}

type getLan struct {
	Command   string `json:"Command"`
	Languages struct {
		ID        string `json:"ID"`
		Name      string `json:"Name"`
		ShortName string `json:"ShortName"`
	}
	ResultCode string `json:"ResultCode	"`
}

// Получение еастройки языка
func getLanguages() {
	// command=cmd_GetLanguages
	// response
	// {"Command":"cmd_GetLanguages","Languages":[{"ID":"0","Name":"Русский","ShortName":"RUS"}],"ResultCode":"0"}

	urlStr := gatewayURL

	// Params
	v := url.Values{}
	v.Set("command", "cmd_GetLanguages")
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)
}

// Получение списка пользователей
func getGetWorkusers(id, name, last, patronimyc string) {
	// command=cmd_GetWorkusers
	// response

	//  // Получение рабочих мест
	// "cmd_GetWorkusers"
	// global.workusers[i] = {
	// 	id: res.WorkUsers[i].ID,
	// 	name: res.WorkUsers[i].FirstName || "",
	// 	last: res.WorkUsers[i].LastName || "",
	// 	second: res.WorkUsers[i].Patronimyc || ""
	// };

	urlStr := gatewayURL

	// Params
	v := url.Values{}
	v.Set("command", "cmd_GetWorkusers")
	v.Set("ID", id)
	v.Set("FirstName", name)
	v.Set("LastName", last)
	v.Set("Patronimyc", patronimyc)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)
}

// Получение информации о предварительной записи по пинкоду
//getSR_InformationByPIN
func getSRInformationByPIN(pin string) {
	// command=cmd_SR_InformationByPIN
	// response

	urlStr := gatewayURL

	// Params
	v := url.Values{}
	v.Set("command", "cmd_SR_InformationByPIN")
	v.Set("LangID", "0")
	v.Set("PIN", pin)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)
}

// получение свободной даты на которую можно сделать предварительную запись
func getIntervals(servid, datefrom, dateto string) {
	// command=cmd_GetIntervals
	// response

	//	command: "cmd_GetIntervals",
	//  ServiceID: params.id,
	//  DateFrom: params.date,
	//  DateTo: params.dateTo || utils.incdate(params.date, 1),
	//  RegToID: params.userId == -1 ? undefined : params.userId

	urlStr := gatewayURL

	// Params
	v := url.Values{}
	v.Set("command", "cmd_GetIntervals")
	v.Set("ServiceID", servid)  // "180"
	v.Set("DateFrom", datefrom) // "09.05.2019"
	v.Set("DateTo", dateto)     // "09.05.2019"
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)
}

// Обнавление статуса талона
func updateTicketStep(tiketid, state string) {
	// command=cmd_UpdateTicketStep
	// response

	urlStr := gatewayURL

	// Params
	v := url.Values{}
	v.Set("command", "cmd_UpdateTicketStep")
	v.Set("TicketStepID", tiketid) // "49978"
	v.Set("Note", "")
	v.Set("State", state) // 3 = неявка заявителя (удаление талона из очереди), 5 = талон отложен , 6 = ожидающие
	//{"Command":"cmd_UpdateTicketStep","QNT":"-1","WaitTime":"0:00:00","ResultCode":"0"}

	// v.Set("ServiceID", "")
	// v.Set("PriorityID", "")
	// v.Set("RatingID", "")
	// v.Set("CustData", "")
	// v.Set("RegToID", "")
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)
}

// Обнавление статуса талона предварительной записи
func srUpdateTicketStep(srticketstepid string) {
	// command=cmd_SRUpdateTicketStep
	// response

	urlStr := gatewayURL

	// Params
	v := url.Values{}
	v.Set("command", "cmd_SRUpdateTicketStep")
	v.Set("SRTicketStepID", srticketstepid) // ""
	// v.Set("ServiceID", "")
	// v.Set("CustData", "")
	// v.Set("Note", "")
	// v.Set("CallTime", "")
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)
}

// Удаление талона в предварительной записи
func srDelTicketStep(srtikcetstepid string) {
	// command=cmd_SRDelTicketStep
	// response

	urlStr := gatewayURL

	// Params
	v := url.Values{}
	v.Set("command", "cmd_SRDelTicketStep")
	v.Set("SRTicketStepID", srtikcetstepid) //
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)
}

// Регистрация талона предварительной записи по пин коду
// sr_RegisterByPIN
func srRegisterByPIN(pin string) {
	// command=cmd_SR_RegisterByPIN
	// response

	urlStr := gatewayURL

	// Params
	v := url.Values{}
	v.Set("command", "cmd_SR_RegisterByPIN")
	v.Set("LangID", "0")
	v.Set("PIN", pin)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)
}
