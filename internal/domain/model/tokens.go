package model

type Tokens struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

type ParseTokens struct {
	UserID    int64  `json:"user_id"`
	SessionID string `json:"session_id"`
	DeviceID  string `json:"device_id"`
}
