package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type detail struct {
	SupervisorHost string
	SupervisorPort int
	TCP            string
	ConnectionHost string
	ConnectionPort int
}

var configuration *detail
var once sync.Once

// GetConfiguration ...
func GetConfiguration() *detail {

	once.Do(func() {
		// Configure `export ENVIRONMENT=dev` for dev
		// and `export ENVIRONMENT=prod` for prod environment
		// in system bash file.

		viper.SetConfigName("config")   // Config file name without extension
		viper.AddConfigPath("./config") // Path to config file
		err := viper.ReadInConfig()
		if err != nil {
			log.Printf("Config file not found: %v", err)
		} else {
			var port, connectionPort int
			var host, connectionHost string
			var tcp string
			port = viper.GetInt("dev.supervisorport")
			host = viper.GetString("dev.supervisorhost")
			tcp = viper.GetString("dev.tcp")

			connectionPort = viper.GetInt("dev.connectionport")
			connectionHost = viper.GetString("dev.connectionhost")
			configuration = &detail{
				SupervisorHost: host,
				SupervisorPort: port,
				TCP:            tcp,
				ConnectionHost: connectionHost,
				ConnectionPort: connectionPort,
			}
		}

	})

	return configuration
}
