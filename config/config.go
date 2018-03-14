package config

// Config holds the configuration value
type Config struct {
	Version    string
	GitFolders []string
	SynFolders []string
}

// Load loads config from the path
func Load(path string) (*Config, error) {
	return nil, nil
}
