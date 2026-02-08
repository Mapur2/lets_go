package storage

import (
	"github.com/mapur2/lets_go/rest-apis/internal/types"
)

type Storage interface {
	CreateStudent(student types.Student) (int, error)
	GetStudentByEmail(email string) (bool,types.Student, error)
}
