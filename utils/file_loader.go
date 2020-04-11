package utils

import (
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

const CONFIG_FILE = "./config.yml"
const USER_DATA_CSV = "./users.csv"

type Conf struct {
	Token         string `yaml:"token"`
	TargetChannel string `yaml:"target_channel"`
	SearchCount   int    `yaml:"search_count"`
}

func LoadConfig() Conf {
	log.Println("# START LoadConfig()")
	c, err := ioutil.ReadFile(CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}
	var conf Conf
	yaml.Unmarshal([]byte(c), &conf)
	validateConf(conf)
	log.Println("## Token: " + conf.Token)
	log.Println("## Target Channel: " + conf.TargetChannel)
	log.Printf("## Search Count: %d", conf.SearchCount)
	log.Println("# END LoadConfig()")
	return conf
}

func validateConf(conf Conf) {
	if conf.Token == "" {
		log.Fatal("### Token not found.")
	}
	if conf.TargetChannel == "" {
		log.Fatal("### Target Channel not found.")
	}
	if conf.SearchCount == 0 {
		log.Fatal("### Search Count must be greater than 0.")
	}
}

func ReadUserCSV() []string {
	log.Println("# START readUserData()")

	fileData, err := ioutil.ReadFile(USER_DATA_CSV)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(fileData), "\n")
	if len(lines) < 1 || lines[0] == "" {
		log.Fatal("No users in users.csv.")
	}

	log.Println("# END readUserData()")
	return lines
}
