package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:     "silent",
				Aliases:  []string{"s"},
				Usage:    "no messages",
				Category: "Miscellaneous:",
			},
			&cli.BoolFlag{
				Name:     "perl-regexp",
				Aliases:  []string{"P"},
				Usage:    "PATTERNS are Perl regular expressions",
				Category: "Pattern selection:",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
