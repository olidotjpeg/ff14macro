package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

type Macro struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var Database *sql.DB

func PostMacro(w http.ResponseWriter, r *http.Request) {
	var macro Macro
	macro.Id = uuid.NewString()

	decoder := json.NewDecoder(r.Body)
	decodeError := decoder.Decode(&macro)
	if decodeError != nil {
		panic(decodeError)
	}

	sqlQuery := fmt.Sprintf(`INSERT INTO macros(id, name, description) VALUES('%s', '%s', '%s')`, macro.Id, macro.Name, macro.Description)

	var err error

	_, err = Database.Exec(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(macro)
}

func GetMacro(w http.ResponseWriter, r *http.Request) {
	routeVariables := mux.Vars(r)
	key := routeVariables["id"]

	sqlQuery := fmt.Sprintf(`SELECT * FROM macros WHERE id = '%s'`, key)

	row := Database.QueryRow(sqlQuery)
	var macro Macro
	err := row.Scan(&macro.Name, &macro.Description)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(macro)
}

func GetMacros(w http.ResponseWriter, r *http.Request) {
	rows, err := Database.Query("SELECT * FROM macros")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var macros []Macro

	for rows.Next() {
		var macro Macro
		err := rows.Scan(&macro.Id, &macro.Name, &macro.Description)

		if err != nil {
			println(fmt.Sprint(err))
			return
		}
		macros = append(macros, macro)
	}

	json.NewEncoder(w).Encode(macros)
}

func main() {
	var err error

	Database, err = sql.Open("sqlite3", "./macro.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer Database.Close()

	r := mux.NewRouter()
	r.HandleFunc("/macro", PostMacro).Methods(http.MethodPost)
	r.HandleFunc("/macros", GetMacros).Methods(http.MethodGet)
	r.HandleFunc("/macro/{id}", GetMacro).Methods(http.MethodGet)
	http.Handle("/", r)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost", "http://localhost:5173"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	srv := &http.Server{
		Handler: c.Handler(r),
		// Handler: r,
		Addr: "0.0.0.0:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
