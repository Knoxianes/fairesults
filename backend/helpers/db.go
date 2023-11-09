package helpers

import (
	"database/sql"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/libsql/libsql-client-go/libsql"
)

func DBConnect() (*sql.DB, error) {
	return sql.Open("libsql", os.Getenv("TURSODB"))
}
