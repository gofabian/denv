package cfg

type DenvConfig struct {
	files []ConfigFile
}

type ConfigFile struct {
	filename string
	configs  []NamedConfig
}

type NamedConfig struct {
	Name  string   `yaml:"name"`
	Image string   `yaml:"image"`
	Shell string   `yaml:"shell"`
	Exec  []string `yaml:"exec"`
}

func (d *DenvConfig) GetAll() []NamedConfig {
	if len(d.files) == 0 {
		return nil
	}
	// stop at first file
	return d.files[0].configs
}

func (d *DenvConfig) GetByNames(names ...string) []NamedConfig {
	var filteredConfigs []NamedConfig

	for _, file := range d.files {
		for _, namedConfig := range file.configs {
			for _, name := range names {
				if namedConfig.Name == name {
					filteredConfigs = append(filteredConfigs, namedConfig)
				}
			}
		}

		// stop at first file with name matches
		if len(filteredConfigs) > 0 {
			break
		}
	}

	return filteredConfigs
}
