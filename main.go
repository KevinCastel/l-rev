package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	args := os.Args[1:]

	var username string

	if len(args) == 1 {
		if StartWith(args[0], "g-") {
			if ContainsAt(args[0], "ls:github/", 1, 10) {
				username = GetGitUsername(args[0], 11)
				fmt.Println("Recherche des repos pour l'utilisateur:", username)
				aRepos := GetAllReposFor("https://github.com/" + username + "/?tab=repositories")
				if len(aRepos) == 0 {
					fmt.Println("Erreur: rient à était trouver pour l'utilisateur:\n\tSoit ", username, " n'a pas de dépots\n\tSoit ", username, " n'existe pass")
				} else {
					fmt.Print("Voici les dépots trouver pour ", username, " :")
					for _, d := range aRepos {
						fmt.Print("\t", d)
					}
					fmt.Println()
				}
			} else if ContainsAt(args[0], "lsf:github/", 1, 10) { // équivalent de ls mais pour git précis montre tous les fichiers d'un dépot
				username = GetGitUsername(args[0], 11)
			} else if ContainsAt(args[0], "get:", 1, 4) {
				url := GetStringAt(args[0], 4)
				fmt.Println("Obtentions du fichier suivant:", url)
			} else {
				fmt.Println("Rient à était trouver")
			}
		} else {
			fmt.Println("La commande doit commencer par 'g-'")
		}
	}
}

/*
Recherche si s commence par r
*/
func StartWith(s string, r string) bool {
	if len(s) < len(r) {
		return false
	}

	var strTest string

	indexStrFound := -1

	for i := len(s) - 1; i >= 0; i-- {
		strTest = string(s[i]) + strTest

		if s[i] == r[len(r)-1] {
			strTest = string(r[len(r)-1])
			indexStrFound = i
		}
	}
	return (strTest == r && indexStrFound+1 == len(r))
}

/*
Recherche si r est dans à partir de la rune à l'index
iStart jusqu'à iEnd
*/
func ContainsAt(s, r string, iStart, iEnd int) bool {
	if len(r) > len(s) {
		return false
	}

	strGet := ""
	for i, char := range s {
		if i >= iStart+1 && i <= iEnd+1 {
			strGet += string(char)
		}
	}

	if strGet == r {
		return true
	} else {
		return false
	}
}

/*
Renvois une string contenu à partir d'un index (iStart) dans une string (s)
*/
func GetStringAt(s string, iStart int) string {
	content := ""
	for i, char := range s {
		if i >= iStart+1 {
			content += string(char)
		}
	}

	return content
}

func GetAllReposFor(url string) []string {
	sliceRepo := []string{}
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Parsing du document HTML avec GoQuery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Sélection de tous les éléments <h3> contenant les noms des dépôts
	doc.Find("h3 a").Each(func(i int, s *goquery.Selection) {
		sliceRepo = append(sliceRepo, s.Text())
	})
	return sliceRepo
}

/*
Renvois le nom d'utilisateur donner dans l'argument de la commande pour exemple:
'g-ls:github/RoiDesRats/', ici sera renvoyer 'RoiDesRats'
*/
func GetGitUsername(s string, iStart int) string {
	username := ""
	isUsernameFound := false
	for i, char := range s {
		if i >= iStart+1 {
			if char == '/' && !isUsernameFound {
				isUsernameFound = true
			} else if char != '/' {
				username += string(char)
			}
		}
	}
	return username
}

/*
Renvois lorsque une string contenus dans une string est terminé par exemple
RoisDesRats, Rois finis à l'index 4 de la string RoisDesRats
*/
func GetContentEndFromString(s, r string) int {
	index := -1
	content := ""

	for _, char := range s {
		content += string(char)
		if content == r {
			return index
		}
		index++
	}
	return index
}
