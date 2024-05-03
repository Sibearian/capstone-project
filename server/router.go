package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func RunServer(config map[string]string, address string) {
	mux := http.NewServeMux()
	db, err := sql.Open("libsql", config["TURSO_DB_URL"])
	if err != nil {
		log.Fatalf("failed to opendb %s: %s", config["TURSO_DB_URL"], err)
		os.Exit(1)
	}

	mux.HandleFunc("GET /", heartbeat)
	mux.HandleFunc("POST /api/data", postData(config["TURSO_DB_URL"]))

	http.ListenAndServe(address, mux)

	defer db.Close()
}
