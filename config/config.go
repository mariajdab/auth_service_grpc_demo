package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type configuration struct {
	AuthToken   string `json:"auth_token"`
	AccountSSID string `json:"account_ssid"`
	ServicesID  string `json:"services_id"`
}

var Configuration configuration

func LoadConfig(filepath string) error {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("could not read file: %w", err)
	}

	err = json.Unmarshal(data, &Configuration)
	if err != nil {
		return fmt.Errorf("could not read file: %w", err)
	}

	return nil
}
