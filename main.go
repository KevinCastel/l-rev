package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot" + "MTEwNTQ3MTk1NzIxNjAxNDQ1Ng.GABr_v.Lf-sJwZgZH2I69tOPYFashOThTAGdmTGyc81ak")

	if err == nil {

	} else {
		fmt.Println("Une erreur à était détécté")
	}
}
