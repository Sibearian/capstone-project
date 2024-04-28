package server

type Organization struct {
	OrgId        string
	SensorNodeId string
	Sensors      int
}

type SensorData struct {
	OrgId        string `json:"org-id"`
	SensorNodeId string `json:"sensor-id"`
	TimeStamp    int64  `json:"timestamp"`
	DO2          int    `json:"DO2"`
	Turbidity    int    `json:"Turbidity"`
	Temperature  int    `json:"Temperature"`
	Ph           int    `json:"Ph"`
}
