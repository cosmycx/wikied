package config

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"runtime"
) // .import

var (
	ElasticHost, AppPort string
)

// init ..
func init() {

	viper.SetDefault("ELASTIC_HOST", "http://elasticsearch:9200")//"http://127.0.0.1:9200")
	viper.SetDefault("APP_PORT", "4040")

	if os.Getenv("ENVIRONMENT") == "DEV" {
		_, dirname, _, _ := runtime.Caller(0)

		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		viper.AddConfigPath(filepath.Dir(dirname))
		viper.ReadInConfig()
		//
	} else {
		viper.AutomaticEnv()
		//
	} // .else

	ElasticHost = viper.GetString("ELASTIC_HOST")
	AppPort = viper.GetString("APP_PORT")

} // .init
