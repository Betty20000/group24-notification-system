package models

import "time"

type UserPreferences struct {
	UserID              string            `json:"user_id"`
	Email               string            `json:"email"`
	Phone               string            `json:"phone,omitempty"`
	Timezone            string            `json:"timezone"`
	Language            string            `json:"language"`
	NotificationEnabled bool              `json:"notification_enabled"`
	Channels            Channels          `json:"channels"`
	Preferences         NotificationPrefs `json:"preferences"`
	UpdatedAt           time.Time         `json:"updated_at"`
}

type Channels struct {
	Email EmailChannel `json:"email"`
	Push  PushChannel  `json:"push"`
}

type EmailChannel struct {
	Enabled    bool       `json:"enabled"`
	Verified   bool       `json:"verified"`
	Frequency  string     `json:"frequency"`
	QuietHours QuietHours `json:"quiet_hours"`
}

type PushChannel struct {
	Enabled    bool       `json:"enabled"`
	Devices    []Device   `json:"devices"`
	QuietHours QuietHours `json:"quiet_hours"`
}

type Device struct {
	DeviceID string    `json:"device_id"`
	Platform string    `json:"platform"`
	Token    string    `json:"token"`
	LastSeen time.Time `json:"last_seen"`
	Active   bool      `json:"active"`
}

type QuietHours struct {
	Enabled  bool   `json:"enabled"`
	Start    string `json:"start,omitempty"`
	End      string `json:"end,omitempty"`
	Timezone string `json:"timezone,omitempty"`
}

type NotificationPrefs struct {
	Marketing     bool   `json:"marketing"`
	Transactional bool   `json:"transactional"`
	Reminders     bool   `json:"reminders"`
	Digest        Digest `json:"digest"`
}

type Digest struct {
	Enabled   bool   `json:"enabled"`
	Frequency string `json:"frequency"`
	Time      string `json:"time"`
}

type OptOutStatus struct {
	UserID    string          `json:"user_id"`
	OptedOut  bool            `json:"opted_out"`
	Channels  map[string]bool `json:"channels"`
	CheckedAt time.Time       `json:"checked_at"`
}
