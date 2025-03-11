package storage

import "github.com/Awais914/go-students-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int)(int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetAllStudents(limit int, page int) ([]types.Student, error)
}