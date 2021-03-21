package cfg

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadConfigFrom(path string) (*DenvConfig, error) {
	configs, err := parseConfigFile(path)
	if err != nil {
		return nil, err
	}

	denvConfig := &DenvConfig{configs: configs}
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
		configs, err := parseConfigFile(path)
		if err != nil {
			return nil, err
		}
		denvConfig.configs = append(denvConfig.configs, configs...)
	}

	return denvConfig, nil
}

func parseConfigFile(path string) ([]NamedConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	configs := []NamedConfig{}

	for {
		cfg := NamedConfig{}
		err = decoder.Decode(&cfg)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		configs = append(configs, cfg)
	}

	return configs, nil
}
