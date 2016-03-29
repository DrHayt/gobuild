package configuration

import (
	"encoding/json"
	"io/ioutil"
)

// Data Is the base of the cofiguration struct
type Data struct {
	Application applicationStruct   `json:"application"`
	Assets      []AssetsFilesStruct `json:"assets"`
	Packages    []PackageStruct     `json:"packages"`
	Versions    versionsStruct      `json:"versions"`
}

type applicationStruct struct {
	Name      string `json:"name"`
	LastBuild string `json:"lastbuilt"`
	Assets    bool   `json:"assets"`
	GoDep     bool   `json:"godep"`
}

type versionsStruct struct {
	Go     string `json:"go"`
	GoBind string `json:"go-bind"`
	GoDep  string `json:"godep"`
}

// AssetsFilesStruct is the base struct for each asset file
type AssetsFilesStruct struct {
	Path string
	Md5  string
}

// PackageStruct is the base struct for each package
type PackageStruct struct {
	Name    string  `json:"name"`
	Percent float64 `json:"coverage"`
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

// Save will save the config to file
func (c *Data) Save(confPath string) error {
	bin, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(confPath, bin, 0600)
}
