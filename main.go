package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	dataSourceName := os.Getenv("HAKARU_DATASOURCENAME")
	if dataSourceName == "" {
		dataSourceName = "root:password@tcp(127.0.0.1:13306)/hakaru"
	}

	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err.Error())
		}
	}(db)

	// See "Important settings" section.
	//db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(30)

	http.HandleFunc("/hakaru", hakaruHandler)
	http.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200) })

	// start server
	log.Println("Listen and Serve")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}

func hakaruHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tx, err := db.BeginTx(ctx, nil)

	// Retrieve get parameters and create insert query
	name := r.URL.Query().Get("name")
	value := r.URL.Query().Get("value")
	log.Println(name, value)
	// Execute query
	_, err = db.ExecContext(ctx, "INSERT INTO `eventlog` (`at`, `name`, `value`) VALUES (NOW(), ?, ?)", name, value)
	if err != nil {
		tx.Rollback()
		log.Println(err)
	}
	tx.Commit()
	/*
		origin := r.Header.Get("Origin")
		if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
	*/
}
