package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Port         string `mapstructure:"port"`
	DbName       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     string `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	// 1. 指定配置文件
	viper.SetConfigFile("config.yaml")

	// 2. 指定查找配置文件的路径
	viper.AddConfigPath(".")

	// 3. 读取配置信息
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("[package: settings] [func: Init()] [viper.ReadInConfig()] failed, err: %v\n", err)
		return
	}

	//// 4. 把读取到的配置信息反序列化到全局变量Conf中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("[package: settings] [func: Init()] [viper.Unmarshal(Conf)] failed, err: %v\n", err)
	}

	// 5. 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config.yaml has been changed...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("[package: settings] [func: Init()] [viper.Unmarshal(Conf)] failed, err: %v\n", err)
		}
	})
	fmt.Printf("%#v\n", Conf)
	return
}
