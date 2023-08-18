package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

// slack 文件上传机器人
func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4991956384768-5756949448114-OYOKn9vLnNMtjV3YPTZCmhpt")
	os.Setenv("CHANNEL_ID", "C04UJMNU5E0")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"1.jpg"}
	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}
