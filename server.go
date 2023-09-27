package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handlePing(w http.ResponseWriter, r *http.Request) {
	log.Println("ping request received")
	io.WriteString(w, "pong")
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	log.Println("get request")
	io.WriteString(w, "this endpoint does nothing")
}

type InsertParams struct {
	key   string
	value string
}

func handleInsert(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	databaseName := params["database"]
	log.Println("insert request, database:", databaseName)
	database, ok := databases[databaseName]
	if !ok {
		// return early and send error
		log.Println("ERROR: datbase doesnt exist:", databaseName)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, fmt.Sprint("ERROR: database doesnt exist: ", databaseName))
		return
	}

	var insertParams InsertParams
	err := json.NewDecoder(r.Body).Decode(insertParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "ERROR: JSON params incorrect")
		return
	}

	err = database.Insert(insertParams.key, insertParams.value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, fmt.Sprint("ERROR: could not insert key/value pair:", err))
		return
	}

	io.WriteString(w, "this endpoint does nothing")
}

func handleCreate(w http.ResponseWriter, r *http.Request) {
	log.Println("create request")
	io.WriteString(w, "this endpoint does nothing")
}

func StartServer() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/api/{database:[a-z]+}/insert", handleInsert)
	rtr.HandleFunc("/api/{database:[a-z]+}/get", handleGet)
	rtr.HandleFunc("/api/{database:[a-z]+}/insert", handleInsert)
	rtr.HandleFunc("/api/create", handleCreate)
	rtr.HandleFunc("/ping", handlePing)

	http.Handle("/", rtr)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
