package main

import (
	"fmt"
	"os/exec"
	"strings"

	msg "github.com/CrowBits/gobuild/packages/messaging"
)

func validateGoVersion() error {
	msg.Action(fmt.Sprintf("Is required version(%s) of go installed", appConfig.Versions.Go), ".")
	if !versionCheck(appConfig.Versions.Go, getGoVersion()) {
		msg.StatOut(*msg.ColHiRed, msg.TxtNo, ".", true)
		return fmt.Errorf("The version of go (%s) on this machine dose not meat the requirements (%s)", getGoVersion(), appConfig.Versions.Go)
	}
	msg.StatOut(*msg.ColHiGreen, msg.TxtYes, ".", true)
	return nil
}

func getGoVersion() string {
	ver, _ := pullVersion([]string{"go", "version"}, "version go")
	return ver
}

func versionCheck(baseVer, testVer string) bool {
	// Split version strins
	baseDat := strings.Split(baseVer, ".")
	testDat := strings.Split(testVer, ".")

	// Length must match
	if len(baseDat) != len(testDat) {
		return false
	}

	// Check if newer
	for idx, val := range baseDat {
		if val > testDat[idx] {
			return false
		}
	}
	return true
}

func pullVersion(cmd []string, parsStr string) (string, error) {
	out, err := exec.Command(cmd[0], cmd[1:]...).Output()
	if err != nil {
		return "", err
	}

	tmp := strings.Split(string(out), parsStr)
	if len(tmp) > 1 {
		tmp = strings.Split(tmp[1], " ")
		return tmp[0], nil
	}
	return "", fmt.Errorf("The Output was not what was expected")
}
