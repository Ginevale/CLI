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

func cociTest(num int) {
	switch {
	case num == 1:
		var ans int
		fmt.Println("How many bits are contained in a byte?")
		fmt.Scanln(&ans)
		ans = int(ans)

		if ans == 8 {
			fmt.Println("In a byte there are 8 bits ")
		} else {
			fmt.Println("In a byte there aren't", ans, "bits. Try again!")
		}
	case num == 2:
		fmt.Println("L = []")
		fmt.Println("L = L.append(4)")
		fmt.Println("print (L)")
		var ans string
		fmt.Println("What does this program print?")
		fmt.Scanln(&ans)
		ans = string(ans)

		if ans == "4" {
			fmt.Println("Good job!")
		} else {
			fmt.Println(ans, "isn't correct. Try again!")
		}
	case num == 3:
		var ans string
		fmt.Println("Write a loop type")
		fmt.Scanln(&ans)
		ans = string(ans)

		if ans == "for" || ans == "while" {
			fmt.Println("That's right!")
		} else {
			fmt.Println(ans, "isn't correct. Try again!")
		}
	case num == 4:
		var ans string
		fmt.Println("What sequence can be modified? A)String B)List C)Tuple")
		fmt.Scanln(&ans)
		ans = string(ans)

		if ans == "B" || ans == "b" {
			fmt.Println("That's right!")
		} else {
			fmt.Println(ans, "isn't correct. Try again!")
		}
	}
}

func cociCuriosity(num int) {
	curiosities := [4]string{"The byte is a unit of digital information that consists of eight bits.",
		"To Insert an element at the bottom of a list: L.append(x) or to insert an element in a specific position i: L.insert(i, x)",
		"A while loop repeats as long as the condition holds true, while a for loop is used to loop through an iterable object (like a list, tuple, set, etc.) and perform the same action for each entry.",
		"Strings and Tuples are immutabile, while lists are mutable. Their values can be changed."}
	fmt.Println(curiosities[num])
}

func costiTest(num int) {
	switch {
	case num == 1:
		var ans string
		fmt.Println("Who wrote Dorian Gray?")
		fmt.Scanln(&ans)
		ans = string(ans)
		if ans == "Oscar Wilde" {
			fmt.Println("Dorian Gray was written by Oscar Wilde")
		} else {
			fmt.Println("Dorian Gray wasn't written by ", ans, "Try again!")
		}
	case num == 2:
		var ans string
		fmt.Println("What's Gulliver's role on the land of Lilliput? A)To satire on British scientific community B)To criticises humanity C)Benevolent giant")
		fmt.Scanln(&ans)
		ans = string(ans)

		if ans == "C" || ans == "c" {
			fmt.Println("Good job!")
		} else {
			fmt.Println(ans, "isn't correct. Try again!")
		}
	case num == 3:
		var ans string
		fmt.Println("When was Robinson Crusoe, by Daniel Defoe, published?")
		fmt.Scanln(&ans)
		ans = string(ans)

		if ans == "1719" {
			fmt.Println("That's right!")
		} else {
			fmt.Println(ans, "isn't correct. Try again!")
		}
	case num == 4:
		var ans string
		fmt.Println("Who wrote Beowulf?")
		fmt.Scanln(&ans)
		ans = string(ans)

		if ans == "anonymous" {
			fmt.Println("That's right, we don't know who wrote this poem!")
		} else {
			fmt.Println(ans, "isn't correct. Try again!")
		}
	}
}

func costiCuriosity(num int) {
	curiosities := [4]string{"In April 1891, Oscar Wilde's first novel The Picture of Dorian Gray was published as a book.",
		"Gulliver plays the role of a benevolent giant for little people who have exaggerated ideas about their self-importance.",
		"Robinson Crusoe is a novel by Daniel Defoe, first published on 25 April 1719.",
		"Beowulf is an anonymous epic poem, written in a West Saxon variant of Anglo-Saxon"}
	fmt.Println(curiosities[num])
}

func main() {
	questions := [4]int{1, 2, 3, 4}
	quest := questions[rand.Intn(len(questions))]
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
							cociTest(quest)
							return nil
						},
					},
					{
						Name:  "curiosity",
						Usage: "gives you explantion on topic",
						Action: func(cCtx *cli.Context) error {
							cociCuriosity(quest)
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
							costiTest(quest)
							return nil
						},
					},
					{
						Name:  "curiosity",
						Usage: "gives you explantion on topic",
						Action: func(cCtx *cli.Context) error {
							costiCuriosity(quest)
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
