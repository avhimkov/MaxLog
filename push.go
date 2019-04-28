package main

import (
	"fmt"

	"github.com/google/go-gcm"
)

func SendGCMToClient(pushText string, pushToken string) {
	serverKey := "<YOUR SERVER KEY>"
	var msg gcm.HttpMessage
	data := map[string]interface{}{"message": pushText}
	regIDs := []string{pushToken}
	msg.RegistrationIds = regIDs
	msg.Data = data
	response, err := gcm.SendHttp(serverKey, msg)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Response ", response.Success)
		fmt.Println("MessageID ", response.MessageId)
		fmt.Println("Failure ", response.Failure)
		fmt.Println("Error ", response.Error)
		fmt.Println("Results ", response.Results)
	}
}
