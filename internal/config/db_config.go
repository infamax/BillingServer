package config

import "fmt"

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	Sslmode  string
}

func (dconfig *DBConfig) GetConnectionString() string {
	return fmt.Sprintf("host = %s port = %d user = %s password = %s name = %s sslmode = %s",
		dconfig.Host, dconfig.Port, dconfig.User, dconfig.Password, dconfig.Name, dconfig.Sslmode)
}
