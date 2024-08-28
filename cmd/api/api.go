package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arionalmond/go-api-boiler-plate-plate/config"
	"github.com/arionalmond/go-api-boiler-plate-plate/pkg/datastore"
	"github.com/arionalmond/go-api-boiler-plate-plate/pkg/router"
)

func main() {
	conf, err := config.GetConf()
	if err != nil {
		log.Fatal(err)
	}

	mds, err := datastore.GetMySQLDS(conf)
	if err != nil {
		log.Fatal(err)
	}

	r := router.GetRouter(mds, conf.AppEnv)

	if conf.Port == 0 {
		conf.Port = 3000
	}

	log.Println("Update api now serving on port ", conf.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), r))

}
