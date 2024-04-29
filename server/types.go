package server

import (
	"database/sql"
	"sync"
)

type Organization struct {
	OrgId        string
	SensorNodeId string
	Sensors      int
}

type SensorData struct {
	OrgId        string
	SensorNodeId string
	TimeStamp    int64 `json:"timestamp"`
	DO2          int   `json:"do2"`
	Turbidity    int   `json:"turbidity"`
	Temperature  int   `json:"temperature"`
	Ph           int   `json:"ph"`
}

type DataBase struct {
	DB *sql.DB
	mu sync.Mutex
}
