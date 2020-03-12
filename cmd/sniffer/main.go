package main

import (
	"fmt"
	log "internal/logging"
	"internal/config"
	"internal/alert"
)


func main() {
	Config, err := config.GetConfig()
	if err != nil {
		log.Logger.Fatal(fmt.Sprintf("Error in loading Configuration. Error: %s", err.Error()))
	}

	//_, err := config.ValidateConfig(Config)
	//if err != nil {
		//log.Logger.Fatal(fmt.Sprintf("Error while Validating Config. Error: %s", err.Error()))
	//}

	for _, item := range Config.Sniff {
		alert.SetupAlertForEntity(item)
	}

	select {}      // this is a hack to run program forever

}
