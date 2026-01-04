package config

import (
	"github.com/spf13/viper"
	"time"
)

// 加载配置

type Config struct {
	App     AppConfig
	MySQL   MySQLConfig
	JWT     JWTConfig
	Log     LogConfig
	Timeout TimeoutConfig
}

type AppConfig struct {
	Name string
	Env  string
	Addr string
}

type MySQLConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	DBName          string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type JWTConfig struct {
	Secret     string
	ExpireDays int
}

type LogConfig struct {
	Level      string
	Format     string
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
}

type TimeoutConfig struct {
	Default time.Duration
	Routes  map[string]time.Duration
}

var Cfg *Config

func InitConfig() {
	v := viper.New()
	v.SetConfigName("service_config")
	v.SetConfigType("yaml")
	v.AddConfigPath("configs")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&Cfg); err != nil {
		panic(err)
	}
}
