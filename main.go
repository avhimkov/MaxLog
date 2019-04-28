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
	fileSMS, _ := os.Open("configsms.json")
	decoderSMS := json.NewDecoder(fileSMS)
	configurationSMS := ConfigSMS{}
	errSMS := decoderSMS.Decode(&configurationSMS)
	if errSMS != nil {
		log.Panic(errSMS)
	}

	fileMax, _ := os.Open("configMax.json")
	decoderMax := json.NewDecoder(fileMax)
	configurationMax := ConfigMax{}
	errMax := decoderMax.Decode(&configurationMax)
	if errMax != nil {
		log.Panic(errMax)
	}

	// configurationSMS.Login
	// configurationSMS.Password

	// configurationMax.IP
	// configurationMax.Port

	// if err != nil {
	// 	log.Panic(err)
	// }

	// sms
	tele2("79827468271")
	// push

	/* SendGMToClient("Hello from GCM", "<CLIENT TOKEN>") */
}
