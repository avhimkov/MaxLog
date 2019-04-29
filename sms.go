package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func tele2(phone string, login string, password string, sender string, text string) {
	// Set variable values
	gatewayURL := "http://bsms-proxy.tele2.ru/api/"
	// login := "AUTH_API_KEY"
	// password := "PASS"
	// sender := "SEDEMO"

	urlStr := gatewayURL + "api/?operation=send" +
		"&login=" + login +
		"&password=" + password +
		"&msisdn=" + phone +
		"&shortcode=" + sender +
		"&text=" + text

	// Params
	v := url.Values{}
	v.Set("to", phone)
	v.Set("login", login)
	v.Set("password", password)
	v.Set("sender", sender)
	v.Set("message", text)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("GET", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)

}
