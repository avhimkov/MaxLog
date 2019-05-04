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

func maxNewTicket(ip string) {
	// Set variable values
	gatewayURL := "http://" + ip + ":5612/QMSSR"
	urlStr := gatewayURL

	// Params
	v := url.Values{}
	v.Set("ip", ip)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)

}
