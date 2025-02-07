package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

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

	// 1 - user sends the location. We need to send him the, send full name reply.
	// 2 - user sends the full name. We send him, thanks! Process complete msg. 

	var req twilioLocationReq

    if err := ctx.ShouldBind(&req); err != nil { 
        ctx.JSON(http.StatusBadRequest, errorResponse(err))
        return
    }

	Latitude := 0.0
	Longitude := 0.0
	Name := ""
	Status := ""

	if req.Latitude != 0.0 && req.Longitude != 0.0 {
		Latitude = req.Latitude
		Longitude = req.Longitude
	}

	// Assuming he sends only the right value in the body. 
	
	if (req.Body != "") {
		Name = req.Body
	} else {
		Name = req.ProfileName
	}
	fmt.Println("NAME: ", Name)

	foundUser, err := server.store.GetUserByNumber(ctx, sql.NullString{String: strings.Split(req.From, ":")[1], Valid: true})
	
	if err != nil {
		if err == sql.ErrNoRows {
			arg := db.CreateUserParams{
				FullName: sql.NullString{String: Name, Valid: true},
				Latitude: Latitude,
				Longitude: Longitude,
				PhoneNumber: sql.NullString{String: strings.Split(req.From, ":")[1], Valid: true},
				Status: db.NullUserStatus{UserStatus: "ADD_LOCATION", Valid: true},
			}
			
			_, err := server.store.CreateUser(ctx, arg)
			
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return 
			} 
			Status = "ADD_LOCATION"
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return 
		}
	} else {
		if Latitude != 0.0 {
			_, err = server.store.UpdateUser(ctx, db.UpdateUserParams{
				ID: foundUser.ID,
				FullName: sql.NullString{String: Name, Valid: true},
				Status: db.NullUserStatus{UserStatus: "ADD_LOCATION", Valid: true},
				Latitude: Latitude,
				Longitude: Longitude,
			})
			Status = "ADD_LOCATION"
		} else {
			_, err = server.store.UpdateUserName(ctx, db.UpdateUserNameParams{
				ID: foundUser.ID,
				FullName: sql.NullString{String: Name, Valid: true},
				Status: db.NullUserStatus{UserStatus: "ADD_NAME", Valid: true},
			})
			Status = "ADD_NAME"
		}
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return 
		}
	}

	msgBody := ""
	switch s := Status; s {
	case "ADD_NAME":
		msgBody = util.MessageResponsesInstance.ProcessComplete
	case "ADD_LOCATION":
		msgBody = util.MessageResponsesInstance.LocationReceived
	default:
		msgBody = util.MessageResponsesInstance.SendLocationPrompt
	}

	response, err := util.TwilioSendMsg(req.From, msgBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response) // Twilio expects a 200 OK response
}

func (server *Server) GetUsers(ctx *gin.Context) {
	users, err := server.store.ListUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Convert DB users to the API response format that a frontend can work with.
	var responseUsers []UserResponse
	for _, user := range users {
		responseUsers = append(responseUsers, newUserResponse(user))
	}

	ctx.JSON(http.StatusOK, responseUsers)
}


type UserResponse struct {
	ID          int64     `json:"id"`
	FullName    *string   `json:"full_name,omitempty"`
	PhoneNumber *string   `json:"phone_number,omitempty"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Address     *string   `json:"address,omitempty"`
	IsFamily    *bool     `json:"is_family,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	Status      string    `json:"status,omitempty"`
}

func newUserResponse(user db.User) UserResponse {
	return UserResponse{
		ID:          user.ID,
		FullName:    nullStringToPtr(user.FullName),
		PhoneNumber: nullStringToPtr(user.PhoneNumber),
		Latitude:    user.Latitude,
		Longitude:   user.Longitude,
		Address:     nullStringToPtr(user.Address),
		IsFamily:    nullBoolToPtr(user.IsFamily),
		CreatedAt:   user.CreatedAt,
		Status:      string(user.Status.UserStatus), // Assuming NullUserStatus has a `.String` field
	}
}

// Helper functions to handle null values
func nullStringToPtr(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

func nullBoolToPtr(nb sql.NullBool) *bool {
	if nb.Valid {
		return &nb.Bool
	}
	return nil
}


