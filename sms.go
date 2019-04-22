package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func tele2(phone string) {
	// Set variable values
	gatewayURL := "http://bsms-proxy.tele2.ru/api/"
	login := "AUTH_API_KEY"
	password := "PASS"
	sender := "SEDEMO"
	urlStr := gatewayURL + "api/?operation=send" +
		"&login=" + "someLogin" +
		"&password=" + "somePassword" +
		"&msisdn=" + "79111234567" +
		"&shortcode=" + "SomeSender" +
		"&text=" + "TestMessage"

	// Params
	v := url.Values{}
	v.Set("to", phone)
	v.Set("login", login)
	v.Set("password", password)
	v.Set("sender", sender)
	v.Set("message", "Hello, this is a test sms")
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("GET", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)

}
