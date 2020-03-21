package watcher


import (
	"time"
	"github.com/shirou/gopsutil/disk"
)


func _watchDisk(dataCh chan float64) {
	data, _ := disk.Usage("/")

	// push this usage percent to the channel 
	// so that utils.Notifier can deal with it
	dataCh <- data.UsedPercent
}

func WatchDisk(tick *time.Ticker, dataCh chan float64) {
	_watchDisk(dataCh)
	for {
		select {
		case <-tick.C:
			_watchDisk(dataCh)
		}
	}
}

