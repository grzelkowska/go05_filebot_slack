package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	// godotenv
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Could not load .env file", err)
	}
	SLACK_BOT_TOKEN := os.Getenv("SLACK_BOT_TOKEN")
	// SLACK_APP_TOKEN := os.Getenv("SLACK_APP_TOKEN")
	CHANNEL_ID := os.Getenv("CHANNEL_ID")

	// os.Setenv("SLACK_BOT_TOKEN", SLACK_BOT_TOKEN)
	// os.Setenv("SLACK_APP_TOKEN", SLACK_APP_TOKEN)
	// os.Setenv("CHANNEL_ID", CHANNEL_ID)

	// api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	api := slack.New(SLACK_BOT_TOKEN)
	channelArr := []string{CHANNEL_ID}
	fileArr := []string{"main_image_star-forming_region_carina_nircam_final-5mb.jpeg"}

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
