// Anek database
package database

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// Get random anek from database
func GetRandomAnek() (string, error) {
	db, err := getDb()
	if err != nil {
		return "", err
	}
	defer db.Close()

	res, err := db.Query("select * from ANEKS order by RANDOM() limit 1;")
	if err != nil {
		return "", err
	}

	res.Next()

	var anek string
	err = res.Scan(&anek)
	if err != nil {
		return "", err
	}

	return anek, nil
}

// Get database instance. Unexported
func getDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "aneks/aneks.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Init database
func Init() error {
	_, err := os.Stat("aneks/aneks.db")
	if err != nil {
		return err
	}

	return nil
}
