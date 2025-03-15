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


type twilioStatusRequest struct {
	MessageSid string `form:"MessageSid"`
	AccountSid string `form:"AccountSid"`
	From string `form:"From"`
	To string `form:"To"`
	Body string `form:"Body"`
	MessageStatus string `form:"MessageStatus"`
	ErrorCode string `form:"ErrorCode"`
	EventType string `form:"EventType"`
}

func (server *Server) twilioStatusMsg(ctx *gin.Context) {
	var req twilioStatusRequest

    if err := ctx.ShouldBind(&req); err != nil { 
        ctx.JSON(http.StatusBadRequest, errorResponse(err))
        return
    }

	fmt.Printf("MessageSid: %s, MessageStatus: %s, ErrorCode: %s, EventType: %s ", req.MessageSid, req.MessageStatus, req.ErrorCode, req.EventType)

	ctx.JSON(http.StatusOK, nil)
}

type twilioLocationRequest struct {
    Latitude  float64 `form:"Latitude"`   
    Longitude float64 `form:"Longitude"`
    Address   string `form:"Address"`
    Label     string `form:"Label"`  
	From      string `form:"From"`    
	Body      string `form:"Body"`
	ProfileName	string `form:"ProfileName"`
}

func (server *Server) twilioReceiveMsg(ctx *gin.Context) {

	var req twilioLocationRequest

    if err := ctx.ShouldBind(&req); err != nil { 
        ctx.JSON(http.StatusBadRequest, errorResponse(err))
        return
    }

	Latitude := 0.0
	Longitude := 0.0
	Name := ""
	Status := ""

	// if req.Longitude == 0.0 && req.Latitude == 0.0 {
	// 	err := errors.New("empty location data (0.0) found in twilio location message")
	// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
    //     return
	// }

	Latitude = req.Latitude
	Longitude = req.Longitude
	// Assuming he sends only the right value in the body. 
	
	if (req.Body != "") {
		Name = req.Body
	} else {
		Name = req.ProfileName
	}
	fmt.Println("NAME: ", Name)

	foundUser, err := server.store.GetMemberByNumber(ctx, sql.NullString{String: strings.Split(req.From, ":")[1], Valid: true})
	
	if err != nil {
		if err == sql.ErrNoRows {
			arg := db.CreateMemberParams{
				FullName: sql.NullString{String: Name, Valid: true},
				Latitude: Latitude,
				Longitude: Longitude,
				PhoneNumber: sql.NullString{String: strings.Split(req.From, ":")[1], Valid: true},
				Status: db.NullUserStatus{UserStatus: "ADD_LOCATION", Valid: true},
			}
			
			_, err := server.store.CreateMember(ctx, arg)
			
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
			_, err = server.store.UpdateMember(ctx, db.UpdateMemberParams{
				ID: foundUser.ID,
				FullName: sql.NullString{String: Name, Valid: true},
				Status: db.NullUserStatus{UserStatus: "ADD_LOCATION", Valid: true},
				Latitude: Latitude,
				Longitude: Longitude,
			})
			Status = "ADD_LOCATION"
		} else {
			_, err = server.store.UpdateMemberName(ctx, db.UpdateMemberNameParams{
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

func (server *Server) getMembers(ctx *gin.Context) {
	users, err := server.store.ListMembers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Convert DB users to the API response format that a frontend can work with.
	var responseUsers []MemberResponse
	for _, user := range users {
		responseUsers = append(responseUsers, newMemberResponse(user))
	}

	ctx.JSON(http.StatusOK, responseUsers)
}


type MemberResponse struct {
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

func newMemberResponse(member db.Member) MemberResponse {
	return MemberResponse{
		ID:          member.ID,
		FullName:    nullStringToPtr(member.FullName),
		PhoneNumber: nullStringToPtr(member.PhoneNumber),
		Latitude:    member.Latitude,
		Longitude:   member.Longitude,
		Address:     nullStringToPtr(member.Address),
		IsFamily:    nullBoolToPtr(member.IsFamily),
		CreatedAt:   member.CreatedAt,
		Status:      string(member.Status.UserStatus), // Assuming NullUserStatus has a `.String` field
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


