package main

import (
	"fmt"
	log "internal/logging"
	"internal/config"
	"internal/watcher"
)

func info(msg...string){
	log.Logger.Info(msg)
}

func main() {
	info("Starting Sniffer...")

	Config, err := config.GetConfig()
	if err != nil {
		log.Logger.Fatal(fmt.Sprintf("Error in loading Configuration. Error: %s", err.Error()))
	}

	info(fmt.Sprintf("Sniffing %d items", len(Config.Sniff)))
	for _, item := range Config.Sniff {
		watcher.Setup(item)
	}

	select {}      // this is a hack to run program forever, could've have used another solid solution.

}
