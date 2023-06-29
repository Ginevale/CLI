package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "cocilova",
				Aliases: []string{"coci"},
				Usage:   "coci extention",
				Subcommands: []*cli.Command{
					{
						Name:  "test",
						Usage: "gives you a question",
						Action: func(cCtx *cli.Context) error {
							var ans int
							fmt.Println("How many bits are contained in a byte?")
							fmt.Scanln(&ans)
							ans = int(ans)

							if ans == 8 {
								fmt.Println("In a byte there are 8 bits ")
							} else {
								fmt.Println("In a byte there aren't", ans, "bits. Try again!")
							}
							return nil
						},
					},
				},
			},
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("completed task: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:    "costarella",
				Aliases: []string{"costi"},
				Usage:   "costi extension",
				Subcommands: []*cli.Command{
					{
						Name:  "test",
						Usage: "gives you a question",
						Action: func(cCtx *cli.Context) error {
							var ans string
							fmt.Println("Who wrote Dorian Gray?")
							fmt.Scanln(&ans)
							ans = string(ans)

							if ans == "Oscar Wilde" {
								fmt.Println("Dorian Gray was written by Oscar Wilde")
							} else {
								fmt.Println("Dorian Gray wasn't written by ", ans, "Try again!")
							}
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "remove an existing template",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("removed task template: ", cCtx.Args().First())
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
