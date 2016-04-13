package packager

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"

	"github.com/CrowBits/gobuild/packages/configuration"
)

const StatInstallSuccess = 0
const StatInstallFail = 1

const StatPassedTest = 0
const StatFailedTest = 1
const StatNoTestFile = 2
const StatProcsError = 3
const StatTestNotRan = 4

// PackageBuild holds it all
type PackageBuild struct {
	TestList map[string]float64
}

// New Will create a new packageBuild Obj
func New(curLst []configuration.PackageStruct) (pb PackageBuild) {
	pb.TestList = make(map[string]float64)
	pb.getCurrePackageList()

	for _, pData := range curLst {
		pb.TestList[pData.Name] = pData.Percent
	}
	return
}

func (p *PackageBuild) getCurrePackageList() error {

	p.TestList[""] = 100
	files, err := ioutil.ReadDir("./packages")
	if err != nil {
		return nil
	}
	for _, file := range files {
		p.TestList[fmt.Sprint("packages/", file.Name())] = 100
	}
	return nil
}

func (p *PackageBuild) ProcessPackage(name string, runTest bool) (statInstall, statTest, statCover int, pcnt float64) {
	var err error

	// Install
	_, err = p.InstallPackage(name)
	if err != nil {
		statInstall = 1
	}

	// Run Test
	if runTest {
		statTest, pcnt = p.testPackage(name)

		if pcnt == p.TestList[name] {
			statCover = StatInstallSuccess
		} else if statTest == StatNoTestFile {
			statCover = StatTestNotRan
		} else {
			statCover = StatFailedTest
		}
	} else {
		statTest = StatTestNotRan
		statCover = StatTestNotRan
	}

	return
}

func (p *PackageBuild) InstallPackage(name string) ([]byte, error) {
	return exec.Command("go", "install", fmt.Sprintf("./%s", name)).Output()
}

func (p *PackageBuild) testPackage(name string) (stat int, cov float64) {
	out, err := exec.Command("go", "test", "-cover", fmt.Sprintf("./%s", name)).Output()

	if strings.Contains(string(out), "[no test files]") {
		return StatNoTestFile, 0
	}

	if strings.HasPrefix(string(out), "ok") {
		myOut := strings.Split(strings.Split(string(out), "coverage: ")[1], "%")[0]
		myPcnt, err2 := strconv.ParseFloat(myOut, 64)
		if err2 != nil {
			return StatProcsError, 0
		}
		return StatPassedTest, myPcnt
	}

	if err.Error() == "exit status 1" {
		return StatFailedTest, 0
	}
	return StatProcsError, 0
}
