package tesla

import (
	"fmt"
)

// DOORS

// UnlockDoors opens car doors
func (c *Client) UnlockDoors(vehicleID string) error {
	url := fmt.Sprintf("https://owner-api.teslamotors.com/api/1/vehicles/%s/command/door_unlock", vehicleID)
	_, err := c.request("POST", url, nil)
	return err
}

// LockDoors closes car doors
func (c *Client) LockDoors(vehicleID string) error {
	url := fmt.Sprintf("https://owner-api.teslamotors.com/api/1/vehicles/%s/command/door_lock", vehicleID)
	_, err := c.request("POST", url, nil)
	return err
}

// CLIMATE CONTROL

// StartClimate enables climate control
func (c *Client) StartClimate(vehicleID string) error {
	url := fmt.Sprintf("https://owner-api.teslamotors.com/api/1/vehicles/%s/command/auto_conditioning_start", vehicleID)
	_, err := c.request("POST", url, nil)
	return err
}

// StopClimate turns off the climate control
func (c *Client) StopClimate(vehicleID string) error {
	url := fmt.Sprintf("https://owner-api.teslamotors.com/api/1/vehicles/%s/command/auto_conditioning_stop", vehicleID)
	_, err := c.request("POST", url, nil)
	return err
}

// CHARGING

// StartCharging starts charging
func (c *Client) StartCharging(vehicleID string) error {
	url := fmt.Sprintf("https://owner-api.teslamotors.com/api/1/vehicles/%s/command/charge_start", vehicleID)
	_, err := c.request("POST", url, nil)
	return err
}

// StopCharging stops charging
func (c *Client) StopCharging(vehicleID string) error {
	url := fmt.Sprintf("https://owner-api.teslamotors.com/api/1/vehicles/%s/command/charge_stop", vehicleID)
	_, err := c.request("POST", url, nil)
	return err
}

// LIGHTS

// FlashLights flashing headlights
func (c *Client) FlashLights(vehicleID string) error {
	url := fmt.Sprintf("https://owner-api.teslamotors.com/api/1/vehicles/%s/command/flash_lights", vehicleID)
	_, err := c.request("POST", url, nil)
	return err
}

// SOUND

// HonkHorn turns on the beep
func (c *Client) HonkHorn(vehicleID string) error {
	url := fmt.Sprintf("https://owner-api.teslamotors.com/api/1/vehicles/%s/command/honk_horn", vehicleID)
	_, err := c.request("POST", url, nil)
	return err
}

// WAKE UP

// WakeUp wakes up the car from sleep
func (c *Client) WakeUp(vehicleID string) error {
	url := fmt.Sprintf("https://owner-api.teslamotors.com/api/1/vehicles/%s/wake_up", vehicleID)
	_, err := c.request("POST", url, nil)
	return err
}
