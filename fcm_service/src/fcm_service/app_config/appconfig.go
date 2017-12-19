package app_config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type AppConfiguration struct {
	// Base settings
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	RuntimeMaxProcs int    `yaml:"maxprocs"`
	Secure          bool   `yaml:"secure"`
	Protocol        string `yaml:"protocol"`
	Framed          bool   `yaml:"framed"`
	Buffered        bool   `yaml:"buffered"`
	FCM_URL         string `yaml:"fcm_url"`
	FCM_API_KEY     string `yaml:"fcm_api_key"`
	ProxyHost       string `yaml:"proxy_host"`
	ProxyPort       int    `yaml:"proxy_port"`
	Proxy           bool   `yaml:"proxy"`
}

var AppConfig *AppConfiguration

func init() {
	AppConfig = &AppConfiguration{
		Host:            "0.0.0.0",
		Port:            2701,
		RuntimeMaxProcs: 4,
		Secure:          false,
		Protocol:        "binary",
		Framed:          false,
		Buffered:        false,
		ProxyHost:       "",
		ProxyPort:       0,
		Proxy:           false,
	}
}

func InitFromYAML() bool {
	filename := os.Args[0] + ".yaml"
	fmt.Println("[FCMService] Start loading configurations from", filename)
	yamlAbsPath, err := filepath.Abs(filename)
	if err != nil {
		panic(err)
	}

	// read the raw contents of the file
	data, err := ioutil.ReadFile(yamlAbsPath)
	if err != nil {
		fmt.Println("[FCMService] Read config file error.")
		panic(err)
	}

	c := AppConfiguration{}
	// put the file's contents as yaml to the default configuration(c)
	if err := yaml.Unmarshal(data, &c); err != nil {
		panic(err)
	}

	AppConfig = &c
	fmt.Println(AppConfig)

	fmt.Println("[FCMService] Load configurations successful!")
	return true
}
