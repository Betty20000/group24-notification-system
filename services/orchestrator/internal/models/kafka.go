package models

import "time"

// KafkaNotificationPayload is the message structure sent to Kafka
// This is what the Email and Push services will consume
type KafkaNotificationPayload struct {
	NotificationID   string                 `json:"notification_id"`
	UserID           string                 `json:"user_id"`
	NotificationType string                 `json:"notification_type"` // "email" or "push"
	To               string                 `json:"to"`                // email address or device token
	Subject          string                 `json:"subject,omitempty"` // for email
	Body             string                 `json:"body"`
	Priority         string                 `json:"priority"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt        time.Time              `json:"created_at"`

	// For retry logic (will be updated by consumer services)
	RetryCount  int       `json:"retry_count,omitempty"`
	LastRetryAt time.Time `json:"last_retry_at,omitempty"`
}
