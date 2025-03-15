package util

type MessageResponses struct {
	LocationReceived string
	SendLocationPrompt string
	ProcessComplete string
}

var MessageResponsesInstance = MessageResponses{
	LocationReceived: "Thank you for sharing your location. \n\nNow, please send your full name. \n(ex: John Doe)",
	SendLocationPrompt: "Please share your location (not live location)",
	ProcessComplete: "We have received all your details. Thank you!",
}

