package models

type TokenPair struct {
	AccessToken  string `json:"access_token" db:"access_token"`
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in" db:"expires_in"`
}

type LogEvent struct {
	ID        string  `json:"id" db:"id"`
	UserID    string  `json:"user_id" db:"user_id"`
	Event     string  `json:"event" db:"event"`
	Timestamp float64 `json:"timestamp" db:"timestamp"`
}
