package main

import (
	"creditcard/m/v2/cardvalidity"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Enter a valid card number")
		os.Exit(0)
	}
	card := strings.Trim(os.Args[1], " ")
	expected := cardvalidity.ValidateCard(card)
	fmt.Println(fmt.Sprintf("The provided card number %v is %v", card, expected))
}
