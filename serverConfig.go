package main

type dbConfig struct {
	TURSO_AUTH   string
	TURSO_DB_URL string
}

var dbServerConfig = dbConfig{
	TURSO_AUTH:   "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3MTMxOTkyMTksImlkIjoiNzNjMzJjNmYtMzhlNC00MTNlLTliZTgtMDM1M2VkNjYzZjRlIn0.oMGJmQe56pBTMt5auJja63N0QEr3i84qhbNJtvkxwdPwlz0l1RmB5C07_HVBVF7USdFvFMvweNubDOrS5rzxAg",
	TURSO_DB_URL: "libsql://capstone-sibearian.turso.io",
}
