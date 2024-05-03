package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func postData(dbURL string) http.HandlerFunc {
	db, _ := sql.Open("libsql", dbURL)

	return func(w http.ResponseWriter, r *http.Request) {
		data := SensorData{}

		fmt.Printf("Got a new request.\n")
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			w.WriteHeader(400)
			w.Write([]byte(fmt.Sprintf("Could not decode json : %s", err)))
			fmt.Println(err)
			return
		}

		fmt.Printf("Got data %v\n", data)
		if data.Id == "" {
			w.WriteHeader(400)
			w.Write([]byte("Empty sensor id."))
			return
		}

		
		fmt.Printf("Got data %v\n", data.TimeStamp)
		if data.TimeStamp == 0 {
			currentTime := time.Now()

			data.TimeStamp = currentTime.Unix()
		}
		fmt.Printf("Got data %v\n", data.TimeStamp)


		res, err := db.Exec(
			"INSERT INTO sensor_data VALUES (?,?,?,?,?);",
			data.Id,
			data.TimeStamp,
			data.Turbidity,
			data.Temperature,
			data.Ph,
		)
		fmt.Printf("exec", data.TimeStamp)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		fmt.Printf("exec", data.TimeStamp)

		rowsEffected, err := res.RowsAffected()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("%d rows effected", rowsEffected)))
		fmt.Printf("%d rows effected\n", rowsEffected)
		return
	}
}
