package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
) // .import

var (
	ElasticHost, AppPort string
)

// init ..
func init() {

	viper.SetDefault("ELASTIC_HOST", "http://elasticsearch:9200") //"http://127.0.0.1:9200")
	viper.SetDefault("APP_PORT", "4040")

	if os.Getenv("ENVIRONMENT") == "dev" {
		_, dirname, _, _ := runtime.Caller(0)

		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		viper.AddConfigPath(filepath.Dir(dirname))
		err := viper.ReadInConfig()
		if err != nil {
			log.Printf("Error loading env vars, err: %v\n", err)
		} // .if

		ElasticHost = viper.GetString("ELASTIC_HOST_DEV")

	} else {
		viper.AutomaticEnv()
		ElasticHost = viper.GetString("ELASTIC_HOST")
	} // .else

	AppPort = viper.GetString("APP_PORT")

	log.Println("ElasticHost: ", ElasticHost)
	log.Println("AppPort: ", AppPort)

} // .init
