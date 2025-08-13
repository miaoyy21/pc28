package base

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type config struct {
	UserId         string  `json:"user_id"`
	DeviceId       string  `json:"device_id"`
	Token          string  `json:"token"`
	UnixNormal     string  `json:"unix_normal"`
	KeyCodeNormal  string  `json:"key_code_normal"`
	UnixMode       string  `json:"unix_mode"`
	KeyCodeMode    string  `json:"key_code_mode"`
	UnixOdds       string  `json:"unix_odds"`
	KeyCodeOdds    string  `json:"key_code_odds"`
	UnixBetting    string  `json:"unix_betting"`
	KeyCodeBetting string  `json:"key_code_betting"`
	Base           int     `json:"base"`
	Sqrt           float64 `json:"sqrt"`
	Sigma          float64 `json:"sigma"`
	Enigma         float64 `json:"enigma"`
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
