package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

func InitConfig(cfg interface{}) error {
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		fmt.Printf("error load config : %s", err.Error())
		return err
	}

	return nil
}
