package main


import (
	"fmt"
	log "internal/logging"
	"internal/config"
	"internal/watcher"
)


func main() {
	log.Logger.Info("Starting Sniffer...")

	Config, err := config.GetConfig()
	if err != nil {
		log.Logger.Fatal(fmt.Sprintf("Error in loading Configuration. Error: %s", err.Error()))
	}

	log.Logger.Info(fmt.Sprintf("Sniffing %d items", len(Config.Sniff)))

	for _, item := range Config.Sniff {
		watcher.Setup(item)
	}

	select {}      // this is a hack to run program forever, could've have used another solid solution.

}
