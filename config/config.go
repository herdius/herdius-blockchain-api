package config

import (
	"fmt"
	"log"
	"os"
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
	DBHost         string
	DBPort         int
	DBName         string
	DBUser         string
	DBPass         string
	DBSslMode      bool
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
		dirname := os.Getenv("GOPATH")
		viper.SetConfigName("config")                                                          // Config file name without extension
		viper.AddConfigPath(dirname + "/src/github.com/herdius/herdius-blockchain-api/config") // Path to config file
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
				DBHost:         viper.GetString(fmt.Sprint(env, ".dbhost")),
				DBPort:         viper.GetInt(fmt.Sprint(env, ".dbport")),
				DBName:         viper.GetString(fmt.Sprint(env, ".dbname")),
				DBUser:         viper.GetString(fmt.Sprint(env, ".dbuser")),
				DBPass:         viper.GetString(fmt.Sprint(env, ".dbpass")),
				DBSslMode:      viper.GetBool(fmt.Sprint(env, ".dbsslmode")),
			}
		}
	})

	return configuration
}

func (d *detail) GetSupervisorAddress() string {
	return d.TCP + "://" + d.SupervisorHost + ":" + strconv.Itoa(d.SupervisorPort)
}

func (d *detail) DBConnString() string {
	mode := "disable"
	if d.DBSslMode {
		mode = "enable"
	}
	connStrFmt := "user=%s host=%s port=%d dbname=%s password=%s sslmode=%s"
	return fmt.Sprintf(connStrFmt, d.DBUser, d.DBHost, d.DBPort, d.DBName, d.DBPass, mode)
}
