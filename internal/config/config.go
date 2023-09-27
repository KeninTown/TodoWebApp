package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Env    string `mapstructure:"env"`
	Server `mapstructure:"server"`
	Db     `mapstructure:"db"`
}

type Server struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type Db struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbname"`
	Host     string `mapstructure:"host"`
	Port     uint   `mapstructure:"port"`
}

func InitConfig(path string) (*Config, error) {
	op := "config.InitConfig()"
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("%s: failed to read config file: %w", op, err)
	}

	var cfg Config

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("%s: failed to unmarshaled config: %w", op, err)
	}

	return &cfg, nil
}
