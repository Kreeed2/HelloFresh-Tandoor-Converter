package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type Args struct {
	BaseUrl string
	Token   string
}

func main() {
	var cliArgs Args

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "Base Url",
				Aliases:     []string{"b"},
				Usage:       "Base url of the Tandoor server",
				Destination: &cliArgs.BaseUrl,
			},
			&cli.StringFlag{
				Name:        "Token",
				Aliases:     []string{"t"},
				Usage:       "Token of the Tandoor server",
				Destination: &cliArgs.Token,
			},
		},

		Name:  "HelloFresh2Tandoor",
		Usage: "Import the entire HelloFresh recipe database to your Tandoor instance",
		Action: func(ctx *cli.Context) error {
			return run(ctx, cliArgs)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(cCtx *cli.Context, args Args) error {
	new()
	return nil
}
