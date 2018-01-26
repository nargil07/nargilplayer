package main

import (
	"fmt"
	"github.com/fatih/color"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nargil07/nargilplayer/commandes"
	"github.com/nargil07/nargilplayer/util/console"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "add":
			commandes.Add()
		case "--help":
			help()
		case "list":
			commandes.List()
		}

	} else {
		fmt.Println("nargilplayer: opérande de fichier manquant")
		fmt.Println("Saisissez \" nargilplayer --help \" pour plus d'informations.")

	}
}

/**

 */
func help() {
	fmt.Println("nargilplayer petit logiciel en ligne de commande pour écouter de la musique, gerer ses playlist.")
	fmt.Println("")
	fmt.Println(color.YellowString("usage : "))
	fmt.Println("  commande [OPTIONS]...")
	fmt.Println("")
	fmt.Println(color.YellowString("Commandes disponibles : "))
	console.NewCommandLine("add", "Ajoute un fichier dans la liste de lecture")
	fmt.Println("")
	fmt.Println("")
}