package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)

func main() {
	//this is how to set environment variables in Go
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-9258691035639-9269585832147-aA81S7vJyimaSFTkA1WpKuhq")
	os.Setenv("CHANNEL_ID", "C097LLBFJ6B")
	//Get the environment variables
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"domain_checker_project_summary.pdf"}

	for i:=0; i<len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("Error: %s\n",err)
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}