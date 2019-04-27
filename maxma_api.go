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
	gatewayURL := "queue/admin/tickets/:ticketId/:ticketProductId/setPriority"
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

func returnOpertorList(phone string) {
	//POST
	//queue/admin/redirect/users/:productId

	// Set variable values
	gatewayURL := "queue/admin/redirect/users/:productId"
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

func returnDesktop(phone string) {
	//POST
	//queue/admin/redirect/workPlaces/:productId

	// Set variable values
	gatewayURL := "queue/admin/redirect/workPlaces/:productId"
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

func redirectUser(phone string) {
	//POST
	// queue/admin/redirect/:productId

	// Set variable values
	gatewayURL := "queue/admin/redirect/:productId"
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

func getInfoTiket(phone string) {
	//GET
	//queue/admin/tickets/:ticketId

	// Set variable values
	gatewayURL := "queue/admin/tickets/:ticketId"
	ip := "192.168.10.176"
	port := "5600"
	urlStr := ip + port + gatewayURL

	// Params
	v := url.Values{}
	v.Set("ip", ip)
	v.Set("port", port)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("GET", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)

}

func returnAllTiketsPredList(phone string) {
	//GET
	//queue/admin/booking

	// Set variable values
	gatewayURL := "queue/admin/booking"
	ip := "192.168.10.176"
	port := "5600"
	urlStr := ip + port + gatewayURL

	// Params
	v := url.Values{}
	v.Set("ip", ip)
	v.Set("port", port)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("GET", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)

}

func returnAllTikets(phone string) {
	//GET
	//queue/admin/tickets

	// Set variable values
	gatewayURL := "queue/admin/tickets"
	ip := "192.168.10.176"
	port := "5600"
	urlStr := ip + port + gatewayURL

	// Params
	v := url.Values{}
	v.Set("ip", ip)
	v.Set("port", port)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("GET", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)

}

func returnTokenAuth(phone string) {
	//POST
	//queue/admin/login/:login/:password

	// Set variable values
	gatewayURL := "queue/admin/login/:login/:password"
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

func getListService(phone string) {
	//GET
	//config/products

	// Set variable values
	gatewayURL := "config/products"
	ip := "192.168.10.176"
	port := "5600"
	urlStr := ip + port + gatewayURL

	// Params
	v := url.Values{}
	v.Set("ip", ip)
	v.Set("port", port)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("GET", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)

}

func getListServiceAndCategory(phone string) {
	//GET
	//config/products/simple

	// Set variable values
	gatewayURL := "config/products/simple"
	ip := "192.168.10.176"
	port := "5600"
	urlStr := ip + port + gatewayURL

	// Params
	v := url.Values{}
	v.Set("ip", ip)
	v.Set("port", port)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("GET", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)

}

func selectMenuService(phone string) {
	//GET
	//config/menus

	// Set variable values
	gatewayURL := "config/menus"
	ip := "192.168.10.176"
	port := "5600"
	urlStr := ip + port + gatewayURL

	// Params
	v := url.Values{}
	v.Set("ip", ip)
	v.Set("port", port)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("GET", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)

}

func selectMenuServiceId(phone string) {
	//tree service
	//GET
	//config/menus/:menuId

	// Set variable values
	gatewayURL := "config/menus/:menuId"
	ip := "192.168.10.176"
	port := "5600"
	urlStr := ip + port + gatewayURL

	// Params
	v := url.Values{}
	v.Set("ip", ip)
	v.Set("port", port)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("GET", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)

}

func returnTimeService(phone string) {
	//GET
	// queue/awaiting

	// Set variable values
	gatewayURL := "queue/awaiting"
	ip := "192.168.10.176"
	port := "5600"
	urlStr := ip + port + gatewayURL

	// Params
	v := url.Values{}
	v.Set("ip", ip)
	v.Set("port", port)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("GET", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)

}

//GET booking/menu
//Возвращает дерево меню выбора услуг для предварительной записи
func returnTimeService(phone string) {
	// Set variable values
	gatewayURL := "booking/menu"
	ip := "192.168.10.176"
	port := "5600"
	urlStr := ip + port + gatewayURL

	// Params
	v := url.Values{}
	v.Set("ip", ip)
	v.Set("port", port)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("GET", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)

}

//GET booking/schedule/days
//Возвращает дни, в которые доступна предварительная запись для указанной услуги
func returnTimeService(phone string) {
	// Set variable values
	gatewayURL := "booking/schedule/days"
	ip := "192.168.10.176"
	port := "5600"
	urlStr := ip + port + gatewayURL

	// Params
	v := url.Values{}
	v.Set("ip", ip)
	v.Set("port", port)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("GET", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)

}

//GET booking/schedule/days/:date
//Возвращает список временных точек, доступных для предварительной записи в указанную дату
func returnTimeService(phone string) {
	// Set variable values
	gatewayURL := "booking/schedule/days/:date"
	ip := "192.168.10.176"
	port := "5600"
	urlStr := ip + port + gatewayURL

	// Params
	v := url.Values{}
	v.Set("ip", ip)
	v.Set("port", port)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("GET", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)

}

//POST booking/login/:userIdentity
//Возвращает токен авторизации. В качестве
//идентификатора пользователя может быть передана
//любая строка. Данный метод аутентификации может
//использоваться только с доверенного адреса,
//который указывается в настройках очереди.
func returnTimeService(phone string) {
	// Set variable values
	gatewayURL := "booking/login/:userIdentity"
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

//POST booking/book/:date/:time
//Предварительная запись на получение услуги. Возвращает элемент с данными о записи,
//включающими идентификатор талона для дальнейшей работы.
func returnTimeService(phone string) {
	// Set variable values
	gatewayURL := "booking/book/:date/:time"
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

//GET booking/tickets
//Возвращает список талонов предварительной записи аутентифицированного пользователя
func returnTimeService(phone string) {
	// Set variable values
	gatewayURL := "booking/tickets"
	ip := "192.168.10.176"
	port := "5600"
	urlStr := ip + port + gatewayURL

	// Params
	v := url.Values{}
	v.Set("ip", ip)
	v.Set("port", port)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("GET", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)
}

//GET booking/tickets/:ticketId
//Получение информации о талоне предварительной записи
func returnTimeService(phone string) {
	// Set variable values
	gatewayURL := "booking/tickets/:ticketId"
	ip := "192.168.10.176"
	port := "5600"
	urlStr := ip + port + gatewayURL

	// Params
	v := url.Values{}
	v.Set("ip", ip)
	v.Set("port", port)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("GET", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)
}

//POST booking/tickets/:ticketId/confirm
//Подтверждение предварительной записи. Требуется, если параметр confirm в методе booking не был передан
func returnTimeService(phone string) {
	// Set variable values
	gatewayURL := "booking/tickets/:ticketId/confirm"
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

//POST booking/tickets/:ticketId/cancel
//Отмена предварительной записи.
func returnTimeService(phone string) {
	// Set variable values
	gatewayURL := "booking/tickets/:ticketId/cancel"
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

//GET /tickets/ticketpng/:ticketId/
//Получение талончика по предварительной записи в печатном виде
func returnTimeService(phone string) {
	// Set variable values
	gatewayURL := "/tickets/ticketpng/:ticketId/"
	ip := "192.168.10.176"
	port := "5600"
	urlStr := ip + port + gatewayURL

	// Params
	v := url.Values{}
	v.Set("ip", ip)
	v.Set("port", port)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("GET", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)
}
