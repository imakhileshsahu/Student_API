// package storage

// import "github.com/imakhileshsahu/students-api/internal/type"

// type Storage interface {
// 	CreateStudent(name string, email string, age int) (int64, error)
// 	GetStudentById(id int64) (types.Student, error)
// 	GetStudents() ([]types.Student, error)
// }

package storage

import "github.com/imakhileshsahu/students-api/internal/type"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
}
