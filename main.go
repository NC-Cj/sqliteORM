package main

import (
	"fmt"
	"os"

	"sqliteORM/session"
)

func main() {
	currentDir, _ := os.Getwd()
	fmt.Println("Current directory:", currentDir)
	session.NewSession()
	session := session.Session{}

	db, _ := session.Open("./database.db")
	defer session.Close()

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS test_table (id INTEGER PRIMARY KEY)")
	if err != nil {
		panic(err)
	}

}
