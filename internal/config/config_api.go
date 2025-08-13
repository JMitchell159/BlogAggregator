package config

import (
	"os"
	"encoding/json"
)

func Read() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	fileName := home + configFileName
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	config := Config{}
	err = json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) SetUser(current_user_name string) error {
	c.Current_User_Name = &current_user_name

	bytes, err := json.Marshal(*c)
	if err != nil {
		return err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	fileName := home + configFileName
	err = os.WriteFile(fileName, bytes, 0777)
	return err
}
