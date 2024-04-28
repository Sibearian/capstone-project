package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
)

func Server(config map[string]string) {
	mux := http.NewServeMux()
	ctx := context.Background()
	db, err := sql.Open("libsql", config["TURSO_DB_URL"])
	if err != nil {
		log.Fatalf("failed to opendb %s: %s", config["TURSO_DB_URL"], err)
		os.Exit(1)
	}

	<-ctx.Done()
	defer db.Close()
}
