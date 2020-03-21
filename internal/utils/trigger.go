package utils


import (
	"fmt"
	"internal/alert"
	"internal/config"
	log "internal/logging"
	"strings"
)


func logit(level string, name string, usage float64, status string){
	name = strings.Title(strings.Replace(name, "_", " ", -1))

	switch level{
		case log.INFO:
			log.Logger.Info(fmt.Sprintf("%s: %.2f%%, Status:: %s", name, usage, status))
		case log.WARN:
			log.Logger.Warn(fmt.Sprintf("%s: %.2f%%, Status:: %s", name, usage, status))
	}
}


func _handleDefaulter(entity *config.EntityToSniff, usage float64){
	a := alert.Alert{Name: entity.Entity, Usage: usage, Notify: entity.Notify}
	a.Dispatch()
}

func Trigger(entity *config.EntityToSniff, dataCh chan float64) {
	for {
		select {
		case usage := <-dataCh:
			if int(usage) > entity.Threshold {
				logit(log.WARN, entity.Entity, usage, alert.STATUS_ALARM)
				_handleDefaulter(entity, usage)
			}else{
				logit(log.INFO, entity.Entity, usage, alert.STATUS_OK)
			}
		}
	}
}
