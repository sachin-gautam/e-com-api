package storage

import "github.com/sachin-gautam/go-crud-api/internal/dtypes"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (dtypes.Student, error)
}
