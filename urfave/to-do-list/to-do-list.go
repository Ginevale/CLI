package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"

	"github.com/urfave/cli/v2"
)

func assessment() {
	names := [5]string{"Mario", "Luigi", "Giada", "Luca", "Viola"}
	student := names[rand.Intn(len(names))]
	fmt.Println("Today", student, "will be assessed")
}

type asessment struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type temp struct {
	Test []asessment `json:"test"`
}

type curiosity struct {
	Curiosities string `json:"curiosities"`
}

type temp0 struct {
	Curiosity []curiosity `json:"curiosity"`
}

func test(num int, prof string) {
	var file_content temp
	var test []asessment
	switch {
	case prof == "coci":
		jsonFile, err := os.Open("./coci/test.json")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Successfully Opened test.json")
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &file_content)

		test = file_content.Test

	case prof == "costi":
		jsonFile, err := os.Open("./costi/test.json")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Successfully Opened test.json")
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &file_content)

		test = file_content.Test

	default:
		fmt.Println("Professor not found")
		return
	}

	if num >= len(test) {
		fmt.Println("Question out of range")
		return
	}
	var ans string
	fmt.Println(test[num].Question)
	fmt.Scanln(&ans)
	ans = string(ans)

	if ans == test[num].Answer {
		fmt.Println("That's right!")
	} else {
		fmt.Println(ans, "isn't correct. Try again!")
	}
}

func cuoriosity(num int, prof string) {
	var file_content0 temp0
	var curiosities []curiosity
	switch {
	case prof == "coci":
		jsonFile, err := os.Open("./coci/curiosity.json")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Successfully Opened curiosity.json")
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &file_content0)

		curiosities = file_content0.Curiosity

	case prof == "costi":
		jsonFile, err := os.Open("./costi/curiosity.json")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Successfully Opened curiosity.json")
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &file_content0)

		curiosities = file_content0.Curiosity
	default:
		fmt.Println("error")
	}
	fmt.Println(curiosities[num])
}

func main() {
	questions := [4]int{1, 2, 3, 4}
	quest := questions[rand.Intn(len(questions))] - 1
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "cocilova",
				Aliases: []string{"coci"},
				Usage:   "coci extension",
				Subcommands: []*cli.Command{
					{
						Name:  "test",
						Usage: "gives you a question",
						Action: func(cCtx *cli.Context) error {
							assessment()
							test(quest, "coci")
							return nil
						},
					},
					{
						Name:  "curiosity",
						Usage: "gives you explantion on topic",
						Action: func(cCtx *cli.Context) error {
							cuoriosity(quest, "coci")
							return nil
						},
					},
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
							assessment()
							test(quest, "costi")
							return nil
						},
					},
					{
						Name:  "curiosity",
						Usage: "gives you explantion on topic",
						Action: func(cCtx *cli.Context) error {
							cuoriosity(quest, "costi")
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
