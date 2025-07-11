package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tesla/tesla-bot-go/internal/pgsql"
	"github.com/tesla/tesla-bot-go/internal/tesla"
)

type CommandHandler struct {
	PG    *pgsql.Client
	Tesla *tesla.Client
}

func (h *CommandHandler) getUserID(telegramID int64) (string, error) {
	telegramUsers, err := h.PG.GetTelegramUsers(telegramID)
	if err != nil || len(telegramUsers) == 0 {
		return "", fmt.Errorf("couldn't find a user with telegram ID %d", telegramID)
	}

	return telegramUsers[0].UserID, nil
}

func (h *CommandHandler) getUserCar(userID string) (string, error) {
	cars, err := h.PG.GetCars("")
	if err != nil {
		return "", err
	}

	for _, car := range cars {
		if car.UserID == userID {
			return car.ID, nil
		}
	}

	return "", fmt.Errorf("the user's car was not found")
}

func (h *CommandHandler) HandleLockCommand(msg *tgbotapi.Message) (string, error) {
	userID, err := h.getUserID(msg.From.ID)
	if err != nil {
		return "", fmt.Errorf("Unable to identify your user. Please try again later or register first: %w", err)
	}

	carID, err := h.getUserCar(userID)
	if err != nil {
		return "", fmt.Errorf("No vehicle found on your profile. Please add your car to proceed: %w", err)
	}

	err = h.Tesla.LockDoors(carID)
	if err != nil {
		return "", fmt.Errorf("Failed to lock the doors. Please try again later: %w", err)
	}

	return "The doors have been successfully locked.", nil
}

func (h *CommandHandler) HandleUnlockCommand(msg *tgbotapi.Message) (string, error) {
	userID, err := h.getUserID(msg.From.ID)
	if err != nil {
		return "", fmt.Errorf("Unable to identify your user. Please try again later or register first: %w", err)
	}

	carID, err := h.getUserCar(userID)
	if err != nil {
		return "", fmt.Errorf("No vehicle found on your profile. Please add your car to proceed: %w", err)
	}

	err = h.Tesla.UnlockDoors(carID)
	if err != nil {
		return "", fmt.Errorf("Failed to unlock the doors. Please try again later: %w", err)
	}

	return "The doors have been successfully unlocked.", nil
}

func (h *CommandHandler) HandleStartClimateCommand(msg *tgbotapi.Message) (string, error) {
	userID, err := h.getUserID(msg.From.ID)
	if err != nil {
		return "", fmt.Errorf("Unable to identify your user. Please try again later or register first: %w", err)
	}

	carID, err := h.getUserCar(userID)
	if err != nil {
		return "", fmt.Errorf("No vehicle found on your profile. Please add your car to proceed: %w", err)
	}

	err = h.Tesla.StartClimate(carID)
	if err != nil {
		return "", fmt.Errorf("Failed to start climate control. Please try again later: %w", err)
	}

	return "Climate control has been activated.", nil
}

func (h *CommandHandler) HandleStopClimateCommand(msg *tgbotapi.Message) (string, error) {
	userID, err := h.getUserID(msg.From.ID)
	if err != nil {
		return "", fmt.Errorf("Unable to identify your user. Please try again later or register first: %w", err)
	}

	carID, err := h.getUserCar(userID)
	if err != nil {
		return "", fmt.Errorf("No vehicle found on your profile. Please add your car to proceed: %w", err)
	}

	err = h.Tesla.StopClimate(carID)
	if err != nil {
		return "", fmt.Errorf("Failed to stop climate control. Please try again later: %w", err)
	}

	return "Climate control has been deactivated.", nil
}

func (h *CommandHandler) HandleStartChargingCommand(msg *tgbotapi.Message) (string, error) {
	userID, err := h.getUserID(msg.From.ID)
	if err != nil {
		return "", fmt.Errorf("Unable to identify your user. Please try again later or register first: %w", err)
	}

	carID, err := h.getUserCar(userID)
	if err != nil {
		return "", fmt.Errorf("No vehicle found on your profile. Please add your car to proceed: %w", err)
	}

	err = h.Tesla.StartCharging(carID)
	if err != nil {
		return "", fmt.Errorf("Failed to start charging. Please try again later: %w", err)
	}

	return "Charging has been started.", nil
}

func (h *CommandHandler) HandleStopChargingCommand(msg *tgbotapi.Message) (string, error) {
	userID, err := h.getUserID(msg.From.ID)
	if err != nil {
		return "", fmt.Errorf("Unable to identify your user. Please try again later or register first: %w", err)
	}

	carID, err := h.getUserCar(userID)
	if err != nil {
		return "", fmt.Errorf("No vehicle found on your profile. Please add your car to proceed: %w", err)
	}

	err = h.Tesla.StopCharging(carID)
	if err != nil {
		return "", fmt.Errorf("Failed to stop charging. Please try again later: %w", err)
	}

	return "Charging has been stopped.", nil
}

func (h *CommandHandler) HandleFlashLightCommand(msg *tgbotapi.Message) (string, error) {
	userID, err := h.getUserID(msg.From.ID)
	if err != nil {
		return "", fmt.Errorf("Unable to identify your user. Please try again later or register first: %w", err)
	}

	carID, err := h.getUserCar(userID)
	if err != nil {
		return "", fmt.Errorf("No vehicle found on your profile. Please add your car to proceed: %w", err)
	}

	err = h.Tesla.FlashLights(carID)
	if err != nil {
		return "", fmt.Errorf("Failed to flash the headlights. Please try again later: %w", err)
	}

	return "The headlights have flashed.", nil
}

func (h *CommandHandler) HandleHonkHorn(msg *tgbotapi.Message) (string, error) {
	userID, err := h.getUserID(msg.From.ID)
	if err != nil {
		return "", fmt.Errorf("Unable to identify your user. Please try again later or register first: %w", err)
	}

	carID, err := h.getUserCar(userID)
	if err != nil {
		return "", fmt.Errorf("No vehicle found on your profile. Please add your car to proceed: %w", err)
	}

	err = h.Tesla.HonkHorn(carID)
	if err != nil {
		return "", fmt.Errorf("Failed to honk the horn. Please try again later: %w", err)
	}

	return "The horn has been sounded.", nil
}

func (h *CommandHandler) HandleWakeUpCommand(msg *tgbotapi.Message) (string, error) {
	userID, err := h.getUserID(msg.From.ID)
	if err != nil {
		return "", fmt.Errorf("Unable to identify your user. Please try again later or register first: %w", err)
	}

	carID, err := h.getUserCar(userID)
	if err != nil {
		return "", fmt.Errorf("No vehicle found on your profile. Please add your car to proceed: %w", err)
	}

	err = h.Tesla.WakeUp(carID)
	if err != nil {
		return "", fmt.Errorf("Failed to wake up the car. Please try again later: %w", err)
	}

	return "The vehicle is now awake.", nil
}
