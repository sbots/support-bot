package viber

type SubscriptionRequest struct {
	Url        string     `json:"url"`
	EventTypes EventTypes `json:"event_types"`
	SendName   bool       `json:"send_name"`
	SendPhoto  bool       `json:"send_photo"`
}

type SubscriptionResponse struct {
	Status        int    `json:"status"`
	StatusMessage string `json:"status_message"`
}

type EventTypes []string

var AllEventsTypes = EventTypes{
	EventTypeDelivered,
	EventTypeSeen,
	EventTypeFailed,
	EventTypeSubscribed,
	EventTypeUnsubscribed,
	EventTypeConversationStarted,
}

const (
	EventTypeDelivered           = "delivered"
	EventTypeSeen                = "seen"
	EventTypeFailed              = "failed"
	EventTypeSubscribed          = "subscribed"
	EventTypeUnsubscribed        = "unsubscribed"
	EventTypeConversationStarted = "conversation_started"
)
