package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		DbName   string `mapstructure:"dbName"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("failed to read configuration %w", err))
	}

	fmt.Println("Server port: ", viper.GetInt("server.port"))
	fmt.Println("JWT key: ", viper.GetString("security.jwt.key"))

	// config := Config{}
	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}

	fmt.Println("Server port: ", config.Server.Port)

	for _, v := range config.Databases {
		fmt.Println(v.User, v.Password, v.Host, v.DbName)
	}
}
