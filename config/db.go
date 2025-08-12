package config

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB(cfg Config) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error abriendo conexión:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos:", err)
	}

	log.Println("Conexión a MySQL exitosa")
}
