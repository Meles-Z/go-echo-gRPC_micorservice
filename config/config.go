package config

import "github.com/spf13/viper"

type Config struct {
	DB     DatabaseConfig
	Server ServerConfig
	Auth   AuthConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Name     string
	Username string
	Password string
}

type ServerConfig struct {
	Host string
	Port int
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() (*Config, error) {
	var cfg Config
	var err error

	viper.AddConfigPath("./")
	viper.SetConfigName(".")
	viper.SetConfigType(".env")
	viper.AutomaticEnv()

	if err = viper.BindEnv("server_host", "SERVER_HOST"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("server_port", "SERVER_PORT"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db_host", "DB_HOST"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db_port", "DB_PORT"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db_name", "DB_NAME"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db_username", "DB_USERNAME"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db_password", "DB_PASSWORD"); err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
