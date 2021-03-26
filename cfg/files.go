package cfg

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

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
