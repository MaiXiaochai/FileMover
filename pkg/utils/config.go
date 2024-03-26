package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Path PathConfig `mapstructure:"path"`
}

type PathConfig struct {
	SrcDir  string `mapstructure:"src_dir"`
	DestDir string `mapstructure:"dest_dir"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error reading config file, %w", err))
	}

	var config Config
	// 将读取的配置信息映射到结构体
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %w", err))
	}

	return &config, nil

}
