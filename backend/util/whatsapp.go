package util

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func TwilioSendMsg(toNumber, msgBody string) (response string, err error) {

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
	// params.SetMediaUrl([]string{"https://media.christcommunitychurch.in/sites/2/2017/06/2017-logo-2.png"})
	// params.SetStatusCallback(fmt.Sprintf("%s/%s", config.PrefixUrl, callbackEndpoint))


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