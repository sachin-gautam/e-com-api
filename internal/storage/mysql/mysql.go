package mysql

import (
	"database/sql"

	"github.com/sachin-gautam/go-crud-api/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("mysql", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
	id INTEGER PRIMARY KEY AUTO_INCREMENT, 
	name TEXT,
	email TEXT,
	age INTEGER
	 )`)

	if err != nil {
		return nil, err
	}
	return &Sqlite{
		Db: db,
	}, nil

}
