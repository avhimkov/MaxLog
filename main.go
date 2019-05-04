package main

import (
	"encoding/json"
	"log"
	"os"
)

type ConfigSMS struct {
	Login    string
	Password string
}

type ConfigMax struct {
	IP   string
	Port string
}

func main() {
	// SMS
	fileSMS, _ := os.Open("configsms.json")
	decoderSMS := json.NewDecoder(fileSMS)
	configurationSMS := ConfigSMS{}
	errSMS := decoderSMS.Decode(&configurationSMS)
	if errSMS != nil {
		log.Panic(errSMS)
	}

	// MAXIMA
	fileMax, _ := os.Open("configMax.json")
	decoderMax := json.NewDecoder(fileMax)
	configurationMax := ConfigMax{}
	errMax := decoderMax.Decode(&configurationMax)
	if errMax != nil {
		log.Panic(errMax)
	}

	// SMS
	login := configurationSMS.Login
	pass := configurationSMS.Password

	// MAXIMA
	ip := configurationMax.IP

	maxNewTicket(ip)

	// sms
	tele2("79827468271", login, pass, "MFC", "Hello")

	// push
	SendGCMToClient("Hello from GCM", "<CLIENT TOKEN>")
}
