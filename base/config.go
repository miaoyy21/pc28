package base

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type config struct {
	Token string `json:"token"`
	Base  int    `json:"base"`
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
