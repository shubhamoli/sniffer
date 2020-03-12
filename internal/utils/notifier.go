package utils


import (
	"fmt"
	"internal/config"
)

func _handleDefaulter(entity *config.EntityToSniff, usage float64){
	fmt.Printf("%s usage: %.2f\n", entity.Entity, usage)
}

func Notifier(entity *config.EntityToSniff, dataCh chan float64) {
	for {
		select {
		case usage := <-dataCh:
			if int(usage) > entity.Threshold {
				_handleDefaulter(entity, usage)
			}
		}
	}
}
