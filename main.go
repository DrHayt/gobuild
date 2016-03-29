package main

import (
	"flag"

	conf "github.com/CrowBits/gobuild/packages/configuration"
	msg "github.com/CrowBits/gobuild/packages/messaging"
)

const (
	// PathConfig is the location from the app of the config file
	PathConfig = "./gobuild.json"
)

var (
	appConfig   = conf.Data{}
	flagNoColor = false
	pathBuild   = true
)

func init() {
	var err error

	// Sett values from argumrnts
	flag.BoolVar(&flagNoColor, "no-color", false, "disable color in display")
	flag.Parse()

	// Load Config
	err = appConfig.Load(PathConfig)
	msg.ErrorMsg(err, 1)
}

func main() {

	// App Info
	msg.SectionHeader("Application Information")
	msg.KeyVal("Name:", appConfig.Application.Name)
	msg.KeyVal("Last Built:", appConfig.Application.LastBuild)

	// Verify System
	msg.SectionHeader("Verify System")
	err := validateGoVersion()
	msg.ErrorMsg(err, 2)

	// Direct to the path to take
	switch true {
	case pathBuild:

	}

}
