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
	var config Config
	var err error
	viper.AddConfigPath("./")
	viper.SetConfigName(".")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err = viper.BindEnv("server.host", "SERVER_HOST"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("server.port", "SERVER_PORT"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.name", "DB_NAME"); err != nil {
		return nil, err
	}

	if err = viper.BindEnv("db.host", "DB_HOST"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.port", "DB_PORT"); err != nil {
		return nil, err
	}

	if err = viper.BindEnv("db.username", "DB_USERNAME"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.password", "DB_PASSWORD"); err != nil {
		return nil, err
	}
	// if err := viper.ReadInConfig(); err != nil {
	// 	return nil, fmt.Errorf("error reading config file: %w", err)
	// }
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
