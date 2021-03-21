package cfg

type DenvConfig struct {
	configs []NamedConfig
}

type NamedConfig struct {
	Name  string   `yaml:"name"`
	Image string   `yaml:"image"`
	Shell string   `yaml:"shell"`
	Exec  []string `yaml:"exec"`
}

func (d *DenvConfig) GetByName(name string) []NamedConfig {
	var filteredConfigs []NamedConfig
	for _, namedConfig := range d.configs {
		if namedConfig.Name == name {
			filteredConfigs = append(filteredConfigs, namedConfig)
		}
	}
	return filteredConfigs
}
