package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/darwinfroese/hacksite/server/models"
)

// ParseConfig reads the environmentconfig file and converts it into a struct
// that can be used to setup the server
func ParseConfig(environmentFile string) models.ServerConfig {
	var config models.ServerConfig

	file, err := ioutil.ReadFile("environments/" + environmentFile)
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
