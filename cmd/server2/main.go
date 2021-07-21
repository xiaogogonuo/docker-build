package main

import (
	"docker-build/internal/server2/api/app"
	"docker-build/pkg/config"
	"docker-build/pkg/logger"
	"github.com/spf13/viper"
)

func setUpConfig() (v *viper.Viper, err error) {
	c := config.Config{
		ConfigName: "config",
		ConfigType: "yaml",
		ConfigPath: "configs/server2",
	}
	return c.NewConfig()
}

func main() {
	v, err := setUpConfig()
	if err != nil {
		logger.Fatal(err.Error())
	}
	app.Run(v)
}