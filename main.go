package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/Doarakko/go-yugioh/yugioh"
)

func main() {
	client := yugioh.NewClient()

	card, _, err := client.RandomCards.One()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println(card.Images[0].URL)

	err = exec.Command("open", card.Images[0].URL).Run()
	if err != nil {
		log.Fatal(err)
	}
}
