package config

import "gopkg.in/yaml.v2"

type file struct {
	Port     int `yaml:"port"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		Sslmode  string `yaml:"sslmode"`
	} `yaml:"database"`
}

func ParseConfig(fileBytes []byte) (*Config, error) {
	var f file
	err := yaml.Unmarshal(fileBytes, &f)
	if err != nil {
		return nil, err
	}

	return &Config{
		Port: f.Port,
		dbConfig: &DBConfig{
			Host:     f.Database.Host,
			Port:     f.Database.Port,
			User:     f.Database.User,
			Password: f.Database.Password,
			Name:     f.Database.Name,
			Sslmode:  f.Database.Sslmode,
		},
	}, nil
}
