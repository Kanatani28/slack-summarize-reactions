package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

const SLACK_URL = "https://slack.com/api/"
const USER_LIST_API_URL = SLACK_URL + "users.list"
const CHANNEL_LIST_API_URL = SLACK_URL + "channels.list"
const CHANNEL_HISTORY_API_URL = SLACK_URL + "channels.history"

type Reaction struct {
	Name  string   `json:"name"`
	Users []string `json:"users"`
	Count int      `json:"count"`
}

const CONFIG_FILE = "./config.yml"
const USER_DATA_CSV = "./users.csv"

type Conf struct {
	Token         string `yaml:"token"`
	TargetChannel string `yaml:"target_channel"`
	SearchCount   int    `yaml:"search_count"`
}

var conf Conf
var userCsvData []string
var users []User

func init() {
	log.Println("START init()")
	c, err := ioutil.ReadFile(CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal([]byte(c), &conf)
	log.Println("# Token: " + conf.Token)
	log.Println("# Target Channel: " + conf.TargetChannel)

	userCsvData = readUserData()
	log.Println("# User CSV Data Count:" + strconv.Itoa(len(userCsvData)))

	users = getUsers()

	log.Println("END init()")
}

func main() {

	log.Println("# START main()")

	targetChannel := getTargetChannel()

	messages := getChannelMsgs(targetChannel.ID)
	for _, message := range messages[0:conf.SearchCount] {
		fmt.Println("============================================================")
		fmt.Println("== Message: " + message.Text)
		showNoReactionUsers(message)
		fmt.Println("============================================================")
	}

	log.Println("# END main()")
}

func showNoReactionUsers(message ChannelMessage) {
	var reactionUserNames []string

	for _, reaction := range message.Reactions {
		fmt.Println("== Reaction: " + reaction.Name)
		reactionUsers := reaction.Users
		fmt.Println("== Reaction Users:")
		for _, reactionUserID := range reactionUsers {
			for _, user := range users {
				if reactionUserID == user.ID {
					fmt.Println(user.RealName)
					reactionUserNames = append(reactionUserNames, user.RealName)
				}
			}
		}
	}

	var noReactionUserNames []string

	for _, userCsvName := range userCsvData {
		userFound := false
		for _, reactionUserName := range reactionUserNames {
			if reactionUserName == userCsvName {
				userFound = true
				break
			}
		}
		if !userFound {
			noReactionUserNames = append(noReactionUserNames, userCsvName)
		}
	}

	fmt.Println("")
	fmt.Println("== No Reaction Users")
	for _, noReactionUserName := range noReactionUserNames {
		fmt.Println(noReactionUserName)
	}
}

// UsersJSON users.listで取得できるJSON
type UsersJSON struct {
	Ok      bool   `json: "ok"`
	Members []User `json: "members"`
}

// User users.listで取得できるJSONのmembersフィールドの要素
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	RealName string `json:"real_name"`
}

func getUsers() []User {
	log.Println("## START getUsers()")

	resp, err := http.Get(USER_LIST_API_URL + "?token=" + conf.Token)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	var usersJSON UsersJSON
	if err := json.Unmarshal(bytes, &usersJSON); err != nil {
		log.Fatal(err)
	}
	if !usersJSON.Ok {
		log.Fatal("### users.list FAILED")
	}

	log.Println("## END getUsers()")
	return usersJSON.Members
}

// ChannelsJSON channels.listで取得できるJSON
type ChannelsJSON struct {
	Ok       bool      `json:"ok"`
	Channels []Channel `json:"channels"`
}

// Channel channels.listで取得できるJSONのchannelsフィールドの要素
type Channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func getTargetChannel() Channel {
	channels := getChannels()
	var targetChannel Channel
	for _, channel := range channels {
		if channel.Name == conf.TargetChannel {
			targetChannel = channel
		}
	}
	return targetChannel
}

func getChannels() []Channel {
	log.Println("## START getChannels()")

	resp, err := http.Get(CHANNEL_LIST_API_URL + "?token=" + conf.Token)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	var channelsJSON ChannelsJSON
	if err := json.Unmarshal(bytes, &channelsJSON); err != nil {
		log.Fatal(err)
	}
	if !channelsJSON.Ok {
		log.Println(channelsJSON)
		log.Fatal("### channel.list FAILED")
	}

	log.Println("## END getChannels()")
	return channelsJSON.Channels
}

type ChannelMsgsJSON struct {
	Ok       bool             `json:"ok"`
	Messages []ChannelMessage `json:"messages"`
}

type ChannelMessage struct {
	Text      string     `json:"text"`
	User      string     `json:"user"`
	Reactions []Reaction `json:"reactions"`
}

func getChannelMsgs(channelID string) []ChannelMessage {
	log.Println("## START getChannelMsgs()")

	resp, err := http.Get(CHANNEL_HISTORY_API_URL + "?token=" + conf.Token + "&channel=" + channelID)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)

	var channelMsgsJSON ChannelMsgsJSON
	if err := json.Unmarshal(bytes, &channelMsgsJSON); err != nil {
		log.Fatal(err)
	}
	if !channelMsgsJSON.Ok {
		log.Fatal("### channel.history FAILED")
	}

	return channelMsgsJSON.Messages
}

func readUserData() []string {

	log.Println("## START readUserData()")

	fileData, err := ioutil.ReadFile(USER_DATA_CSV)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(fileData), "\n")

	log.Println("## END readUserData()")
	return lines
}
