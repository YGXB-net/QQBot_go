package handler

import (
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

// EventHandler 事件处理
func EventHandler(message string) {
	if gjson.Parse(message).Get("post_type").String() == "message" {

		var messageType = ""
		switch gjson.Parse(message).Get("message_type").String() {
		case "guild":

			if gjson.Parse(message).Get("userID").String() != "144115218684493716" {
				messageType = "Guild"

				var guildID = gjson.Parse(message).Get("guildID").String()
				var channelID = gjson.Parse(message).Get("channelID").String()
				var userID = gjson.Parse(message).Get("userID").String()
				var msg = gjson.Parse(message).Get("message").String()
				GuildMessage(guildID, channelID, userID, msg)
			}
		case "group":
			messageType = "Group"

			var groupID = gjson.Parse(message).Get("groupID").String()
			var userID = gjson.Parse(message).Get("userID").String()
			var msg = gjson.Parse(message).Get("message").String()

			GroupMessage(groupID, userID, msg)
		case "private":
			messageType = "Private"

			PrivateMessage(message)
		}
		log.Info("接受到["+messageType+"]消息: ", string(message))
	} else {
		return
	}
}
