package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func wrapperPostData(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		data := SensorData{}
		var err error
		data.OrgId, err = getHeader("org-id", r)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		data.SensorNodeId, err = getHeader("sensor-id", r)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			w.WriteHeader(400)
			w.Write([]byte(fmt.Sprintf("Could not decode json : %s", err)))
			return
		}

		res, err := db.Exec(
			"INSERT INTO sensor_data VALUES (?,?,?,?,?,?,?);",
			data.OrgId,
			data.SensorNodeId,
			data.TimeStamp,
			data.DO2,
			data.Turbidity,
			data.Temperature,
			data.Ph,
		)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			sendError(err, 500, w)
			return
		}

		rowsEffected, err := res.RowsAffected()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("%d rows effected", rowsEffected)))
		return
	}
}
