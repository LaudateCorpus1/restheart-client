package restheart

import (
	"fmt"
	"io/ioutil"

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
	//_, err := os.OpenFile("../config.yaml", os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	//if !os.IsNotExist(err) {
	fileData, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		return fmt.Errorf("Unable to read config file: %s", err)
	}
	yaml.Unmarshal(fileData, &configuration)
	//} else {
	//	configuration.Endpoint = "https://odie.vianttech.com/odie"
	//}
	return nil
}
