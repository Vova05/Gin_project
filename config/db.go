package config
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/mattn/go-sqlite3"
)
var DB *sql.DB
func InitializeDB() (*sql.DB, error) {
	// Initialize connection to the database
	var err error
	DB, err = sql.Open("mysql", "root:1234@/Kurs5")
	return DB, err
}