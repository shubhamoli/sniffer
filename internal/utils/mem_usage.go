package utils


import (
	"time"
	"github.com/shirou/gopsutil/mem"
)


func _watchMemory(dataCh chan float64) {
	data, _ := mem.VirtualMemory()

	// push this usage percent to the channel 
	// so that utils.Notifier can deal with it
	dataCh <- data.UsedPercent

}

func WatchMemory(tick *time.Ticker, dataCh chan float64) {
	_watchMemory(dataCh)
	for {
		select {
		case <-tick.C:
			_watchMemory(dataCh)
		}
	}
}
