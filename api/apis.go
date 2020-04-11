package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"slack-summarize-reactions/structs"
)

const SLACK_URL = "https://slack.com/api/"
const USER_LIST_API_URL = SLACK_URL + "users.list"
const CHANNEL_LIST_API_URL = SLACK_URL + "channels.list"
const CHANNEL_HISTORY_API_URL = SLACK_URL + "channels.history"

func GetUsers(token string) []structs.User {
	log.Println("# START getUsers()")

	resp, err := http.Get(USER_LIST_API_URL + "?token=" + token)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	var usersJSON structs.UsersJSON
	if err := json.Unmarshal(bytes, &usersJSON); err != nil {
		log.Fatal(err)
	}
	if !usersJSON.Ok {
		log.Fatal("## users.list FAILED")
	}

	log.Println("# END getUsers()")
	return usersJSON.Members
}

func GetChannels(token string) []structs.Channel {
	log.Println("## START getChannels()")

	resp, err := http.Get(CHANNEL_LIST_API_URL + "?token=" + token)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	var channelsJSON structs.ChannelsJSON
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

func GetChannelMsgs(channelID string, token string) []structs.ChannelMessage {
	log.Println("## START getChannelMsgs()")

	resp, err := http.Get(CHANNEL_HISTORY_API_URL + "?token=" + token + "&channel=" + channelID)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)

	var channelMsgsJSON structs.ChannelMsgsJSON
	if err := json.Unmarshal(bytes, &channelMsgsJSON); err != nil {
		log.Fatal(err)
	}
	if !channelMsgsJSON.Ok {
		log.Fatal("### channel.history FAILED")
	}

	return channelMsgsJSON.Messages
}
