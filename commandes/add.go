package commandes

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/nargil07/nargilplayer/util"
	"github.com/nargil07/nargilplayer/util/console"
	"os"
	"path/filepath"
	"github.com/nargil07/nargilplayer/entity"
	"github.com/nargil07/nargilplayer/metier/bdd"
)

func Add() {
	var _ string
	_ = "noname"
	_ = false
	showHelp := false
	if len(os.Args) > 2 {
		var input []string
		for i := 2; i < len(os.Args); i++ {
			switch os.Args[i] {
			case "--help":
				showHelp = true
				i = len(os.Args)
			case "-r":
				_ = true
			case "-p":
				i++
				_ = os.Args[i]
			default:
				file, err := os.Open(os.Args[i])
				if !util.LogErr(err) {
					absolutPath, err := filepath.Abs(file.Name())
					if !util.LogErr(err) {
						input = append(input, absolutPath)
					}
				}
			}
		}
		if showHelp {
			help()
		} else {
			mMetMusique := bdd.MetierBDD{}
			for j := 0; j < len(input) ; j++{
				musique := entity.Musique{}
				musique.SetName(input[j])

				fmt.Println(musique)
				mMetMusique.AddMusique(musique)
			}

		}

	} else {

		fmt.Println("nargilplayer: manque 1 Argument supplementaire")
		fmt.Println("Saisissez \" nargilplayer add --help \" pour plus d'informations.")
	}
}

func ajout() {

}

func help() {
	fmt.Println(color.YellowString("usage : "), "add [OPTIONS]... FILE...")
	fmt.Println("")
	fmt.Println("Ajoute les fichiers données dans la liste de lecture.")
	fmt.Println("de lecture")
	fmt.Println("")
	fmt.Println(color.YellowString("Options disponibles : "))
	console.NewOptionsLine("-r", "", "Ajoute recursivement les répertoires")
	console.NewOptionsLine("", "--help", "Affiche cette aide")
	console.NewOptionsLine("-p", "--playlist", "Ajoute les musiques dans une playlist nommées")
	fmt.Println("")
	fmt.Println("")
}
