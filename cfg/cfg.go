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

func (cfg *DenvConfig) GetByName(name string) *NamedConfig {
	for _, c := range cfg.configs {
		if c.Name == name {
			return &c
		}
	}
	return nil
}
