package config

import (
	"encoding/json"
	"fmt"
	"github.com/caarlos0/env/v6"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	ConfigFileName = "config.yaml"
)

// LoadEnv gets configuration from environment then parse to Env variable
func LoadEnv() error {
	function := "LoadEnv"

	err := env.Parse(&Env)
	if err != nil {
		return fmt.Errorf("%s: unable to read environment configuration, %s", function, err.Error())
	}

	return nil
}

// LoadFile gets configuration from file then parse to App variable
func LoadFile() error {
	function := "LoadFile"

	// read configuration file
	data, err := os.ReadFile(ConfigFileName)
	if err != nil {
		return fmt.Errorf("%s: unable to read file configuration, %s", function, err.Error())
	}

	// unmarshal configuration data to map[string]interface{}
	var dataMap map[string]interface{}
	err = yaml.Unmarshal(data, &dataMap)
	if err != nil {
		return fmt.Errorf("%s: unable to unmarshal configuration data to map, %s", function, err.Error())
	}

	// get configuration by environment(local, dev, or etc.) and marshal to binary
	envData, err := yaml.Marshal(dataMap[Env.AppEnv])
	if err != nil {
		return fmt.Errorf("%s: unable to marshal configuration data, %s", function, err.Error())
	}

	// unmarshal configuration data to App variable
	err = yaml.Unmarshal(envData, &App)
	if err != nil {
		return fmt.Errorf("%s: unable to unmarshal configuration data to App, %s", function, err.Error())
	}

	return nil
}

func LoadOverrideENV() error {
	function := "LoadOverrideENV"
	err := env.Parse(&App)
	if err != nil {
		return fmt.Errorf("%s: unable to read environment configuration, %s", function, err.Error())
	}

	return nil
}

func LoadStandardVersion() error {
	function := "LoadStandardVersion"

	var version StandardVersion
	configFile, err := os.Open("package.json")
	defer configFile.Close()
	if err != nil {
		return fmt.Errorf("%s: unable to read file configuration, %s", function, err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&version)
	if err != nil {
		return fmt.Errorf("%s: unable to decode, %s", function, err.Error())
	}
	err = os.Setenv("APP_VERSION", version.Version)
	if err != nil {
		return fmt.Errorf("%s: unable to set ENV, %s", function, err.Error())
	}

	return nil
}

type StandardVersion struct {
	Version string `json:"version"`
}
