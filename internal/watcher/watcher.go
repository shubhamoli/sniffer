package watcher


import (
	"internal/config"
	"internal/utils"
	"time"
)


type fn func(*time.Ticker, chan float64)
type hn func(*config.EntityToSniff, chan float64)


func _setupWatcher(_function fn, _handler hn, entity *config.EntityToSniff) {
	tick := time.NewTicker(time.Second * time.Duration(entity.Frequency))

	// Creating common channel for watcher and its notifier
	// so that they can communicate with each other
	dataCh := make(chan float64)

	// watching Entities
	go _function(tick, dataCh)
	go _handler(entity, dataCh)
}

func Setup(entity config.EntityToSniff) {

	switch entity.Entity {
		case "disk_usage": _setupWatcher(WatchDisk, utils.Trigger, &entity)
		case "mem_usage": _setupWatcher(WatchMemory, utils.Trigger, &entity)
		case "cpu_usage": _setupWatcher(WatchCPU, utils.Trigger, &entity)
	}

}
