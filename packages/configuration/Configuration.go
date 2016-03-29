package configuration

import (
	"encoding/json"
	"io/ioutil"
)

// Data Is the base of the cofiguration struct
type Data struct {
	Application applicationStruct `json:"application"`
	Versions    versionsStruct    `json:"versions"`
}

type applicationStruct struct {
	Name      string `json:"name"`
	LastBuild string `json:"lastbuilt"`
}

type versionsStruct struct {
	Go string `json:"go"`
}

// Load gobuild.json file into struct ( settings )
func (c *Data) Load(confPath string) error {
	// Setup Basics
	var err error

	// Read in Config file bytes
	configBytes, err := ioutil.ReadFile(confPath)
	if err != nil {
		return err
	}

	// Import to struct
	err = json.Unmarshal(configBytes, c)
	if err != nil {
		return err
	}

	return nil
}
