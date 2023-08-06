package config

import (
	"encoding/json"
	"os"
)

type (
	Conf struct {
		Api     Api     `json:"api"`
		MongoDB MongoDB `json:"mongoDB"`
	}
	Api struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
	MongoDB struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		Host       string `json:"host"`
		Port       string `json:"port"`
		DBName     string `json:"dbName"`
		AuthSourse string `json:"authSourse"`
	}
)

// Загрузка конфигов
func NewConfig(path string) (*Conf, error) {
	var newConfig Conf
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	if err := json.NewDecoder(file).Decode(&newConfig); err != nil {
		return nil, err
	}
	return &newConfig, nil
}
