package configservice

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"sync"
	"time"

	yaml "gopkg.in/yaml.v3"
)

type APIConfigFile struct {
	Endpoint []Endpoint `yaml:"routes"`
}

type Endpoint struct {
	Route        string   `yaml:"route"`
	Name         string   `yaml:"name"`
	Port         int      `yaml:"port"`
	Host         string   `yaml:"host"`
	RequestTypes []string `yaml:"request-types"`
}

// var instance *APIConfigFile
// var once sync.Once

var (
	instance          *APIConfigFile
	once              sync.Once
	fileCheckInterval *int
	configFileName    *string
)

func prepareRoutes() (APIConfigFile, error) {
	t := APIConfigFile{}
	dat, err := ioutil.ReadFile(*configFileName)
	if err != nil {
		fmt.Errorf("ERROR => ", err)
		return APIConfigFile{}, err
	}

	err = yaml.Unmarshal(dat, &t)
	if err != nil {
		fmt.Errorf("ERROR => ", err)
		return APIConfigFile{}, err
	}
	return t, nil
}

func listenFileChanges() {
	checkInterval := *fileCheckInterval
	for {
		newFile, err := prepareRoutes()
		if err != nil {
			fmt.Errorf("Invalid YAML")
			return
		}
		if !reflect.DeepEqual(newFile, *instance) {
			fmt.Println("Setting new config")
			instance = &newFile
		}
		time.Sleep(time.Duration(checkInterval) * time.Second)
	}
}

func GetInstance(confFileName *string, checkInterval *int) *APIConfigFile {
	once.Do(func() {
		fileCheckInterval = checkInterval
		configFileName = confFileName
		config, err := prepareRoutes()
		if err != nil {
			panic("Config file is invalid")
		}
		instance = &config
		go listenFileChanges()
	})
	return instance
}
