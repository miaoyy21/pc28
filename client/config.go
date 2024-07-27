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

	IssueURL   string `json:"issue_url"`
	HistoryURL string `json:"history_url"`
	RiddleURL  string `json:"riddle_url"`
	GoldURL    string `json:"gold_url"`
	BettingURL string `json:"betting_url"`
	SpaceURL   string `json:"space_url"`

	UserId    string `json:"user_id"`
	Token     string `json:"token"`
	PType     string `json:"p_type"`
	Unix      string `json:"unix"`
	KeyCode   string `json:"key_code"`
	DeviceId  string `json:"device_id"`
	ChannelId string `json:"channel_id"`
	Base      int    `json:"base"`
	Stop      int    `json:"stop"`
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
