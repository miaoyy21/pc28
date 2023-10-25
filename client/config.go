package client

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Dsn string `json:"dsn"`

	Origin    string `json:"origin"`
	Cookie    string `json:"cookie"`
	UserAgent string `json:"user_agent"`

	UserId    string `json:"user_id"`
	Token     string `json:"token"`
	Unix      string `json:"unix"`
	KeyCode   string `json:"key_code"`
	DeviceId  string `json:"device_id"`
	ChannelId string `json:"channel_id"`
}

var conf Config

func InitConfig() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	bs, err := os.ReadFile(filepath.Join(dir, "config.json"))
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bs, &conf); err != nil {
		return err
	}

	return nil
}
