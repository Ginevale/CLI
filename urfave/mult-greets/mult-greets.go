package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:  "greeting",
				Usage: "Pass multiple greetings",
			},
		},
		Action: func(cCtx *cli.Context) error {
			fmt.Println(strings.Join(cCtx.StringSlice("greeting"), `, `))
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
