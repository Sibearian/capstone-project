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
	DO2          int    `json:"do2"`
	Turbidity    int    `json:"turbidity"`
	Temperature  int    `json:"temperature"`
	Ph           int    `json:"ph"`
}
