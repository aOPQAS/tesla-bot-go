package pgsql

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tesla/tesla-bot-go/pkg/models"
)

func (c *Client) GetUsers(ID string) ([]models.User, error) {
	s := c.GetSession()
	resp := []models.User{}
	stmt := s.Select("*").From("users")

	if ID != "" {
		stmt = stmt.Where("id = ?", ID)
	}

	_, err := stmt.Load(&resp)
	if err != nil {
		return resp, fmt.Errorf("failed to get user: %w", err)
	}

	return resp, nil
}

func (c *Client) RegisterUser(email string, password string) error {
	session := c.GetSession()
	id := uuid.New().String()
	now := float64(time.Now().Unix())

	_, err := session.Exec(
		`INSERT INTO users (id, email, password, personal_access_token, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`,
		id, email, password, "", now, now,
	)
	return err
}

func (c *Client) CreateUser() (string, error) {
	s := c.GetSession()
	createdID := uuid.New().String()
	now := float64(time.Now().Unix())

	stmt := s.InsertInto("users").Columns(
		"id", "created_at", "updated_at",
	).Values(
		createdID, now, now,
	)

	if _, err := stmt.Exec(); err != nil {
		return "", fmt.Errorf("failed to CreateUser: %w", err)
	}

	return createdID, nil
}

