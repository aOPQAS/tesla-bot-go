package models

type Car struct {
	ID         string  `json:"id" db:"id"`
	UserID     string  `json:"user_id" db:"user_id"`
	Latitude   float64 `json:"latitude" db:"latitude"`
	Longitude  float64 `json:"longitude" db:"longitude"`
	Battery    int     `json:"battery" db:"battery"`
	IsLocked   bool    `json:"is_locked" db:"is_locked"`
	IsCharging bool    `json:"is_charging" db:"is_charging"`
	ClimateOn  bool    `json:"climate_on" db:"climate_on"`
	LastUpdate float64 `json:"last_update" db:"last_update"`
}
