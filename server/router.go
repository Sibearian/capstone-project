package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func Server(config map[string]string) {
	mux := http.NewServeMux()
	ctx := context.Background()
	db, err := sql.Open("libsql", config["TURSO_DB_URL"])
	if err != nil {
		log.Fatalf("failed to opendb %s: %s", config["TURSO_DB_URL"], err)
		os.Exit(1)
	}

	mux.HandleFunc("POST /", wrapperPostData(db))

	http.ListenAndServe(":3000", mux)

	<-ctx.Done()
	defer db.Close()
}
