package config

import (
	"encoding/json"
	"os"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Connect struct {
	Broker   string   `json:"broker"`
	ClientID string   `json:"clientID"`
	Auth     Auth     `json:"auth"`
	Topics   []string `json:"topics"`
}

// Load reads the JSON file at path and デコードして返す
func Load(path string) (*Connect, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	cfg := &Connect{}
	if err := json.NewDecoder(f).Decode(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
