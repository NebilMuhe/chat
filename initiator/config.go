package initiator

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.AddConfigPath("config")
	viper.SetConfigName("test_config")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("failed to initialize config")
	}

	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("config changed", in.Name)
	})
}
