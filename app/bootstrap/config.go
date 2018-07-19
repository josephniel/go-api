package bootstrap

import (
	"fmt"

	"github.com/josephniel/go-api/app/config"
	"github.com/josephniel/go-api/app/utils"
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

func readFromFile(env *string) []byte {
	return utils.ReadFromFile(fmt.Sprintf("app/config/%s.yaml", *env))
}
