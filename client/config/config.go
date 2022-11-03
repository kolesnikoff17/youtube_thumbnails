package config

import (
	"flag"
	"github.com/spf13/viper"
)

type (
	// Config -.
	Config struct {
		Grpc Client `mapstructure:"client"`
		Cli  Cli
	}
	// Client -.
	Client struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	}
	// Cli -.
	Cli struct {
		Update bool
		Async  bool
	}
)

// New is a config parser
func New() (*Config, error) {
	vp := viper.New()
	var c Config

	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("./client/config")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = vp.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	flag.BoolVar(&c.Cli.Update, "force-update", false,
		"If enabled, ignores existing cache and gets a fresh version of thumbnail")
	flag.BoolVar(&c.Cli.Async, "async", false,
		"If enabled, executes rpc and creates new file concurrently")
	flag.Parse()

	return &c, nil
}
