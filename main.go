package main

import "capstone/server"

var serverConfig = map[string]string{
	"TURSO_AUTH":   "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3MTMxOTkyMTksImlkIjoiNzNjMzJjNmYtMzhlNC00MTNlLTliZTgtMDM1M2VkNjYzZjRlIn0.oMGJmQe56pBTMt5auJja63N0QEr3i84qhbNJtvkxwdPwlz0l1RmB5C07_HVBVF7USdFvFMvweNubDOrS5rzxAg",
	"TURSO_DB_URL": "libsql://capstone-sibearian.turso.io",
}

func main() {
	server.Server(serverConfig)
}
