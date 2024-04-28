package server

import "time"

type OrgId string
type SData int

type Organization struct {
	OrgId
	SensorNodeId string
	Sensors      int
}

type SensorData struct {
	OrgId
	SensorNodeId string
	TimeStamp    time.Time
	Sensor1      SData
	Sensor2      SData
	Sensor3      SData
	Sensor4      SData
	Sensor5      SData
	Sensor6      SData
}
