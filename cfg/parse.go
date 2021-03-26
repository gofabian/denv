package cfg

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadConfigFrom(path string) (*DenvConfig, error) {
	configFile, err := parseConfigFile(path)
	if err != nil {
		return nil, err
	}

	denvConfig := &DenvConfig{files: []ConfigFile{*configFile}}
	return denvConfig, nil
}

func LoadConfigFromDefaultDirs() (*DenvConfig, error) {
	denvConfig := &DenvConfig{}

	paths, err := findConfigFiles()
	if err != nil {
		return nil, err
	}
	if len(paths) == 0 {
		return denvConfig, nil
	}

	for _, path := range paths {
		configFile, err := parseConfigFile(path)
		if err != nil {
			return nil, err
		}
		denvConfig.files = append(denvConfig.files, *configFile)
	}

	return denvConfig, nil
}

func parseConfigFile(path string) (*ConfigFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	configFile := &ConfigFile{}

	for {
		cfg := NamedConfig{}
		err = decoder.Decode(&cfg)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		configFile.configs = append(configFile.configs, cfg)
	}

	return configFile, nil
}
