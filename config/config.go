package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

func InitConfig(cfgPath string, cfg interface{}) error {
	err := cleanenv.ReadConfig(cfgPath, cfg)
	if err != nil {
		fmt.Printf("error load config : %s", err.Error())
		return err
	}

	return nil
}
