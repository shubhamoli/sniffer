package utils


import (
	"time"
	"github.com/shirou/gopsutil/cpu"
)


func _watchCPU(dataCh chan float64) {
	data, _ := cpu.Percent(0, false)

	// push this usage percent to the channel 
	// so that utils.Notifier can deal with it
	dataCh <- data[0]
}

func WatchCPU(tick *time.Ticker, dataCh chan float64) {
	_watchCPU(dataCh)
	for {
		select {
		case <-tick.C:
			_watchCPU(dataCh)
		}
	}
}

