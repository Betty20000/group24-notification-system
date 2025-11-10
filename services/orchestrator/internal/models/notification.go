package models

import "time"

type NotificationRequest struct {
	UserID       string                 `json:"user_id" binding:"required"`
	TemplateID   string                 `json:"template_id" binding:"required"`
	Channel      string                 `json:"channel" binding:"required,oneof=email push sms"`
	Variables    map[string]interface{} `json:"variables"`
	Priority     string                 `json:"priority,omitempty"` // high, normal, low
	ScheduledFor *time.Time             `json:"scheduled_for,omitempty"`
}

type NotificationResponse struct {
	NotificationID string    `json:"notification_id"`
	Status         string    `json:"status"`
	Message        string    `json:"message,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
}

type Notification struct {
	ID         string                 `json:"id"`
	UserID     string                 `json:"user_id"`
	TemplateID string                 `json:"template_id"`
	Channel    string                 `json:"channel"`
	To         string                 `json:"to"` // email address or device token
	Subject    string                 `json:"subject,omitempty"`
	Body       string                 `json:"body"`
	Variables  map[string]interface{} `json:"variables"`
	Priority   string                 `json:"priority"`
	Status     string                 `json:"status"`
	CreatedAt  time.Time              `json:"created_at"`
}
