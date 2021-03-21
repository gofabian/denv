package cfg

import (
	"io"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

type DenvConfig struct {
	configs []NamedConfig
}

type NamedConfig struct {
	Name  string `yaml:"name"`
	Image string `yaml:"image"`
}

func (cfg *DenvConfig) GetByName(name string) *NamedConfig {
	for _, c := range cfg.configs {
		if c.Name == name {
			return &c
		}
	}
	return nil
}

func LoadConfigFrom(path string) (*DenvConfig, error) {
	configs, err := readConfigFile(path)
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
		configs, err := readConfigFile(path)
		if err != nil {
			return nil, err
		}
		denvConfig.configs = append(denvConfig.configs, configs...)
	}

	return denvConfig, nil
}

func readConfigFile(path string) ([]NamedConfig, error) {
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

func findConfigFiles() ([]string, error) {
	paths, err := findConfigFilesInWorkingDir()
	if err != nil {
		return nil, err
	}

	p, err := findConfigFileInHomeDir()
	if err != nil {
		return nil, err
	}
	if p != "" {
		paths = append(paths, p)
	}
	return paths, nil
}

func findConfigFilesInWorkingDir() ([]string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	paths := []string{}
	prevDir := dir + "something"
	for dir != prevDir {
		path := filepath.Join(dir, ".denv.yml")
		if existsFile(path) {
			paths = append(paths, path)
		}
		prevDir = dir
		dir = filepath.Dir(dir)
	}

	return paths, nil
}

func findConfigFileInHomeDir() (string, error) {
	dir, err := homedir.Dir()
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
