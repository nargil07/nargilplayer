package bdd

import (
	"database/sql"
	"log"
	"os"
	"github.com/nargil07/nargilplayer/entity"
	"fmt"
	"os/user"
)

type MetierBDD struct {
}

func (m *MetierBDD) init() {
	var err error
	var db *sql.DB
	usr,_ := user.Current()
	_, err = os.Stat(usr.HomeDir + "/.nargilplayer")
	if os.IsNotExist(err) {
		err = os.Mkdir(usr.HomeDir + "/.nargilplayer", 0777)
		if err != nil{
			fmt.Printf("ligne 20 : %s", err.Error())
			log.Fatal(err)
		}
	}
	_, err = os.Stat(usr.HomeDir + "/.nargilplayer/base.db")

	if os.IsNotExist(err) { //Si la base n'existe pas.
		db, err = sql.Open("sqlite3", usr.HomeDir + "/.nargilplayer/base.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close() //Fermeture du fichier a la fin de la methode

		sqlStmt := `
					CREATE TABLE musique (
						id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
						name TEXT,
						path TEXT,
						playlist TEXT
					);
				`
		db.Exec(sqlStmt)
	}
}

func (m *MetierBDD) AllMusiques() []entity.Musique {
	m.init()

	usr,_ := user.Current()
	db, err := sql.Open("sqlite3", usr.HomeDir + "/.nargilplayer/base.db")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT name FROM musique")
	if err != nil{
		log.Fatal(err)
	}
	var aMusiques []entity.Musique
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)
		mus := entity.Musique{}
		mus.SetName(name)
		aMusiques = append(aMusiques, mus)
	}

	rows.Close()

	return aMusiques
}

func (m *MetierBDD) AddMusique(musique entity.Musique) {
	m.init()

	usr,_ := user.Current()
	db, err := sql.Open("sqlite3", usr.HomeDir + "/.nargilplayer/base.db")
	if err != nil {
		log.Fatal(err)
	}
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("INSERT INTO musique (name, path) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(musique.Name, musique.Path)
	tx.Commit()
}
