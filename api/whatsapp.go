package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	db "github.com/reubenthomasjohn/location-heatmap/db/sqlc"
	"github.com/reubenthomasjohn/location-heatmap/util"
)


type TwilioStatusReq struct {
	MessageSid string `form:"MessageSid"`
	AccountSid string `form:"AccountSid"`
	From string `form:"From"`
	To string `form:"To"`
	Body string `form:"Body"`
	MessageStatus string `form:"MessageStatus"`
	ErrorCode string `form:"ErrorCode"`
	EventType string `form:"EventType"`
}

func (server *Server) TwilioStatusMsg(ctx *gin.Context) {
	var req TwilioStatusReq

    if err := ctx.ShouldBind(&req); err != nil { 
        ctx.JSON(http.StatusBadRequest, errorResponse(err))
        return
    }

	fmt.Printf("MessageSid: %s, MessageStatus: %s, ErrorCode: %s, EventType: %s ", req.MessageSid, req.MessageStatus, req.ErrorCode, req.EventType)

	ctx.JSON(http.StatusOK, nil)
}

type twilioLocationReq struct {
    Latitude  float64 `form:"Latitude"`   
    Longitude float64 `form:"Longitude"`
    Address   string `form:"Address"`
    Label     string `form:"Label"`  
	From      string `form:"From"`    
	Body      string `form:"Body"`
	ProfileName	string `form:"ProfileName"`
}

func (server *Server) TwilioReceiveMsg(ctx *gin.Context) {

	var req twilioLocationReq

    if err := ctx.ShouldBind(&req); err != nil { 
        ctx.JSON(http.StatusBadRequest, errorResponse(err))
        return
    }

	msgBody := util.MessageResponsesInstance.SendLocationPrompt
	if req.Latitude != 0.0 && req.Longitude != 0.0 {
		msgBody = util.MessageResponsesInstance.LocationReceived

		foundUser, err := server.store.GetUserByNumber(ctx, sql.NullString{String: strings.Split(req.From, ":")[1], Valid: true})

		if err != nil {
			if err == sql.ErrNoRows {
				arg := db.CreateUserParams{
					FullName: sql.NullString{String: req.ProfileName, Valid: true},
					Latitude: req.Latitude,
					Longitude: req.Longitude,
					PhoneNumber: sql.NullString{String: strings.Split(req.From, ":")[1], Valid: true},
				}
				
				_, err := server.store.CreateUser(ctx, arg)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, errorResponse(err))
					return 
				}
		
			} else {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return 
			}
		}
		_, err = server.store.UpdateUser(ctx, db.UpdateUserParams{
			ID: foundUser.ID,
			Latitude: req.Latitude,
			Longitude: req.Longitude,
			FullName: sql.NullString{String: req.ProfileName, Valid: true},
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return 
		}
	}

	response, err := util.TwilioSendMsg(req.From, msgBody, "twilio-status")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response) // Twilio expects a 200 OK response
}

