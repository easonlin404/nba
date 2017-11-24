package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Usage = "Watch NBA live game preview, box score and player information on your console."

	app.Commands = []cli.Command{
		{
			Name:    "game",
			Aliases: []string{"i"},
			Usage:   "show today game information",
			Action: func(c *cli.Context) error {

				//TODO:

				return nil
			},
		},
	}

	app.Run(os.Args)
}
