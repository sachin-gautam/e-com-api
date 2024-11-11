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

func (m *Mysql) GetList() ([]dtypes.Student, error) {
	stmt, err := m.Db.Prepare("SELECT id, name, email, age FROM students")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []dtypes.Student

	for rows.Next() {
		var student dtypes.Student

		if err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age); err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func (m *Mysql) UpdateById(id int64, name string, email string, age int) (dtypes.Student, error) {
	stmt, err := m.Db.Prepare("UPDATE students SET name = ?, email = ?, age = ? WHERE id = ?")
	if err != nil {
		return dtypes.Student{}, fmt.Errorf("prepare error: %w", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, email, age, id)
	if err != nil {
		return dtypes.Student{}, fmt.Errorf("execution error: %w", err)
	}

	return m.GetStudentById(id)
}
func (m *Mysql) DeleteById(id int64) (int64, error) {
	stmt, err := m.Db.Prepare("DELETE FROM students WHERE id = ?")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare delete statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return 0, fmt.Errorf("failed to execute delete statement: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return 0, fmt.Errorf("no student found with id %d", id)
	}

	return id, nil
}
