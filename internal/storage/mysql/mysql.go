package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/imakhileshsahu/students-api/internal/config"
	types "github.com/imakhileshsahu/students-api/internal/type"
)

type MySQL struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*MySQL, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS student (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255),
			email VARCHAR(255),
			age INT
		)
	`)
	if err != nil {
		return nil, err
	}

	return &MySQL{Db: db}, nil
}

func (m *MySQL) CreateStudent(name, email string, age int) (int64, error) {
	res, err := m.Db.Exec("INSERT INTO student(name, email, age) VALUES (?, ?, ?)", name, email, age)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (m *MySQL) GetStudentById(id int64) (types.Student, error) {
	var student types.Student
	err := m.Db.QueryRow("SELECT id, name, email, age FROM student WHERE id = ?", id).
		Scan(&student.Id, &student.Name, &student.Email, &student.Age)
	if err != nil {
		return student, err
	}
	return student, nil
}

func (m *MySQL) GetStudents() ([]types.Student, error) {
	rows, err := m.Db.Query("SELECT id, name, email, age FROM student")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []types.Student
	for rows.Next() {
		var s types.Student
		if err := rows.Scan(&s.Id, &s.Name, &s.Email, &s.Age); err != nil {
			return nil, err
		}
		students = append(students, s)
	}
	return students, nil
}
