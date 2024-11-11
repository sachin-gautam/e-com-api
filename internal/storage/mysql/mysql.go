package mysql

import (
	"database/sql"
	"fmt"

	"github.com/sachin-gautam/go-crud-api/internal/config"
	"github.com/sachin-gautam/go-crud-api/internal/dtypes"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Mysql, error) {
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
	return &Mysql{
		Db: db,
	}, nil

}

func (m *Mysql) CreateStudent(name string, email string, age int) (int64, error) {

	stmt, err := m.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (m *Mysql) GetStudentById(id int64) (dtypes.Student, error) {
	stmt, err := m.Db.Prepare("SELECT * FROM students WHERE id = ? LIMIT 1")
	if err != nil {
		return dtypes.Student{}, err
	}

	defer stmt.Close()

	var student dtypes.Student

	if err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age); err != nil {
		if err == sql.ErrNoRows {
			return dtypes.Student{}, fmt.Errorf("no studen with id %s ", fmt.Sprint(id))
		}
		return dtypes.Student{}, fmt.Errorf("query error: %w", err)
	}
	return student, nil
}
