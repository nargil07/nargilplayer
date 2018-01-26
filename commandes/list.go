package commandes

import (
	"fmt"
	"github.com/nargil07/nargilplayer/metier/bdd"
)

func List(){
	mMetier := bdd.MetierBDD{}
	aMusiques := mMetier.AllMusiques()

	for i := 0; i<len(aMusiques); i++{
		aMusique := aMusiques[i]
		fmt.Println(aMusique.Name)
	}
}