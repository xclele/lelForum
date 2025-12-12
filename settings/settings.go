package settings

import (
	"fmt"
	"sync/atomic"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var conf atomic.Value // stores *AppConfig

// GetConf returns current configuration (thread-safe)
func GetConf() *AppConfig {
	return conf.Load().(*AppConfig)
}

// Conf is kept for backward compatibility (direct usage not recommended)
// Prefer using GetConf() function
var Conf = new(AppConfig)

type AppConfig struct {
	Mode              string `mapstructure:"mode"`
	Port              int    `mapstructure:"port"`
	*LogConfig        `mapstructure:"log"`
	*PostgreSQLConfig `mapstructure:"postgresql"`
	*RedisConfig      `mapstructure:"redis"`
}

type PostgreSQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"db"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	SSLMode      string `mapstructure:"sslmode"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

func Init() error {
	viper.SetConfigFile("./conf/config.yaml")

	// Load configuration for the first time
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("ReadInConfig failed, err: %v", err)
	}

	// Unmarshal to temporary variable
	tempConf := new(AppConfig)
	if err := viper.Unmarshal(tempConf); err != nil {
		return fmt.Errorf("unmarshal to Conf failed, err:%v", err)
	}

	// Atomically store configuration
	conf.Store(tempConf)
	// Also update Conf for backward compatibility
	*Conf = *tempConf

	// Start config file watcher for hot reload
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file modified, reloading...")

		newConf := new(AppConfig)
		if err := viper.Unmarshal(newConf); err != nil {
			fmt.Printf("Hot reload config failed: %v\n", err)
			return
		}

		// Atomically replace configuration
		conf.Store(newConf)
		// Also update Conf for backward compatibility
		*Conf = *newConf

		fmt.Println("Config hot reload successful!")
	})

	return nil
}
