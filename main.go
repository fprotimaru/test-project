package main

import (
	"github.com/fprotimaru/test-project/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Commands = []*cli.Command{
		&cmd.SendSMS, &cmd.SendEmail,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
