package logic

import (
	"errors"
	"os"
)

type Config struct {
	SonarToken string
	Proxy      string
}

var instance *Config

func GetConfig() (*Config, error) {
	if instance == nil {
		instance = &Config{}
		st, isSet := os.LookupEnv("SONAR_TOKEN")
		if isSet {
			instance.SonarToken = st
		} else {
			return nil, errors.New("sonar token env not set")
		}
		proxy, isSet := os.LookupEnv("HTTP_PROXY")
		if isSet {
			instance.Proxy = proxy
		} else {
			instance.Proxy = ""
		}

	}
	return instance, nil
}
