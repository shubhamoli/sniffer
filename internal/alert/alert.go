package alert


import (
	"internal/config"
	"internal/utils"
	"time"
)


type fn func(*time.Ticker, chan float64)
type hn func(*config.EntityToSniff, chan float64)


func _setupAlert(_function fn, _handler hn, entity *config.EntityToSniff) {
	tick := time.NewTicker(time.Second * time.Duration(entity.Frequency))

	// Creating common channel for watcher and notifier
	// so that they can communicate with each other
	dataCh := make(chan float64)

	// watching Entities
	go _function(tick, dataCh)
	go _handler(entity, dataCh)
}

func SetupAlertForEntity(entity config.EntityToSniff) {

	switch entity.Entity {
		case "disk_usage": _setupAlert(utils.WatchDisk, utils.Notifier, &entity)
		case "mem_usage": _setupAlert(utils.WatchMemory, utils.Notifier, &entity)
		case "cpu_usage": _setupAlert(utils.WatchCPU, utils.Notifier, &entity)
	}

}
