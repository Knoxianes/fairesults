package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/libsql/libsql-client-go/libsql"
)

var DB *sql.DB

func InitDB() {
	dbUrl := os.Getenv("TURSOURL") + "?authToken=" + os.Getenv("TURSOTOKEN")
	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		log.Panic(err)
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	DB = db
}

func CloseDB() error {
	return DB.Close()
}

func GetLastInsertedIDFromUsers() (int, error){
	var id int
	row := DB.QueryRow("select rowid from users order by rowid desc limit 1;")
	err := row.Scan(&id)
	if err != nil{
		log.Println(err)
		return -1, err
	}
	return id, nil
}

func GetLastInsertedIDFromResults() (int, error){
	var id int
	row := DB.QueryRow("select rowid from results order by rowid desc limit 1;")
	err := row.Scan(&id)
	if err != nil{
		log.Println(err)
		return -1, err
	}
	return id, nil
}
