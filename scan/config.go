package main

import (
	"fmt"
	"path/filepath"
	"github.com/spf13/viper"
)

type User struct {
	Username string
	Email    string
}

var user = User{Username: "user", Email: "user@email.com"}

type Config struct {
	Name   string
	Micros []MicroInfo
}

func NewAppConfig(dir, name string, micros []*MicroInfo) *Config {

	cfg := Config{
		Name: name,
	}

	for _, micro := range micros {
		if micro != nil {
			cfg.Micros = append(cfg.Micros, (*micro))
		}
	}

	return &cfg
}

func (cfg *Config) SaveConfig(dir string) error {
	viper.SetConfigType("toml")
	path := filepath.Join(dir, "deta.toml")

	viper.Set("name", (*cfg).Name)
	for _, micro := range (*cfg).Micros {
		header := fmt.Sprintf("micros.%s", micro.Name)
		viper.Set(fmt.Sprintf("%s.directory", header), micro.Directory)
		viper.Set(fmt.Sprintf("%s.engine", header), micro.Engine)
	}

	viper.WriteConfigAs(path)
	return nil
}

func GetUser() User {
	return user
}