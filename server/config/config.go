package config

import "github.com/spf13/viper"

type (
	// Config -.
	Config struct {
		Log    Logger     `mapstructure:"logger"`
		Db     Db         `mapstructure:"db"`
		Client HttpClient `mapstructure:"http_client"`
		Grpc   Grpc       `mapstructure:"server"`
	}
	// Logger -.
	Logger struct {
		Lvl string `mapstructure:"lvl"`
	}
	// Db -.
	Db struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	}
	// HttpClient -.
	HttpClient struct {
		MaxConn int `mapstructure:"max_conn"`
		IdleTO  int `mapstructure:"idle_timeout"`
	}
	// Grpc -.
	Grpc struct {
		Port int `mapstructure:"port"`
	}
)

// New is a config parser
func New() (*Config, error) {
	vp := viper.New()
	var c Config

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("./server/config")
	vp.AddConfigPath(".")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = vp.Unmarshal(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
