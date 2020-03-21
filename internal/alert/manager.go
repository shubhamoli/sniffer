package alert


type Alert struct{
    Name        string
    Usage       float64
    Notify      []string
}

const (
    STATUS_OK = "OK"
    STATUS_ALARM = "ALARM"

    SLACK = "slack"
    SNS = "sns"
)


func(a *Alert) Dispatch() {
    for _, notifier := range a.Notify {
        switch notifier {
            case SLACK: _triggerSlack(a)
            case SNS: _triggerEmail(a)
        }
    }
}
