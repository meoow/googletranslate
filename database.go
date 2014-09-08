package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbfile        string = ".googletranslate.db"
	createTblStmt string = `create table if not exists gt (words text not null primary key, trans text);`
)

func GetTrans(text string) (string, error) {
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		return "", err
	}
	defer db.Close()

	stmt, err := db.Prepare("select trans from gt where words = ?")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	var name string
	err = stmt.QueryRow(text).Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

func SaveTrans(words, trans string) error {
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(createTblStmt)
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into gt(words, trans) values(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(words, trans)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}
