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

func main() {
	// SMS
	fileSMS, _ := os.Open("configsms.json")
	decoderSMS := json.NewDecoder(fileSMS)
	configurationSMS := ConfigSMS{}
	errSMS := decoderSMS.Decode(&configurationSMS)
	if errSMS != nil {
		log.Panic(errSMS)
	}

	// SMS
	login := configurationSMS.Login
	pass := configurationSMS.Password

	// sms
	tele2("79827468271", login, pass, "MFC", "Hello")

	// push
	SendGCMToClient("Hello from GCM", "<CLIENT TOKEN>")
}
