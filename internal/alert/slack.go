package alert


import (
	"os"
	"strings"
    "bytes"
    "encoding/json"
    "fmt"
    "internal/config"
    "net/http"
    log "internal/logging"
)


//func _triggerSlack(entityName string, usage float64) {
func _triggerSlack(a *Alert) {

	hostname, _ := os.Hostname()
    alertName := strings.Title(strings.Replace(a.Name, "_", " ", -1))

    requestBody, err := json.Marshal(map[string]string{
        "text" : fmt.Sprintf(":warning: %s high, %.2f%%, in `%v` instance", alertName, a.Usage, hostname),
    })

    if err != nil {
        log.Logger.Fatal("Error while creating requestBody")
    }

    // Naive way to search through config
    // TODO: improve this
    var webhook_url string
    for _, item := range config.Configuration.Global.Notifiers {
        if item.Type == SLACK {
            webhook_url = item.Webhook_url
        }
    }

    slackWebhookURL := webhook_url
    resp, err := http.Post(slackWebhookURL, "application/json", bytes.NewBuffer(requestBody))

    if err != nil {
        log.Logger.Fatal("Error hitting slack webhook URL, Error: %s\n", err.Error())
    }

    defer resp.Body.Close()
}
