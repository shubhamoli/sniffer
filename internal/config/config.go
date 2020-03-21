package config


import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	log "internal/logging"
)


// Global variable to be imported in another packages
// see: slack.go
var (
	Configuration Config
)

func init(){
	c, err := GetConfig()
	if err != nil {
		log.Logger.Fatal("No config found!")
	}

	Configuration = *c
}


// Config structure for yaml file
type Config struct {
	Global      Global
	Sniff       []EntityToSniff
}

type Global struct {
	Notifiers   []Notifiers
}

type Notifiers struct {
	Type        string
	Sns_arn		string  `yaml:",omitempty"`
	Webhook_url string  `yaml:",omitempty"`
}

type EntityToSniff struct {
	Entity		string
	Threshold	int
	Frequency	int
	Realert		int
	Notify		[]string
}


func GetConfig() (*Config, error) {
	c := &Config{}

	configFilePath := os.Getenv("CONFIG_PATH")
	if configFilePath == "" {
		configFilePath = "."
	}
	configFilePathAbs, _ := filepath.Abs(filepath.Join(configFilePath, "config.yml"))
	configFile, err := os.Open(configFilePathAbs);
	defer configFile.Close()

	if err != nil{
		return c, err
	}

	conf, err := ioutil.ReadAll(configFile)
	if err != nil {
		return c, err
	}

	if len(conf) != 0 {
		err := yaml.Unmarshal(conf, c)
		if err != nil {
			return c, errors.New("Unable to Parse Config File")
		}
	}

	return c, nil
}
