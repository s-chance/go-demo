package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"os"
	"strconv"
	"time"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

// slackBot 年龄计算机器人
func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4991956384768-4953793551207-KNopxcMo3N8BgKPfUdiIFB23")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A04UG4QDF0B-4970762418372-3703bf04d5db9e19ba4cb4f1a59b569cfddc0f12412dc7905f996294bde568bc")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("你好", &slacker.CommandDefinition{
		Description: "打招呼",
		Examples:    nil,
		Handler: func(botContext slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			r := "hello"
			err := response.Reply(r)
			if err != nil {
				return
			}
		},
	})

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:    []string{"my yob is 2020"},
		Handler: func(botContext slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := time.Now().Year() - yob
			r := fmt.Sprintf("age is %d", age)
			err = response.Reply(r)
			if err != nil {
				return
			}
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
