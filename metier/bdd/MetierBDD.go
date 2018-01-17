package bdd

import (
	"database/sql"
	"log"
)

type MetierBDD struct {
	db* sql.DB
}

func (m* MetierBDD) init()  {
	var err error;
	m.db, err = sql.Open("sqlite3", "./base.db")
	if err != nil{
		log.Fatal(err)
	}
	m.db.Begin();

}