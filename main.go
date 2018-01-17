package main


import(
	"fmt"
	"os"
	"github.com/fatih/color"
	"github.com/nargil/nargilplayer/commandes"
	"github.com/nargil/nargilplayer/util/console"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if len(os.Args) > 1{
		switch os.Args[1] {
		case "add":
			commandes.Add()
		case "--help":
			help()
		case "test":
			testSql()
		}

	}else {
		fmt.Println("nargilplayer: opérande de fichier manquant")
		fmt.Println("Saisissez \" nargilplayer --help \" pour plus d'informations.")

	}
}


/**

 */
func help(){
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

func testSql(){
	db, err := sql.Open("sqlite3", "./base.db")
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(db)
		db.Close()

	}
}
