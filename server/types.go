package server

type Organization struct {
	Id   string
	Name string
}

type SensorData struct {
	Id          string `json:"id"`
	TimeStamp   int64    `json:"timestamp"`
	DO2         float64    `json:"do2"`
	Turbidity   float64    `json:"turbidity"`
	Temperature float64    `json:"temperature"`
	Ph          float64    `json:"ph"`
}
