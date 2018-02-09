package restheart

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Configuration creates a struct with restheart config items
type Configuration struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Endpoint string `yaml:"endpoint"`
}

// Load reads a yaml config file to source various config items
// and updates referenced struct and returns error if a read
// operation fails
func (configuration *Configuration) Load() error {
	_, err := os.Open("./config.yaml")
	if !os.IsNotExist(err) {
		fileData, err := ioutil.ReadFile("./config.yaml")
		if err != nil {
			return fmt.Errorf("Unable to read config file: %s", err)
		}
		yaml.Unmarshal(fileData, &configuration)
	} else if envExists("RHUSERNAME") && envExists("RHPASSWORD") && envExists("RHENDPOINT") {
		configuration.Username = os.Getenv("RHUSERNAME")
		configuration.Password = os.Getenv("RHPASSWORD")
		configuration.Endpoint = os.Getenv("RHENDPOINT")
	} else if envExists("RHENDPOINT") {
		configuration.Endpoint = os.Getenv("RHENDPOINT")
	}
	return nil
}

func envExists(v string) bool {
	exists := true
	env := os.Getenv(v)
	if len(env) < 1 {
		exists = false
	}
	return exists
}
