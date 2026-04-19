package config

import (
	"encoding/json"
	"fmt"
	"os"
	"wvtrserv/logger"
)

type NanapiConfig struct {
	NanapiDomain   string `json:"nanapi_domain"`
	ClientUsername string `json:"client_username"`
	ClientSecret   string `json:"client_secret"`
	ClientId       string `json:"client_id"`
}

func (api NanapiConfig) String() string {
	return fmt.Sprintf("nanapi_domain: %s\nclient_username: %s\nclient_secret: %s\nclient_id: %s", api.NanapiDomain, api.ClientUsername, api.ClientSecret, api.ClientId)
}

var nanapi_config *NanapiConfig = nil

func GetNanapiConfig() *NanapiConfig {
	if nanapi_config != nil {
		return nanapi_config
	}
	nanapi_config = &NanapiConfig{}
	configFile := "nanapi_client_config.json"
	file, err := os.Open(configFile)
	if err != nil {
		logger.ErrLog.Printf("Can't open client config file (%s) : %s\n", configFile, err)
		return nil
	}

	err2 := json.NewDecoder(file).Decode(nanapi_config)
	if err2 != nil {
		logger.ErrLog.Printf("Can't open client config file (%s): %s \n", configFile, err2)
		return nil
	}

	return nanapi_config
}
