package main

import (
	"fmt"
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
	question string
	answer   string
}

func test(num int, prof string) {
	var test []asessment
	switch {
	case prof == "coci":
		test = []asessment{
			{"How many bits are contained in a byte?", "8"},
			{`L = []
L = L. append (4)
print (L)
What does this program print?`, "4"},
			{"Write a loop type", "for"},
			{"What sequence can be modified? A)String B)List C)Tuple", "b"},
		}
	case prof == "costi":
		test = []asessment{
			{"Who wrote Dorian Gray?", "Oscar Wilde"},
			{"What's Gulliver's role on the land of Lilliput? A)To satire on British scientific community B)To criticises humanity C)Benevolent giant", "c"},
			{"When was Robinson Crusoe, by Daniel Defoe, published?", "1719"},
			{"Who wrote Beowulf?", "anonymous"},
		}

	default:
		fmt.Println("error")
	}
	var ans string
	fmt.Println(test[num].question)
	fmt.Scanln(&ans)
	ans = string(ans)

	if ans == test[num].answer {
		fmt.Println("That's right!")
	} else {
		fmt.Println(ans, "isn't correct. Try again!")
	}
}

func cuoriosity(num int, prof string) {
	var curiosities []string
	switch {
	case prof == "coci":
		curiosities = []string{"The byte is a unit of digital information that consists of eight bits.",
			"To Insert an element at the bottom of a list: L.append(x) or to insert an element in a specific position i: L.insert(i, x)",
			"A while loop repeats as long as the condition holds true, while a for loop is used to loop through an iterable object (like a list, tuple, set, etc.) and perform the same action for each entry.",
			"Strings and Tuples are immutabile, while lists are mutable. Their values can be changed."}

	case prof == "costi":
		curiosities = []string{"In April 1891, Oscar Wilde's first novel The Picture of Dorian Gray was published as a book.",
			"Gulliver plays the role of a benevolent giant for little people who have exaggerated ideas about their self-importance.",
			"Robinson Crusoe is a novel by Daniel Defoe, first published on 25 April 1719.",
			"Beowulf is an anonymous epic poem, written in a West Saxon variant of Anglo-Saxon"}

	default:
		fmt.Println("error")
	}
	fmt.Println(curiosities[num])
}

func main() {
	questions := [4]int{1, 2, 3, 4}
	quest := questions[rand.Intn(len(questions))] - 1
	quest = 1
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