func (c *Client) UpdateUser(req models.User) error {
	s := c.GetSession()

	stmt := s.Update("users").SetMap(map[string]interface{}{
		"email": req.Email,
	}).Where("id = ?", req.ID)

	if _, err := stmt.Exec(); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (c *Client) DeleteUser(id string) error {
	s := c.GetSession()

	stmt := s.DeleteFrom("users").Where("id = ?", id)

	if _, err := stmt.Exec(); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

func (c *Client) GetTelegramUsers(telegramID int64) ([]models.TelegramUser, error) {
	s := c.GetSession()
	resp := []models.TelegramUser{}
	stmt := s.Select("*").From("telegram_users")

	if telegramID != 0 {
		stmt = stmt.Where("telegram_id = ?", telegramID)
	}

	_, err := stmt.Load(&resp)
	if err != nil {
		return resp, fmt.Errorf("failed to get telegram user: %w", err)
	}

	return resp, nil
}

func (c *Client) CreateTelegramUsers(user models.TelegramUser) error {
	s := c.GetSession()

	stmt := s.InsertInto("telegram_users").Columns(
		"telegram_id",
		"user_id",
	).Values(
		user.TelegramID,
		user.UserID,
	)

	if _, err := stmt.Exec(); err != nil {
		return fmt.Errorf("failed to create telegram user: %w", err)
	}

	return nil
}

func (c *Client) UpdateTelegramUsers(user models.TelegramUser) error {
	s := c.GetSession()

	stmt := s.Update("telegram_users").SetMap(map[string]interface{}{
		"user_id": user.UserID,
	}).Where("telegram_id = ?", user.TelegramID)

	if _, err := stmt.Exec(); err != nil {
		return fmt.Errorf("failed to update telegram user: %w", err)
	}

	return nil
}

func (c *Client) DeleteTelegramUsers(telegramID int64) error {
	s := c.GetSession()

	stmt := s.DeleteFrom("telegram_users").Where("telegram_id = ?", telegramID)

	if _, err := stmt.Exec(); err != nil {
		return fmt.Errorf("failed to delete telegram user: %w", err)
	}

	return nil
}

func (c *Client) GetCars(ID string) ([]models.Car, error) {
	s := c.GetSession()
	resp := []models.Car{}
	stmt := s.Select("*").From("cars")

	if ID != "" {
		stmt = stmt.Where("id = ?", ID)
	}

	_, err := stmt.Load(&resp)
	if err != nil {
		return resp, fmt.Errorf("failed to get cars: %w", err)
	}

	return resp, nil
}

func (c *Client) CreateCar(car models.Car) (string, error) {
	s := c.GetSession()
	createdID := uuid.New().String()
	car.ID = createdID

	stmt := s.InsertInto("cars").Columns(
		"id", "user_id", "latitude", "longitude",
		"battery", "is_locked", "is_charging",
		"climate_on", "last_update",
	).Values(
		car.ID, car.UserID, car.Latitude, car.Longitude,
		car.Battery, car.IsLocked, car.IsCharging,
		car.ClimateOn, car.LastUpdate,
	)

	if _, err := stmt.Exec(); err != nil {
		return "", fmt.Errorf("failed to create car: %w", err)
	}

	return createdID, nil
}

func (c *Client) UpdateCar(car models.Car) error {
	s := c.GetSession()
	car.LastUpdate = float64(time.Now().Unix())

	stmt := s.Update("cars").SetMap(map[string]interface{}{
		"user_id":     car.UserID,
		"latitude":    car.Latitude,
		"longitude":   car.Longitude,
		"battery":     car.Battery,
		"is_locked":   car.IsLocked,
		"is_charging": car.IsCharging,
		"climate_on":  car.ClimateOn,
		"last_update": car.LastUpdate,
	}).Where("id = ?", car.ID)

	if _, err := stmt.Exec(); err != nil {
		return fmt.Errorf("failed to update car: %w", err)
	}

	return nil
}

func (c *Client) DeleteCar(id string) error {
	s := c.GetSession()

	stmt := s.DeleteFrom("cars").Where("id = ?", id)

	if _, err := stmt.Exec(); err != nil {
		return fmt.Errorf("failed to delete car: %w", err)
	}

	return nil
}

func (c *Client) GetTokenPair(accessToken, refreshToken string, expiresIn int64) ([]models.TokenPair, error) {
	s := c.GetSession()
	resp := []models.TokenPair{}
	stmt := s.Select("*").From("token_pairs")

	if accessToken != "" {
		stmt = stmt.Where("access_token = ?", accessToken)
	}
	if refreshToken != "" {
		stmt = stmt.Where("refresh_token = ?", refreshToken)
	}
	if expiresIn != 0 {
		stmt = stmt.Where("expires_in = ?", expiresIn)
	}

	_, err := stmt.Load(&resp)
	if err != nil {
		return resp, fmt.Errorf("failed to get token_pair: %w", err)
	}

	return resp, nil
}

func (c *Client) CreateTokenPair(req models.TokenPair) (string, error) {
	s := c.GetSession()

	stmt := s.InsertInto("token_pairs").Columns(
		"access_token", "refresh_token", "expires_in",
	).Values(
		req.AccessToken, req.RefreshToken, req.ExpiresIn,
	)

	if _, err := stmt.Exec(); err != nil {
		return "", fmt.Errorf("failed to create token_pair: %w", err)
	}

	return req.AccessToken, nil
}

func (c *Client) UpdateTokenPair(pair models.TokenPair) error {
	s := c.GetSession()

	stmt := s.Update("token_pairs").SetMap(map[string]interface{}{
		"refresh_token": pair.RefreshToken,
		"expires_in":    pair.ExpiresIn,
	}).Where("access_token = ?", pair.AccessToken)

	if _, err := stmt.Exec(); err != nil {
		return fmt.Errorf("failed to update token pair: %w", err)
	}

	return nil
}

func (c *Client) DeleteTokenPair(accessToken string) error {
	s := c.GetSession()

	stmt := s.DeleteFrom("token_pairs").Where("access_token = ?", accessToken)

	if _, err := stmt.Exec(); err != nil {
		return fmt.Errorf("failed to delete token_pair: %w", err)
	}

	return nil
}

func (c *Client) GetLogEvent(userID string, timestamp string) ([]models.LogEvent, error) {
	s := c.GetSession()
	resp := []models.LogEvent{}
	stmt := s.Select("*").From("log_events")

	if userID != "" {
		stmt = stmt.Where("user_id = ?", userID)
	}
	if timestamp != "" {
		stmt = stmt.Where("timestamp = ?", timestamp)
	}

	_, err := stmt.Load(&resp)
	if err != nil {
		return resp, fmt.Errorf("failed to get log_event: %w", err)
	}

	return resp, nil
}

func (c *Client) CreateLogEvent(event models.LogEvent) (string, error) {
	s := c.GetSession()
	createdID := uuid.New().String()
	event.ID = createdID

	stmt := s.InsertInto("log_events").Columns(
		"id", "user_id", "event", "timestamp",
	).Values(
		event.ID, event.UserID, event.Event, event.Timestamp,
	)

	if _, err := stmt.Exec(); err != nil {
		return "", fmt.Errorf("failed to create log_event: %w", err)
	}

	return createdID, nil
}

func (c *Client) UpdateLogEvent(event models.LogEvent) error {
	s := c.GetSession()

	stmt := s.Update("log_events").SetMap(map[string]interface{}{
		"event":     event.Event,
		"timestamp": event.Timestamp,
	}).Where("id = ?", event.ID)

	if _, err := stmt.Exec(); err != nil {
		return fmt.Errorf("failed to update log event: %w", err)
	}

	return nil
}

func (c *Client) DeleteLogEvent(id string) error {
	s := c.GetSession()

	stmt := s.DeleteFrom("log_events").Where("id = ?", id)

	if _, err := stmt.Exec(); err != nil {
		return fmt.Errorf("failed to delete log_event: %w", err)
	}

	return nil
}

func (c *Client) GetUserByEmail(email string) (*models.User, error) {
	s := c.GetSession()

	var user models.User

	stmt := s.Select("*").From("users").Where("email = ?", email)

	_, err := stmt.Load(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	return &user, nil
}
