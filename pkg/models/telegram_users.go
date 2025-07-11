package models

type TelegramUser struct {
	TelegramID int64  `json:"telegram_id" db:"telegram_id"`
	UserID     string `json:"user_id" db:"user_id"`
}
