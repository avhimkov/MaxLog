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

func sms() (string, string) {
	fileSMS, _ := os.Open("./config/configsms.json")
	decoderSMS := json.NewDecoder(fileSMS)
	configurationSMS := ConfigSMS{}
	errSMS := decoderSMS.Decode(&configurationSMS)
	if errSMS != nil {
		log.Panic(errSMS)
	}
	login := configurationSMS.Login
	pass := configurationSMS.Password

	return login, pass
}

func tele2(phone string, login string, password string, sender string, text string) {
	// Set variable values
	gatewayURL := "http://bsms-proxy.tele2.ru/api/"

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

func megafon(phone string, login string, password string, sender string, text string) {
	// Set variable values
	gatewayURL := "https://a2p-api.megalabs.ru/"

	//В заголовке передаются авторизационные данные в формате login:pass закодированные Base64. Например, cG8nbX54dGAyqx==.
	urlStr := gatewayURL + "/sms/v1/sms" +
		"&login=" + login +
		"&password=" + password +
		"&shortcode=" + sender +
		"&text=" + text

	/* "from": "Johnny_Viper",
	"to": 79001234212,
	"message": "Это тест",
	"callback_url": "http://www.sombodyserver.ru/sms_callback" */

	// Params
	v := url.Values{}
	v.Set("to", phone)
	v.Set("login", login)
	v.Set("password", password)
	v.Set("sender", sender)
	v.Set("message", text)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)

	//print response
	fmt.Println(resp.Status)

}

func mts(phone string, login string, password string, sender string, text string) {

	/* HTTP GET
	   The following is a sample HTTP GET request and response. The placeholders shown need to be replaced with actual values.

	   GET /m2m/m2m_api.asmx/SendMessage?msid=string&message=string&naming=string&login=string&password=string HTTP/1.1
	   Host: mcommunicator.ru
	   HTTP/1.1 200 OK
	   Content-Type: text/xml; charset=utf-8
	   Content-Length: length

	   <?xml version="1.0" encoding="utf-8"?>
	   <long xmlns="http://mcommunicator.ru/M2M">long</long>
	   HTTP POST
	   The following is a sample HTTP POST request and response. The placeholders shown need to be replaced with actual values.

	   POST /m2m/m2m_api.asmx/SendMessage HTTP/1.1
	   Host: mcommunicator.ru
	   Content-Type: application/x-www-form-urlencoded
	   Content-Length: length

	   msid=string&message=string&naming=string&login=string&password=string
	   HTTP/1.1 200 OK
	   Content-Type: text/xml; charset=utf-8
	   Content-Length: length

	   <?xml version="1.0" encoding="utf-8"?>
	   <long xmlns="http://mcommunicator.ru/M2M">long</long> */

	// Set variable values
	gatewayURL := "http://mcommunicator.ru/"

	urlStr := gatewayURL + "/m2m_api.asmx/SendMessage?" +
		"msid=" + phone +
		"&message=" + text +
		"&naming=" + sender +
		"&login=" + login +
		"&password=" + password

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
