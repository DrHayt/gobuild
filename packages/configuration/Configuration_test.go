package configuration

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const tstJSONSuccess = "{\"application\":{\"name\":\"TestAppName\",\"lastbuilt\":\"00/00/00\"},\"versions\":{\"go\":\"1.1.1\"}}"
const tstJSONBad = "\"application\":{\"name\":\"TestAppName\",\"lastbuilt\":\"00/00/00\"},\"versions\":{\"go\":\"1.1.1\"}}"

func TestConfig(t *testing.T) {
	assert := assert.New(t)
	conf := Data{}

	// Test Dir
	testDir, err := ioutil.TempDir("", "GOBUILD_TESTING")

	// ===== Write Bad File

	// Write out config
	err = ioutil.WriteFile(fmt.Sprint(testDir, "/config.json"), []byte(tstJSONBad), 0600)
	assert.Nil(err)

	err = conf.Load(fmt.Sprint(testDir, "/config.json"))
	assert.Error(err)

	// ===== Write Good File

	// Write out config
	err = ioutil.WriteFile(fmt.Sprint(testDir, "/config.json"), []byte(tstJSONSuccess), 0600)
	assert.Nil(err)

	// Load Config
	err = conf.Load(fmt.Sprint(testDir, "/config.json"))
	assert.Nil(err)

	// Test App Info
	assert.Equal(conf.Application.Name, "TestAppName")
	assert.Equal(conf.Application.LastBuild, "00/00/00")

	// Version
	assert.Equal(conf.Versions.Go, "1.1.1")

	// ===== Make Folder unreadable

	// Change Folder Permisions
	err = os.Chmod(testDir, os.FileMode(0000))
	assert.Nil(err)

	err = conf.Load(fmt.Sprint(testDir, "/config.json"))
	assert.Error(err)

}

func saveFile() {

}
