package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func newTicket(phone string) {

	//POST queue/admin/tickets/new
	// Set variable values
	gatewayURL := "queue/admin/tickets/new"
	ip := "192.168.10.176"
	port := "5600"
	urlStr := ip + port + gatewayURL

	// Params
	v := url.Values{}
	v.Set("ip", ip)
	v.Set("port", port)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)

}

func serviceInTicket(phone string) {

	// Set variable values
	gatewayURL := "queue/admin/tickets/:ticketId/addProduct/:productId"
	ip := "192.168.10.176"
	port := "5600"
	urlStr := ip + port + gatewayURL

	// Params
	v := url.Values{}
	v.Set("ip", ip)
	v.Set("port", port)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)

}

func cancelServiceTiket(phone string) {
	//queue/admin/tickets/:ticketId/:ticketProductId/cancel
	// Set variable values
	gatewayURL := "queue/admin/tickets/:ticketId/:ticketProductId/cancel"
	ip := "192.168.10.176"
	port := "5600"
	urlStr := ip + port + gatewayURL

	// Params
	v := url.Values{}
	v.Set("ip", ip)
	v.Set("port", port)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)

}

func prioritetServiceTiket(phone string) {
	//POST
	//queue/admin/tickets/:ticketId/:ticketProductId/setPriority

	// Set variable values

}

func returnOpertorList(phone string) {
	//POST
	//queue/admin/redirect/users/:productId

	// Set variable values

}

func returnDesktop(phone string) {
	//POST
	//queue/admin/redirect/workPlaces/:productId

	// Set variable values

}

func redirectUser(phone string) {
	//POST
	// queue/admin/redirect/:productId

	// Set variable values

}

func getInfoTiket(phone string) {
	//GET
	//queue/admin/tickets/:ticketId

	// Set variable values

}

func returnAllTiketsPredList(phone string) {
	//GET
	//queue/admin/booking

	// Set variable values

}

func returnAllTikets(phone string) {
	//GET
	//queue/admin/tickets

	// Set variable values

}

func returnTokenAuth(phone string) {
	//POST
	//queue/admin/login/:login/:password

	// Set variable values

}

func getListService(phone string) {
	//GET
	//config/products

	// Set variable values

}

func getListServiceAndCategory(phone string) {
	//GET
	//config/products/simple

	// Set variable values

}

func selectMenuService(phone string) {
	//GET
	//config/menus

	// Set variable values

}

func selectMenuServiceId(phone string) {
	//tree service
	//GET
	//config/menus/:menuId

	// Set variable values

}

func returnTimeService(phone string) {
	//GET
	// queue/awaiting

	// Set variable values

}

//GET booking/menu
//Возвращает дерево меню выбора услуг для предварительной записи

//GET booking/schedule/days
//Возвращает дни, в которые доступна предварительная запись для указанной услуги

//GET booking/schedule/days/:date
//Возвращает список временных точек, доступных для предварительной записи в указанную дату

//POST booking/login/:userIdentity
//Возвращает токен авторизации. В качестве
//идентификатора пользователя может быть передана
//любая строка. Данный метод аутентификации может
//использоваться только с доверенного адреса,
//который указывается в настройках очереди.

//POST booking/book/:date/:time
//Предварительная запись на получение услуги. Возвращает элемент с данными о записи,
//включающими идентификатор талона для дальнейшей работы.

//GET booking/tickets
//Возвращает список талонов предварительной записи аутентифицированного пользователя

//GET booking/tickets/:ticketId
//Получение информации о талоне предварительной записи

//POST booking/tickets/:ticketId/confirm
//Подтверждение предварительной записи. Требуется, если параметр confirm в методе booking не был передан

//POST booking/tickets/:ticketId/cancel
//Отмена предварительной записи.

//GET /tickets/ticketpng/:ticketId/
//Получение талончика по предварительной записи в печатном виде
