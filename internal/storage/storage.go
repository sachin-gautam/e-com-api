package storage

import "github.com/sachin-gautam/go-crud-api/internal/model"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (model.Student, error)
	GetList() ([]model.Student, error)
	UpdateById(id int64, name, email string, age int) (model.Student, error)
	DeleteById(id int64) (int64, error)
}
