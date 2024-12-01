package student

import (
	"database/sql"

	"github.com/Hharithsa/student-course-registration/entity"
)

type Repository interface {
	CreateStudents(student entity.Student) error
	QueryStudents() ([]*entity.Student, error)
	QueryStudentsByID(studentId int) (*entity.Student, error)
	UpdateStudents(student entity.Student) error
	DeleteStudents(studentID int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateStudents(student entity.Student) error {
	_, err := r.db.Exec("INSERT INTO students (name, age, college, year) VALUES (?, ?, ?, ?)", student.Name, student.Age, student.College, student.Year)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) QueryStudents() ([]*entity.Student, error) {
	rows, err := r.db.Query("SELECT * from students")
	if err != nil {
		return nil, err
	}

	students := make([]*entity.Student, 0)
	for rows.Next() {
		student, err := scanRowsIntoStudent(rows)
		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}

func (r *repository) QueryStudentsByID(studentId int) (*entity.Student, error) {
	rows, err := r.db.Query("SELECT * FROM students WHERE id = ?", studentId)
	if err != nil {
		return nil, err
	}

	student := new(entity.Student)
	for rows.Next() {
		student, err = scanRowsIntoStudent(rows)
		if err != nil {
			return nil, err
		}
	}

	if student.ID == 0 {
		return nil, nil
	}
	return student, nil
}

func (r *repository) UpdateStudents(student entity.Student) error {
	_, err := r.db.Exec("UPDATE students SET name = ?, age = ?, college = ?, year = ? WHERE id = ?", student.Name, student.Age, student.College, student.Year, student.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteStudents(studentID int) error {
	_, err := r.db.Exec("DELETE FROM students WHERE id = ?", studentID)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoStudent(rows *sql.Rows) (*entity.Student, error) {
	student := new(entity.Student)

	err := rows.Scan(
		&student.ID,
		&student.Name,
		&student.Age,
		&student.College,
		&student.Year,
	)
	if err != nil {
		return nil, err
	}

	return student, nil
}
