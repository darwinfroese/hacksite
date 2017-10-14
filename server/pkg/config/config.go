package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// EnvironmentConfig contains the configuration for the enivornment
type EnvironmentConfig struct {
	Server   serverConfig
	Logger   logConfig
	Database dbConfig
	Aws      awsConfig
}

type serverConfig struct {
	Port, KeyLocation, CertLocation, WebFileLocation string
}
type logConfig struct {
	LogFileLocation string
}
type dbConfig struct {
	System string
}
type awsConfig struct {
	AccessKey, SecretKey, Token, Region string
}

// ParseConfig reads the environmentconfig file and converts it into a struct
// that can be used to setup the server
func ParseConfig(environmentFile string) EnvironmentConfig {
	var config EnvironmentConfig

	fmt.Printf("Using config file: %s\n", environmentFile)

	if _, err := os.Stat(environmentFile); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Couldn't find the file: %s\n", environmentFile)
		return config
	}

	file, err := ioutil.ReadFile(environmentFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return config
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return config
	}

	return config
}
