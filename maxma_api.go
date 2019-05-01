package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

//http://" + location.hostname + ":5612/QMSSR

/**
 * Асинхронная регистрация талона
 * @param {{serviceId, languageId, custData, noteText, priorityId, userId}} params
 * @returns promise
 */

/** Асинхронно получить информацию по пину
 * [registerByPIN description]
 * @param {number}   pin      [description]
 * @param {[type]}   attr     [description]
 * @param {Function} callback [description]
 */

func newTicket(ip string, port string) {
	// Set variable values
	gatewayURL := "queue/admin/tickets/new"
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

func serviceInTicket(ip string, port string) {
	// Set variable values
	gatewayURL := "queue/admin/tickets/:ticketId/addProduct/:productId"
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

func cancelServiceTiket(ip string, port string) {
	// Set variable values
	gatewayURL := "queue/admin/tickets/:ticketId/:ticketProductId/cancel"
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

func prioritetServiceTiket(ip string, port string) {
	// Set variable values
	gatewayURL := "queue/admin/tickets/:ticketId/:ticketProductId/setPriority"
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

func returnOpertorList(ip string, port string) {
	// Set variable values
	gatewayURL := "queue/admin/redirect/users/:productId"
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

func returnDesktop(ip string, port string) {
	// Set variable values
	gatewayURL := "queue/admin/redirect/workPlaces/:productId"
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

func redirectUser(ip string, port string) {
	// Set variable values
	gatewayURL := "queue/admin/redirect/:productId"
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

func getInfoTiket(ip string, port string) {
	// Set variable values
	gatewayURL := "queue/admin/tickets/:ticketId"
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

func returnAllTiketsPredList(ip string, port string) {
	// Set variable values
	gatewayURL := "queue/admin/booking"
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

func returnAllTikets(ip string, port string) {
	// Set variable values
	gatewayURL := "queue/admin/tickets"
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

func returnTokenAuth(ip string, port string) {
	// Set variable values
	gatewayURL := "queue/admin/login/:login/:password"
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

func getListService(ip string, port string) {
	// Set variable values
	gatewayURL := "config/products"
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

func getListServiceAndCategory(ip string, port string) {
	// Set variable values
	gatewayURL := "config/products/simple"
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

func selectMenuService(ip string, port string) {
	// Set variable values
	gatewayURL := "config/menus"
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

func selectMenuServiceId(ip string, port string) {
	//tree service
	// Set variable values
	gatewayURL := "config/menus/:menuId"
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

func returnTimeService(ip string, port string) {
	//GET
	// queue/awaiting

	// Set variable values
	gatewayURL := "queue/awaiting"
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

//Возвращает дерево меню выбора услуг для предварительной записи
func returnMenuServicePre(ip string, port string) {
	// Set variable values
	gatewayURL := "booking/menu"
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

//Возвращает дни, в которые доступна предварительная запись для указанной услуги
func returnDayPre(ip string, port string) {
	// Set variable values
	gatewayURL := "booking/schedule/days"
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

//Возвращает список временных точек, доступных для предварительной записи в указанную дату
func returnListTimePreService(ip string, port string) {
	// Set variable values
	gatewayURL := "booking/schedule/days/:date"
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

//Возвращает токен авторизации. В качестве
//идентификатора пользователя может быть передана
//любая строка. Данный метод аутентификации может
//использоваться только с доверенного адреса,
//который указывается в настройках очереди.
func returnTokeAuth(ip string, port string) {
	// Set variable values
	gatewayURL := "booking/login/:userIdentity"
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

//Предварительная запись на получение услуги. Возвращает элемент с данными о записи,
//включающими идентификатор талона для дальнейшей работы.
func returnIdTiket(ip string, port string) {
	// Set variable values
	gatewayURL := "booking/book/:date/:time"
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

//Возвращает список талонов предварительной записи аутентифицированного пользователя
func returnListTicketAuthUser(ip string, port string) {
	// Set variable values
	gatewayURL := "booking/tickets"
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

//Получение информации о талоне предварительной записи
func infoPreTicket(ip string, port string) {
	// Set variable values
	gatewayURL := "booking/tickets/:ticketId"
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

//Подтверждение предварительной записи. Требуется, если параметр confirm в методе booking не был передан
func confurmPre(ip string, port string) {
	// Set variable values
	gatewayURL := "booking/tickets/:ticketId/confirm"
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

//Отмена предварительной записи.
func cancelPre(ip string, port string) {
	// Set variable values
	gatewayURL := "booking/tickets/:ticketId/cancel"
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

//Получение талончика по предварительной записи в печатном виде
func PrintPreTicket(ip string, port string) {
	// Set variable values
	gatewayURL := "/tickets/ticketpng/:ticketId/"
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
