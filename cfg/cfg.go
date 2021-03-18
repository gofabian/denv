package cfg

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type DenvConfig struct {
	Image string `yaml:"image"`
}

func ReadConfigFromFile() (*DenvConfig, error) {
	cfg := &DenvConfig{}

	if !existsFile(".denv.yml") {
		return cfg, nil
	}

	yamlContent, err := ioutil.ReadFile(".denv.yml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlContent, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func existsFile(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
