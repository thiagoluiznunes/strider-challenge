package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort  int  `mapstructure:"SERVER_PORT"`
	ServerDebug bool `mapstructure:"SERVER_DEBUG"`

	DBName string `mapstructure:"DB_NAME"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
}

// initConfig reads in config file and ENV variables if set.
func Read() (config Config, err error) {

	var cfg Config
	err = setViperDefaults(cfg)
	if err != nil {
		return config, err
	}

	viper.SetTypeByDefaultValue(true)
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	viper.ReadInConfig()
	viper.AutomaticEnv()

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return config, err
	}

	return cfg, err
}

func setViperDefaults(cfg Config) (err error) {

	cfgMap := make(map[string]interface{})
	err = mapstructure.Decode(cfg, &cfgMap)
	if err != nil {
		return err
	}

	err = viper.MergeConfigMap(cfgMap)

	return err
}
