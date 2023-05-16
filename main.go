package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	canBeLaunched bool   = false
	token         string = ""
)

func main() {
	if canBeLaunched {
		Discord()
	} else {
		tok, err := GetToken()
		if err == nil {

		}
	}
}

func Discord() {
	discord, err := discordgo.New("Bot" + "MTEwNTQ3MTk1NzIxNjAxNDQ1Ng.GABr_v.Lf-sJwZgZH2I69tOPYFashOThTAGdmTGyc81ak")

	if err == nil {

	} else {
		fmt.Println("Une erreur à était levé:", err.Error())
	}

}

/*
Get the discord token for running it
*/
func GetToken() ([]byte, error) {
	file, err := os.Open("..\tk.d")
	if err != nil {
		fmt.Println("Une erreur est survenue :", err.Error())
	} else {
		t, err := os.ReadFile("..\tk.d")
		file.Close()
		if err != nil {
			return nil, errors.New("Une erreur est survenu")
		}

		return t, nil
	}

	file.Close()
	return nil, errors.New("Une erreur est survenue")
}
