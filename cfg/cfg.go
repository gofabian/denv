package cfg

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

type DenvConfig struct {
	Image string `yaml:"image"`
}

func ReadConfigFromFile() (*DenvConfig, error) {
	cfg := &DenvConfig{}

	path, err := findConfigFile()
	if err != nil {
		return nil, err
	}
	if path == "" {
		return cfg, nil
	}

	yamlContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlContent, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func findConfigFile() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// search in parent directories
	prevDir := dir + "x"
	for dir != prevDir {
		path := filepath.Join(dir, ".denv.yml")
		if existsFile(path) {
			return path, nil
		}
		prevDir = dir
		dir = filepath.Dir(dir)
	}

	// search in home folder
	dir, err = homedir.Dir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(dir, ".config", "denv", ".denv.yml")
	if existsFile(path) {
		return path, nil
	}
	return "", nil
}

func existsFile(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
