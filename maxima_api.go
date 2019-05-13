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

	return uri

}

// Регистрация талона
func register(serid, custdata, note, priorid, toid string) {
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

// Регистрация на предварительную запись
func sr_Register(servid, custdata, note, priorid, calltime string) {
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

// Получение списка талонов
func getTicketSteps(state string) {
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
func getSR_InformationByPIN(pin string) {
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
func sr_RegisterByPIN(pin string) {
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

// 			// Проверка соединения с сервером
// 			checkConnection: function () {
// 				var self = this;
// 				$.ajax({
// 					async: true,
// 					dataType: "json",
// 				});
// 			},
