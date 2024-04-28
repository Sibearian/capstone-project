package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func StoreSensorData(w http.ResponseWriter, r *http.Request) {
	data := SensorData{}
	org_id, ok := r.Header["org-id"]
	if !ok {
		w.WriteHeader(400)
		w.Write([]byte("Missing org-id in the header."))
		return
	}
	data.OrgId = org_id[0]

	sensor_id, ok := r.Header["sensor-id"]
	if !ok {
		w.WriteHeader(400)
		w.Write([]byte("Missing org-id in the header."))
		return
	}
	data.SensorNodeId = sensor_id[0]

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("Could not decode json", err)))
		return
	}

	fmt.Printf("%v", data)
}
