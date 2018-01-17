package commandes

import (
	"os"
	"path/filepath"
	"fmt"
	"github.com/nargil/nargilplayer/util"
	"github.com/fatih/color"
	"github.com/nargil/nargilplayer/util/console"
)

func Add(){
	var playlist string
	playlist = "noname"
	isDirectory := false
	showHelp := false
	if len(os.Args)>2 {
		var input []string
		for i := 2; i < len(os.Args); i++{
			switch os.Args[i] {
			case "--help":
				showHelp = true
				i = len(os.Args)
			case "-r":
				isDirectory = true
			case "-p":
				i++
				playlist = os.Args[i]
			default:
				file,err := os.Open(os.Args[i])
				if !util.LogErr(err){
					absolutPath, err := filepath.Abs(file.Name())
					if !util.LogErr(err){
						input = append(input,absolutPath)
					}
				}
			}
		}
		if showHelp {
			help()
		}else{
			fmt.Println(input, isDirectory, playlist)
		}

	}else{

		fmt.Println("nargilplayer: manque 1 Argument supplementaire")
		fmt.Println("Saisissez \" nargilplayer add --help \" pour plus d'informations.")
	}
}


func ajout(){

}

func help(){
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
