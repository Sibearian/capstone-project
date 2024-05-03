package main

import "capstone/server"

var serverConfig = map[string]string{
	"TURSO_AUTH":   "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3MTQzODgxMjgsImlkIjoiNzNjMzJjNmYtMzhlNC00MTNlLTliZTgtMDM1M2VkNjYzZjRlIn0.cqpCIn2DobcmLen-dTaD3Joh4u-7d14a_7YuyjALfzFKC57nMWi0GLhtBQWfRpdJkLwDCaEKTYhaSiBuOne0Cg",
	"TURSO_DB_URL": "libsql://capstone-sibearian.turso.io?authToken=eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhIjoicnciLCJpYXQiOjE3MTQ1MDMyNTksImlkIjoiNzNjMzJjNmYtMzhlNC00MTNlLTliZTgtMDM1M2VkNjYzZjRlIn0.QONgvapYpJXCfzrvQONm591p_HbFQpd9yAVCrdIGFoboaHTbeT_n72U904PqyVVtZqMcqVeM6sqTktZQTSPeAA",
}

func main() {
	// Specify the address to listen on
	address := "0.0.0.0:8080"
	// Pass the address and serverConfig to the server
	server.RunServer(serverConfig, address)
}
