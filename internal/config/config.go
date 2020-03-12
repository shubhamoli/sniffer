package config


import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)


// Config structure of yaml file
type Config struct {
	Global      Global
	Sniff       []EntityToSniff
}

type Global struct {
	Notifiers   []Notifiers
}

type Notifiers struct {
	Type        string
	Smtp		SmtpConfig `yaml:",omitempty"`
	Webhook_url string  `yaml:",omitempty"`
}

type SmtpConfig struct {
	Host		string
	Username	string
	Password	string
	From		string
}

type EntityToSniff struct {
	Entity		string
	Threshold	int
	Frequency	int
	Realert		int
	Alert		Alert
}

type Alert struct {
	Notify		[]Notify
}

type Notify struct {
	Notifier	string
	Group		string
}



func GetConfig() (*Config, error) {
	c := &Config{}

	fileName := "config.yml"
	configDir, err := filepath.Abs("src/tests")
	configFile, err := os.Open(filepath.Join(configDir, fileName))
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
			return c, errors.New("Invalid Config File") 
		}
	}

	return c, nil
}
