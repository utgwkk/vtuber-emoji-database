package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func handleTopPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>It works!</h1>")
	db, err := sql.Open("sqlite3", "production.db")
	if err != nil {
		panic(err)
	}
	rows, err := db.Query(`select emoji from emojis order by id desc limit 5`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var emoji string
		if err := rows.Scan(&emoji); err != nil {
			log.Fatal("rows.Scan()", err)
			return
		}
		fmt.Fprintf(w, "<p>%s</p>\n", emoji)
	}
}

func main() {
	http.HandleFunc("/", handleTopPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
