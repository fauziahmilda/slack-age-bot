package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
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

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4439748214674-4463554075744-aPlGy1T44ibC6ZQocrw9RJEf")                                         //this is bot token
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A04DACL7NDP-4442258537636-564e34b440c005bcbc8c1154c60c19f630a035f293501207a74e248f22762b52") //this is socket token

	//create a bot
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN")) //just make more extensible

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my year of birth is <year>", &slacker.CommandDefinition{
		Description: "year of birth calculator",
		Example:     "my year of birth is 1998",
		Handler: func(botCtx slacker.Botcontext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("error")
			}
			age := 2022 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() //make sure cancel called towards the end

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
