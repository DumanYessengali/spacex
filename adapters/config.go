package adapters

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type ConfigStruct struct {
	Port             string `yaml:"port"`
	DataBaseHost     string `yaml:"DataBaseHost"`
	DataBasePort     string `yaml:"DataBasePort"`
	DataBaseUsername string `yaml:"DataBaseUsername"`
	DataBaseDbname   string `yaml:"DataBaseDbname"`
	DataBaseSslmode  string `yaml:"DataBaseSslmode"`
	DataBasePassword string `yaml:"DataBasePassword"`
}

func (c *ConfigStruct) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

func ParseConfig() ConfigStruct {
	yfile, err := ioutil.ReadFile("configs/config.yml")

	if err != nil {

		log.Fatal(err)
	}

	var config ConfigStruct
	if err := config.Parse(yfile); err != nil {
		log.Fatal(err)
	}
	return config
}
