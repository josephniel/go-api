package bootstrap

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/josephniel/go-api/app/config"
	yaml "gopkg.in/yaml.v2"
)

// GetConfiguration gets the appropriate configuration for a specified environment
func GetConfiguration(env *string) (conf *config.Config) {
	conf = &config.Config{}
	err := yaml.Unmarshal(readFromFile(env), &conf)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}
	return
}

func readFromFile(env *string) (b []byte) {
	envValue := *env
	absPath, err := filepath.Abs(fmt.Sprintf("config/%s.yaml", envValue))
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	b, err = ioutil.ReadFile(absPath)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return
}
