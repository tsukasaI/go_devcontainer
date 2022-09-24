package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	openDB()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8001", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func openDB() {
	_, err := sql.Open("mysql", "go_db:go_db@/go_db")
	if err != nil {
		panic(err)
	}
	fmt.Println("DB opend")
}
