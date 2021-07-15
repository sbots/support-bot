package models

type ViberSubscriptionRequest struct {
	Url        string          `json:"url"`
	EventTypes ViberEventTypes `json:"event_types"`
	SendName   bool            `json:"send_name"`
	SendPhoto  bool            `json:"send_photo"`
}

type ViberSubscriptionResponse struct {
	Status        int    `json:"status"`
	StatusMessage string `json:"status_message"`
}

type ViberEventTypes []string

var AllViberEventTypes = ViberEventTypes{
	ViberEventTypeDelivered,
	ViberEventTypeSeen,
	ViberEventTypeFailed,
	ViberEventTypeSubscribed,
	ViberEventTypeUnsubscribed,
	ViberEventTypeConversationStarted,
}

const (
	ViberEventTypeDelivered           = "delivered"
	ViberEventTypeSeen                = "seen"
	ViberEventTypeFailed              = "failed"
	ViberEventTypeSubscribed          = "subscribed"
	ViberEventTypeUnsubscribed        = "unsubscribed"
	ViberEventTypeConversationStarted = "conversation_started"
)
