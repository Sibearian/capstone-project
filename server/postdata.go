package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func StoreSensorData(w http.ResponseWriter, r *http.Request) {
	data := SensorData{}
	data.OrgId = r.Header.Get("org-id")
	if data.OrgId == "" {
		w.WriteHeader(400)
		w.Write([]byte("Missing org-id in the header."))
		return
	}

	data.SensorNodeId = r.Header.Get("sensor-id")
	if data.SensorNodeId == "" {
		w.WriteHeader(400)
		w.Write([]byte("Missing org-id in the header."))
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("Could not decode json : %s", err)))
		return
	}

	fmt.Printf("%v\n", data)
}
