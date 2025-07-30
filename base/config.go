package base

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type config struct {
	UserId   string  `json:"user_id"`
	DeviceId string  `json:"device_id"`
	Token    string  `json:"token"`
	Unix     string  `json:"unix"`
	KeyCode  string  `json:"key_code"`
	Base     int     `json:"base"`
	Sigma    float64 `json:"sigma"`
	Enigma   float64 `json:"enigma"`
}

var Config config

func InitConfig() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	bs, err := os.ReadFile(filepath.Join(dir, "config.json"))
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bs, &Config); err != nil {
		return err
	}

	return nil
}
