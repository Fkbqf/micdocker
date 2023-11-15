package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "mydocker"

	app.Commands = []cli.Command{
		initCommand,
		runCommand,
		commitCommand,
		listCommand,
		logCommand,
		execCommand,
		stopCommand,
		removeCommand,
		networkCommand,
	}

	app.Before = func(context *cli.Context) error {
		// Log as JSON instead of the default ASCII formatter.
		log.SetFormatter(&log.JSONFormatter{})

		log.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}