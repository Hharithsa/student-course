package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Hharithsa/student-course-registration/cmd/api"
	"github.com/Hharithsa/student-course-registration/config"
	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := newDBConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = checkConnection(db)
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(fmt.Sprintf(":%s", config.Envs.Port), db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func newDBConnection(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func checkConnection(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}

	log.Println("DB: Successfully connected!")
	return nil
}
