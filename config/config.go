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
	SelfIP         string
}

var configuration *detail
var once sync.Once

// GetConfiguration ...
func GetConfiguration(env string) *detail {
	if env != "staging" {
		env = "dev"
	}
	log.Println("environment sourced from config:", env)
	once.Do(func() {
		viper.SetConfigName("config")   // Config file name without extension
		viper.AddConfigPath("./config") // Path to config file
		err := viper.ReadInConfig()
		if err != nil {
			log.Printf("Config file not found: %v", err)
		} else {
			configuration = &detail{
				SupervisorHost: viper.GetString(fmt.Sprint(env, ".supervisorhost")),
				SupervisorPort: viper.GetInt(fmt.Sprint(env, ".supervisorport")),
				TCP:            viper.GetString(fmt.Sprint(env, ".tcp")),
				ConnectionPort: viper.GetInt(fmt.Sprint(env, ".connectionport")),
				SelfIP:         viper.GetString(fmt.Sprint(env, ".selfip")),
			}
		}
	})

	return configuration
}

func (d *detail) GetSupervisorAddress() string {
	return d.TCP + "://" + d.SupervisorHost + ":" + strconv.Itoa(d.SupervisorPort)
}
