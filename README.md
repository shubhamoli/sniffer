# Sniffer

Author: Shubham Oli <oli.shubham@gmail.com>
---


## Brief
Sniffer is a binary which monitors CPU, Disk and Memory usage of any node. Every resource is monitored by a GoRoutine which instrument the resource for usage at set `frequency` and once the `threshold` is crossed it notifies you via `Slack` and `Email`.


## Run it locally
```
$ git clone https://github.com/shubhamoli/sniffer
$ cd sniffer
$ go get ./...
$ export CONFIG_PATH=/path/to/config
$ go run cmd/sniffer/main.go
```


## Configuration
Configuration is done in `YML` format and following is the sample config file
```
# config.yml
---
global:
  notifiers:
    - type: slack
      webhook_url: <slack-webhook-url>


sniff:
  - entity: disk_usage       
    threshold: 20          # Percentage without % sign
    frequency: 3           # Seconds
    realert: 30            # Seconds
    notify:
      - slack              # array of notifier in global
      - email

  - entity: mem_usage
    threshold: 20
    frequency: 3
    realert: 30
    notify:
      - slack

  - entity: cpu_usage
    threshold: 80
    frequency: 3
    realert: 30
    notify:
      - slack
      - email
```

## Running as systemd service unit
A service file is already present in the repository so that Sniffer can be run a systemd service but you are free to run it via any process manager you want.


## TODOs
- [ ] Add Support for email
- [ ] Add support for reminder/re-alert
- [ ] Write tests and more tests
- [ ] Clean the code (write in more idiomatic way as per Go)
- [ ] Support Docker
- [ ] Add support for multiple slack channels and email groups.
