package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Login    string
	Password string
}

func main() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Panic(err)
	}

	configuration.Login
	configuration.Password

	if err != nil {
		log.Panic(err)
	}

	// sms
	tele2("79827468271")
	// push

	/* SendGMToClient("Hello from GCM", "<CLIENT TOKEN>") */
}
