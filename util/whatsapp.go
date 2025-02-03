package util

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func TwilioSendMsg(toNumber string, msgBody string) (response string, err error) {

	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: config.AccountSid,
		Password: config.AuthToken,
	})

	params := &api.CreateMessageParams{}
	params.SetTo(toNumber)
	params.SetFrom(fmt.Sprintf("whatsapp:%s", config.SenderNumber))
	params.SetBody(msgBody)
	// params.SetStatusCallback("https://8ce9-2401-4900-1c43-f672-4517-c91c-90d3-d4eb.ngrok-free.app/twilio-status")


	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending Whatsapp message: " + err.Error())
		return "", err
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
		return string(response), nil
	}
}