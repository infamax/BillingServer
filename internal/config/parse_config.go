package config

import "gopkg.in/yaml.v2"

type file struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		Sslmode  string `yaml:"sslmode"`
	} `yaml:"database"`
}

func ParseConfig(fileBytes []byte) (*config, error) {
	var cf file
	err := yaml.Unmarshal(fileBytes, &cf)
	if err != nil {
		return nil, err
	}

	return &config{
		dbConfig{
			host:     cf.Database.Host,
			port:     cf.Database.Port,
			user:     cf.Database.User,
			password: cf.Database.Password,
			name:     cf.Database.Name,
			sslmode:  cf.Database.Sslmode,
		},
	}, nil
}
