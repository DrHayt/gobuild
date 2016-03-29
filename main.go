package main

import (
	"flag"

	conf "github.com/CrowBits/gobuild/packages/configuration"
	msg "github.com/CrowBits/gobuild/packages/messaging"
)

const (
	// PathConfig is the location from the app of the config file
	PathConfig   = "./gobuild.json"
	PathAssets   = "./assets"
	PathPackages = "./packages"
)

var (
	appConfig     = conf.Data{}
	flagNoColor   = false
	pathBuild     = true
	buildRunTest  bool
	buildSkipCovr bool
	buildSkipTime bool
)

func init() {
	var err error

	// Sett values from argumrnts
	flag.BoolVar(&buildRunTest, "t", false, "Run Test for each package")
	flag.BoolVar(&buildSkipCovr, "c", false, "Dont error when coverage dose not match")
	flag.BoolVar(&buildSkipTime, "no-time", false, "Pass to update the last build time")
	flag.BoolVar(&flagNoColor, "no-color", false, "disable color in display")
	flag.Parse()

	// Load Config
	err = appConfig.Load(PathConfig)
	msg.ErrorMsg(err, 1)

	if flagNoColor {
		msg.SetNoColor()
	}
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

	if appConfig.Application.Assets {
		err := validateGoBindVersion()
		msg.ErrorMsg(err, 2)
	}

	if appConfig.Application.GoDep {
		err := validateGoDepVersion()
		msg.ErrorMsg(err, 2)
	}

	// Direct to the path to take
	switch true {
	case pathBuild:
		actionBuild(appConfig)
	}

}
