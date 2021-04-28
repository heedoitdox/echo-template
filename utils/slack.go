package utils

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func SendMessageToSlack() error {
	s := "[utils/SendMessageToSlack]"
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channelID, timestamp, err := api.PostMessage(
		"C020ZCUUDNC",
		slack.MsgOptionText("alert! you must fix it!", false),
	)

	if err != nil {
		fmt.Printf("%s %v\n", s, err)
		return err
	}

	fmt.Printf("slack message post successfully %s at %s", channelID, timestamp)
	return nil
}
