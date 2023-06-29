package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Action: func(cCtx *cli.Context) error {
			fmt.Printf("Hello %q", cCtx.Args().Get(0))
			//when running function include name to greet
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
