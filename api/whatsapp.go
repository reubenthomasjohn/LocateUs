package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reubenthomasjohn/location-heatmap/util"
)

type sendMsgToNumReq struct {
	ToNumber string `json:"to_number" binding:"required"`
}

// func SendWhatsAppMsg(ctx *gin.Context) {
// 	var req sendMsgToNumReq
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}
// 	body, err := util.SendWhatsAppMessage(req.ToNumber)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message": body,
// 	})
// }

func TwilioStatusMsg(ctx *gin.Context) {
	messageSid := ctx.PostForm("MessageSid")
	messageStatus := ctx.PostForm("MessageStatus")
	errorCode := ctx.PostForm("ErrorCode")
	eventType := ctx.PostForm("EventType")

	fmt.Printf("MessageSid: %s, MessageStatus: %s, ErrorCode: %s, EventType: %s ", messageSid, messageStatus, errorCode, eventType)

	ctx.JSON(http.StatusOK, nil)
}


type twilioLocationReq struct {
    Latitude  string `form:"Latitude"`   
    Longitude string `form:"Longitude"`
    Address   string `form:"Address"`
    Label     string `form:"Label"`  
	From      string `form:"From"`    
	Body      string `form:"Body"`
}

func TwilioReceiveMsg(ctx *gin.Context) {

	var req twilioLocationReq

    if err := ctx.ShouldBind(&req); err != nil { 
        ctx.JSON(http.StatusBadRequest, errorResponse(err))
        return
    }

    fmt.Printf("Latitude: %s, Longitude: %s\n", req.Latitude, req.Longitude)
    fmt.Printf("Address: %s, Label: %s, From: %s\n", req.Address, req.Label, req.From)

	fromNumber := req.From
	msgBody := "Please share your location (not live location)."
	if req.Latitude != "" && req.Longitude != "" {
		msgBody = "Thank you for sharing your location. Please visit https://www.google.com to view the heatmap"
	}

	response, err := util.TwilioSendMsg(fromNumber, msgBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response) // Twilio expects a 200 OK response
}

