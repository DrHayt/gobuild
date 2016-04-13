package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	ast "github.com/CrowBits/gobuild/packages/assets"
	conf "github.com/CrowBits/gobuild/packages/configuration"
	msg "github.com/CrowBits/gobuild/packages/messaging"
	pac "github.com/CrowBits/gobuild/packages/packager"
)

func actionBuild(config conf.Data) error {

	if config.Application.Assets {
		msg.SectionHeader("Build Assets")

		// Verify We have a Assets Folder
		msg.Action(fmt.Sprintf("Is ther an  \"%s\" folder", PathAssets), ".")
		if _, err := os.Stat(PathAssets); os.IsNotExist(err) {
			msg.StatOut(*msg.ColHiRed, msg.TxtNo, ".", true)
			msg.ErrorMsg(fmt.Errorf("If you set assets to true you must have an assets folder"), 159)
		}
		msg.StatOut(*msg.ColHiGreen, msg.TxtYes, ".", true)

		// Test For any changes
		msg.Action("Have any assets been chagned since last build", ".")
		if ast.CheckForChange(config.Assets) {
			msg.StatOut(*msg.ColHiYellow, msg.TxtYes, ".", true)

			msg.Action("Build the files in the asset folder", ".")

			// Build Assets
			err := ast.BuildAssets()
			if err != nil {
				msg.StatOut(*msg.ColHiRed, msg.TxtFail, ".", true)
				msg.ErrorMsg(err, 300)
			}
			msg.StatOut(*msg.ColHiGreen, msg.TxtSuccess, ".", true)

			// Update Config
			appConfig.Assets = ast.GetFilesForConfig()

			appConfig.Save(PathConfig)
		} else {
			msg.StatOut(*msg.ColHiGreen, msg.TxtNo, ".", true)
		}
	}

	if config.Application.GoDep {
		msg.SectionHeader("GoDep .. Update Import Paths")
		msg.Action("Run GoDep on Project", ".")
		out, err := exec.Command("godep", "save", "-r", "./...").Output()
		if err != nil {
			msg.StatOut(*msg.ColHiRed, msg.TxtFail, ".", true)

			fmt.Println("-----------------------------------")
			fmt.Println(out)

			msg.ErrorMsg(err, 98)

		}
		msg.StatOut(*msg.ColHiGreen, msg.TxtSuccess, ".", true)
	}

	myPack := pac.New(appConfig.Packages)

	msg.SectionHeader("Install/Test/Cover Packages")

	msg.Action(fmt.Sprintf("Is ther a \"%s\" folder", PathPackages), ".")
	if _, err := os.Stat(PathPackages); os.IsNotExist(err) {
		msg.StatOut(*msg.ColHiYellow, msg.TxtNo, ".", true)
	} else {
		msg.StatOut(*msg.ColHiGreen, msg.TxtYes, ".", true)
	}
	var packError error

	fmt.Println("")
	msg.Col4Bar("Package", "Install", "Test", "Cover")

	for packName, _ := range myPack.TestList {

		dispName := strings.TrimSpace(packName)
		if dispName == "" {
			dispName = "main"
		}
		msg.Col4Text(dispName)

		stInstall, stTest, stCov, Prcnt := myPack.ProcessPackage(packName, buildRunTest)

		// Display Install Stat
		switch stInstall {
		case pac.StatInstallSuccess:
			msg.StatOut(*msg.ColHiGreen, msg.TxtSuccess, " ", false)
		case pac.StatInstallFail:
			packError = fmt.Errorf("Error Installing package %s", packName)
			msg.StatOut(*msg.ColHiRed, msg.TxtFail, " ", false)
		}

		// Display Text Stat
		switch stTest {
		case pac.StatPassedTest:
			msg.StatOut(*msg.ColHiGreen, msg.TxtSuccess, " ", false)
		case pac.StatFailedTest:
			packError = fmt.Errorf("Error Testing package %s", packName)
			msg.StatOut(*msg.ColHiGreen, "WHAA", " ", false)
		case pac.StatTestNotRan:
			msg.StatOut(*msg.ColHiYellow, "NotRan", " ", false)
		case pac.StatNoTestFile:
			msg.StatOut(*msg.ColHiCyan, "NoTest", " ", false)
		}

		// Display Cover Stat
		switch stCov {
		case pac.StatInstallSuccess:
			msg.StatOut(*msg.ColHiGreen, fmt.Sprint(Prcnt), " ", true)
		case pac.StatTestNotRan:
			msg.StatOut(*msg.ColHiYellow, "NotRan", " ", true)
		case pac.StatFailedTest:
			if !buildSkipCovr {
				packError = fmt.Errorf("Error with coverage on package %s", packName)
			}
			msg.StatOut(*msg.ColHiRed, fmt.Sprint(Prcnt), " ", true)
		}

	}
	msg.ErrorMsg(packError, 99)

	msg.SectionHeader("Finish Up")
	msg.Action("Install main application", ".")
	_, err := myPack.InstallPackage("")
	if err != nil {
		msg.StatOut(*msg.ColHiRed, msg.TxtFail, ".", true)
		msg.ErrorMsg(err, 2)
	}
	msg.StatOut(*msg.ColHiGreen, msg.TxtSuccess, ".", true)

	if !buildSkipTime {
		t := time.Now().UTC()
		appConfig.Application.LastBuild = fmt.Sprint(t.Format(time.RFC3339))
	}

	msg.Action("Save Config", ".")
	err = appConfig.Save(PathConfig)
	if err != nil {
		msg.StatOut(*msg.ColHiRed, msg.TxtFail, ".", true)
	} else {
		msg.StatOut(*msg.ColHiGreen, msg.TxtSuccess, ".", true)
	}

	return nil
}
