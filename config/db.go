package config

import (
	"database/sql"
	"log"
	_ "github.com/alexbrainman/odbc" // Importing the ODBC driver for database connection
)

var DB *sql.DB // Declaring a global variable DB of type *sql.DB

func ConnectDB(connString string) {
	var err error
	DB, err = sql.Open("odbc", connString) // Opening a database connection using the ODBC driver and connection string
	if err != nil {
		log.Fatalf("Error opening database: %v", err) // Logging a fatal error if the connection fails
	}
	err = DB.Ping() // Pinging the database to ensure the connection is valid
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err) // Logging a fatal error if the ping fails
	}

	log.Println("Conexción exitosa a la base de datos") // Logging a success message if the connection is successful
 }