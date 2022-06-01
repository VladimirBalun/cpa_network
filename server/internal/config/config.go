package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type (
	DBConfig struct {
		Network  NetConfig `json:"network"`
		Login    string    `json:"login"`
		Password string    `json:"password"`
		DBName   string    `json:"dbname"`
	}

	ServerConfig struct {
		Network NetConfig `json:"network"`
	}

	NetConfig struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	}

	Config struct {
		Mysql  DBConfig     `json:"mysql"`
		Server ServerConfig `json:"server"`
	}
)

func ReadConfig(filename string) (cfg *Config, err error) {
	configFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	configBytes, _ := ioutil.ReadAll(configFile)

	err = json.Unmarshal(configBytes, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
