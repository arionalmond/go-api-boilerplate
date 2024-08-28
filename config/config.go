package config

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

// Config is the struct that encapsulates the config data necessary to run the application
type Config struct {
	//AppEnv is the application's environment
	AppEnv string `envconfig:"APP_ENV" required:"true"`

	//Port is the application's port
	Port int `required:"true"`
	// UseAuth bool `required:"true"`

	WorkdayEnvironment string `required:"true"`
	WorkdayTenant      string `required:"true"`
	WorkdayUserName    string `required:"true"`
	WorkdayPassword    string `required:"true"`

	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string

	// DBMaxIdleConns             int `default`
	// DBMaxConnLifetimeInMinutes int
}

// GetConf gets the Config object for the application.
// Returns default config if not 'dev' or 'prod'
// Returns an error if problem retrieving the config data
func GetConf() (conf Config, err error) {
	appEnv := strings.ToUpper(os.Getenv("APP_ENV"))
	if appEnv == "PROD" || appEnv == "PRODUCTION" || appEnv == "DEV" || appEnv == "DEVELOPMENT" {
		err = envconfig.Process("WWSC", &conf)
		if err != nil {
			return
		}
	}

	conf.AppEnv = "dev"

	file, err := os.Open("jsonConfigs/dev.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		log.Fatal(err)
	}

	// file2, err := os.Open("jsonConfigs/dev-secrets.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file2.Close()

	// err = json.NewDecoder(file2).Decode(&conf)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//Print config values in local development
	confMap := make(map[string]interface{})
	b, _ := json.Marshal(conf)
	json.Unmarshal(b, &confMap)

	// delete(confMap, "WorkdayPassword")

	log.Println("using configuration values...")
	for key, value := range confMap {
		log.Println(key, " : ", value)
	}
	log.Println("*************************")
	return
}
