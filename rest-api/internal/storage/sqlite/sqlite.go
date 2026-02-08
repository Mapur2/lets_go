package sqlite

import (
	"database/sql"

	"github.com/mapur2/lets_go/rest-apis/internal/config"
	"github.com/mapur2/lets_go/rest-apis/internal/types"
	_ "modernc.org/sqlite"
)

type SQlite struct {
	Db *sql.DB
}

func (s *SQlite) GetStudentByEmail(email string) (bool, types.Student, error) {
	var student types.Student

	err := s.Db.QueryRow(
		"SELECT id, name, age, email FROM students WHERE email = ?",
		email,
	).Scan(
		&student.Id,
		&student.Name,
		&student.Age,
		&student.Email,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, types.Student{}, err
		}
		return false, types.Student{}, err
	}

	return true, student, nil
}

func (s *SQlite) CreateStudent(student types.Student) (int, error) {
	statement, err := s.Db.Prepare("INSERT INTO students (name,email, age) VALUES(?,?,?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	res, errs := statement.Exec(student.Name, student.Email, student.Age)
	if errs != nil {
		return 0, errs
	}
	lastId, er := res.LastInsertId()
	if er != nil {
		return 0, errs
	}
	return int(lastId), nil
}

func New(cfg *config.Config) (*SQlite, error) {

	db, err := sql.Open("sqlite", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER,
		email TEXT
	)
	`)
	if err != nil {
		return nil, err
	}

	return &SQlite{
		Db: db,
	}, nil
}
