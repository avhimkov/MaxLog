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
