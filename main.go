package main

import (
	"fmt"
	"log"
	"strconv"

	"slack-summarize-reactions/api"
	"slack-summarize-reactions/structs"
	"slack-summarize-reactions/utils"
)

var conf utils.Conf
var userCsvData []string
var users []structs.User

func init() {
	log.Println("# START init()")

	conf = utils.LoadConfig()
	userCsvData = utils.ReadUserCSV()
	log.Println("## User CSV Data Count: " + strconv.Itoa(len(userCsvData)))

	users = api.GetUsers(conf.Token)

	log.Println("# END init()")
}

func main() {

	log.Println("# START main()")

	targetChannel := getTargetChannel()

	messages := api.GetChannelMsgs(targetChannel.ID, conf.Token)
	for _, message := range messages[0:conf.SearchCount] {
		fmt.Println("============================================================")
		fmt.Println("== Message: " + message.Text)
		showNoReactionUsers(message)
		fmt.Println("============================================================")
	}

	log.Println("# END main()")
}

func showNoReactionUsers(message structs.ChannelMessage) {
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

func getTargetChannel() structs.Channel {
	channels := api.GetChannels(conf.Token)

	for _, channel := range channels {
		if channel.Name == conf.TargetChannel {
			return channel
		}
	}
	var noChannel structs.Channel
	log.Fatal("Channel not found: " + conf.TargetChannel)
	return noChannel
}
