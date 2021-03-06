package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	msg "github.com/CrowBits/gobuild/packages/messaging"
)

func validateGoVersion() error {
	msg.Action(fmt.Sprintf("Is required version(%s) of go installed", appConfig.Versions.Go), ".")
	if !versionCheck(appConfig.Versions.Go, getGoVersion()) {
		msg.StatOut(*msg.ColHiRed, msg.TxtNo, ".", true)
		return fmt.Errorf("The version of go (%s) on this machine does not meet the requirements (%s)", getGoVersion(), appConfig.Versions.Go)
	}
	msg.StatOut(*msg.ColHiGreen, msg.TxtYes, ".", true)
	return nil
}

func validateGoBindVersion() error {
	msg.Action(fmt.Sprintf("Is required version(%s) of go-bindata installed", appConfig.Versions.GoBind), ".")
	if !versionCheck(appConfig.Versions.GoBind, getBindDataVersion()) {
		msg.StatOut(*msg.ColHiRed, msg.TxtNo, ".", true)
		return fmt.Errorf("The version of go-bindata (%s) on this machine does not meet the requirements (%s)", getBindDataVersion(), appConfig.Versions.GoBind)
	}
	msg.StatOut(*msg.ColHiGreen, msg.TxtYes, ".", true)
	return nil
}

func validateGoDepVersion() error {
	msg.Action(fmt.Sprintf("Is required version(%s) of godep installed", appConfig.Versions.GoDep), ".")
	if !versionCheck(appConfig.Versions.GoDep, getGoDepVersion()) {
		msg.StatOut(*msg.ColHiRed, msg.TxtNo, ".", true)
		return fmt.Errorf("The version of godep (%s) on this machine does not meet the requirements (%s)", getGoDepVersion(), appConfig.Versions.GoDep)
	}
	msg.StatOut(*msg.ColHiGreen, msg.TxtYes, ".", true)
	return nil
}

func getGoVersion() string {
	ver, _ := pullVersion([]string{"go", "version"}, "version go")
	return ver
}

func getBindDataVersion() string {
	ver, _ := pullVersion([]string{"go-bindata", "-version"}, "go-bindata ")
	return ver
}

func getGoDepVersion() string {
	ver, _ := pullVersion([]string{"godep", "version"}, " v")
	return ver
}

func versionCheck(a, b string) bool {
	var ret int
	as := strings.Split(a, ".")
	bs := strings.Split(b, ".")
	loopMax := len(bs)
	if len(as) > len(bs) {
		loopMax = len(as)
	}
	for i := 0; i < loopMax; i++ {
		var x, y string
		if len(as) > i {
			x = as[i]
		}
		if len(bs) > i {
			y = bs[i]
		}
		xi, _ := strconv.Atoi(x)
		yi, _ := strconv.Atoi(y)
		if xi > yi {
			ret = -1
		} else if xi < yi {
			ret = 1
		}
		if ret != 0 {
			break
		}
	}

	if ret < 0 {
		return false
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
