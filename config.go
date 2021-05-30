package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Recipient struct {
	Key string
}

type Config struct {
	Patches struct {
		Filename   string
		Seed       string
		Length     int64
		Interval   string
		DateFormat string
	}
	Pushover struct {
		Server string
		API    string
	}
	Recipients struct {
		One struct {
			Name string
			Key  string
		}
		Two struct {
			Name string
			Key  string
		}
	}
	Quotes struct {
		Server string
		API    string
	}
}

func LoadConfig() *Config {

	config := Config{}

	data, err := ioutil.ReadFile("config.yml")

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(data), &config)

	return &config
}

/* Config class based on Jake Wright's YouTube Video */

// type Config struct {
// 	config map[string]interface{}
// }

// func (c *Config) SetFromBytes(data []byte) error {
// 	var rawConfig interface{}
// 	if err := yaml.Unmarshal(data, &rawConfig); err != nil {
// 		return err
// 	}

// 	untypedConfig, ok := rawConfig.(map[interface{}]interface{})

// 	if !ok {
// 		return fmt.Errorf("Config is not a map")
// 	}

// 	config, err := convertKeysToStrings(untypedConfig)

// 	if err != nil {
// 		return err
// 	}

// 	c.config = config

// 	return nil
// }

// func (c *Config) Get(serviceName string) (map[string]interface{}, error) {
// 	// a, ok := c.config["base"].(map[string]interface{})

// 	// if !ok {
// 	// 	return nil, fmt.Errorf("Base config is not a map")
// 	// }

// 	// if _, ok := c.config[serviceName]; !ok {
// 	// 	return a, nil
// 	// }

// 	b, ok := c.config[serviceName].(map[string]interface{})

// 	if !ok {
// 		return nil, fmt.Errorf("Service %q config is not a map", serviceName)
// 	}

// 	config := make(map[string]interface{})

// 	// for k, v := range a {
// 	// 	config[k] = v
// 	// }

// 	for k, v := range b {
// 		config[k] = v
// 	}

// 	return config, nil
// }

// func convertKeysToStrings(m map[interface{}]interface{}) (map[string]interface{}, error) {
// 	n := make(map[string]interface{})

// 	for k, v := range m {
// 		str, ok := k.(string)

// 		if !ok {
// 			return nil, fmt.Errorf("config key is not a string")
// 		}

// 		if vMap, ok := v.(map[interface{}]interface{}); ok {
// 			var err error
// 			v, err = convertKeysToStrings(vMap)

// 			if err != nil {
// 				return nil, err
// 			}
// 		}

// 		n[str] = v

// 	}

// 	return n, nil

// }
