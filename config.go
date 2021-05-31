package main

import (
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

type Recipient struct {
	Name string `yaml:"name"`
	Key  string `yaml:"key"`
}

type Config struct {
	Patches struct {
		Seed     time.Time     `yaml:"seed"`
		Interval time.Duration `yaml:"interval"`
	}
	Pushover struct {
		Server string `yaml:"server"`
		API    string `yaml:"api"`
		Token  string `yaml:"token"`
	}
	Recipients []Recipient
	Quotes     struct {
		Server string `yaml:"server"`
		API    string `yaml:"api"`
	}
}

func SaveConfig(config *Config) (bool, error) {

	data, err := yaml.Marshal(config)

	if err != nil {
		return false, err
	}

	err = ioutil.WriteFile("config.yml", data, 0644)

	if err != nil {
		return false, err
	}

	return true, nil

}

func LoadConfig() *Config {

	config := Config{}

	data, err := ioutil.ReadFile("config.yml")

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(data), &config)

	if err != nil {
		panic(err)
	}

	return &config
}
