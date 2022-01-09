package common

import (
	"encoding/json"
	"os"
)

type Config struct {
	JwtSigningSecret string `json:"jwt_signing_secret"`
}

func (c *Config) Load() error {
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(c)
	if err != nil {
		return err
	}
	return nil
}
