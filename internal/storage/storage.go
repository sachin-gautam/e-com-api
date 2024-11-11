package storage

import "github.com/sachin-gautam/go-crud-api/internal/dtypes"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (dtypes.Student, error)
	GetList() ([]dtypes.Student, error)
	UpdateById(id int64, name, email string, age int) (dtypes.Student, error)
	DeleteById(id int64) (int64, error)
}
