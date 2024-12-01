package registration

import (
	"database/sql"
	"time"

	"github.com/Hharithsa/student-course-registration/entity"
)

type Repository interface {
	CreateRegistration(entity.Registration) error
	QueryRegistrations() ([]*entity.Registration, error)
	QueryRegistrationsByStudentID(studentID int) ([]*entity.Registration, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateRegistration(registration entity.Registration) error {
	_, err := r.db.Exec("INSERT INTO registrations (student_id, course_id, created_at) VALUES (?, ?, ?)", registration.StudentID, registration.CourseID, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) QueryRegistrations() ([]*entity.Registration, error) {
	rows, err := r.db.Query("SELECT * from registrations")
	if err != nil {
		return nil, err
	}

	registrations := make([]*entity.Registration, 0)
	for rows.Next() {
		registration, err := scanRowsIntoRegistration(rows)
		if err != nil {
			return nil, err
		}

		registrations = append(registrations, registration)
	}

	return registrations, nil
}

func (r *repository) QueryRegistrationsByStudentID(studentID int) ([]*entity.Registration, error) {
	rows, err := r.db.Query("SELECT * from registrations WHERE student_id = ?", studentID)
	if err != nil {
		return nil, err
	}

	registrations := make([]*entity.Registration, 0)
	for rows.Next() {
		registration, err := scanRowsIntoRegistration(rows)
		if err != nil {
			return nil, err
		}

		registrations = append(registrations, registration)
	}

	return registrations, nil
}

func scanRowsIntoRegistration(rows *sql.Rows) (*entity.Registration, error) {
	registration := new(entity.Registration)

	err := rows.Scan(
		&registration.StudentID,
		&registration.CourseID,
		&registration.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return registration, nil
}
