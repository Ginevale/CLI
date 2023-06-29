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
						Usage: "gives you question",
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
					{
						Name:    "costarella",
						Aliases: []string{"costi"},
						Usage:   "costi extention",
						Subcommands: []*cli.Command{
							{
								Name:  "test",
								Usage: "gives you question",
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
