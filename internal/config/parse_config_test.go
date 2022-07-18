package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseConfig(t *testing.T) {
	expectedRes := Config{
		Port: 8080,
		dbConfig: &DBConfig{
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: "postgres",
			Name:     "postgres",
			Sslmode:  "disable",
		},
	}

	fileBytes, err := os.ReadFile("../../config.yaml")

	assert.Nil(t, err)

	res, err := ParseConfig(fileBytes)

	assert.Nil(t, err)

	assert.Equal(t, expectedRes.Port, res.Port)
	assert.Equal(t, expectedRes.dbConfig, res.dbConfig)
}
