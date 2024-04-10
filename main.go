package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/percipia/eslgo"
	"github.com/percipia/eslgo/command"
)

func main() {
	host := flag.String("H", "127.0.0.1", "pass the host of freeswitch")
	cmd := flag.String("x", "", "the command to run")
	flag.Parse()
	// Connect to FreeSWITCH
	eslgo.DefaultOptions.Logger = eslgo.NilLogger{}
	conn, err := eslgo.Dial(*host+":8021", "ClueCon", func() {
		fmt.Println("Inbound Connection Disconnected")
	})

	if err != nil {
		fmt.Println("Error connecting", err)
		return
	}

	// Create a basic context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Place the call in the background(bgapi) to user 100 and playback an audio file as the bLeg and no exported variables
	response, err := conn.SendCommand(ctx, command.API{
		Command:    *cmd,
		Background: false,
	})
	if err != nil {

		fmt.Println(err)
	} else {

		fmt.Println(string(response.Body))
	}

	conn.ExitAndClose()
}
