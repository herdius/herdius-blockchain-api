package config

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/spf13/viper"
)

type detail struct {
	SupervisorHost string
	SupervisorPort int
	TCP            string
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
		fmt.Println("Test it")
		viper.SetConfigName("config")   // Config file name without extension
		viper.AddConfigPath("./config") // Path to config file
		err := viper.ReadInConfig()
		if err != nil {
			log.Printf("Config file not found: %v", err)
		} else {
			configuration = &detail{
				SupervisorHost: viper.GetString("dev.supervisorhost"),
				SupervisorPort: viper.GetInt("dev.supervisorport"),
				TCP:            viper.GetString("dev.tcp"),
				ConnectionPort: viper.GetInt("dev.connectionport"),
			}
		}

	})

	return configuration
}

func (d *detail) GetSupervisorAddress() string {
	supervisorAddress := d.TCP + "://" + d.SupervisorHost + ":" + strconv.Itoa(d.SupervisorPort)
	return supervisorAddress
}
