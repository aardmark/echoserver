package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type user struct {
	ID             int
	firstname      string
	lastname       string
	email          string
	hashedPassword string
}

func getUserByEmail(email string) (user, error) {
	var ret user

	db, err := sql.Open("sqlite3", "./echoserver.db")
	if err != nil {
		log.Fatal(err)
		return ret, err
	}
	defer db.Close()

	stmt, err := db.Prepare("select ID, Firstname, Lastname, Email, Password from users where email = ?")
	if err != nil {
		log.Fatal(err)
		return ret, err
	}
	defer stmt.Close()

	ret = user{}

	err = stmt.QueryRow(email).Scan(&ret.ID, &ret.firstname, &ret.lastname, &ret.email, &ret.hashedPassword)
	if err != nil {
		log.Fatal(err)
		return ret, err
	}
	fmt.Println(ret)

	return ret, nil
}
