package course

import (
	"database/sql"

	"github.com/Hharithsa/student-course-registration/entity"
)

type Repository interface {
	CreateCourses(course entity.Course) error
	QueryCourses() ([]*entity.Course, error)
	QueryCoursesByID(courseId int) (*entity.Course, error)
	UpdateCourses(course entity.Course) error
	DeleteCourses(courseID int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateCourses(course entity.Course) error {
	_, err := r.db.Exec("INSERT INTO courses (name, description) VALUES (?, ?)", course.Name, course.Description)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) QueryCourses() ([]*entity.Course, error) {
	rows, err := r.db.Query("SELECT * from courses")
	if err != nil {
		return nil, err
	}

	courses := make([]*entity.Course, 0)
	for rows.Next() {
		course, err := scanRowsIntoCourse(rows)
		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}

func (r *repository) QueryCoursesByID(courseId int) (*entity.Course, error) {
	rows, err := r.db.Query("SELECT * FROM courses WHERE id = ?", courseId)
	if err != nil {
		return nil, err
	}

	course := new(entity.Course)
	for rows.Next() {
		course, err = scanRowsIntoCourse(rows)
		if err != nil {
			return nil, err
		}
	}

	if course.ID == 0 {
		return nil, nil
	}
	return course, nil
}

func (r *repository) UpdateCourses(course entity.Course) error {
	_, err := r.db.Exec("UPDATE courses SET name = ?, description = ? WHERE id = ?", course.Name, course.Description, course.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteCourses(courseID int) error {
	_, err := r.db.Exec("DELETE FROM courses WHERE id = ?", courseID)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoCourse(rows *sql.Rows) (*entity.Course, error) {
	course := new(entity.Course)

	err := rows.Scan(
		&course.ID,
		&course.Name,
		&course.Description,
	)
	if err != nil {
		return nil, err
	}

	return course, nil
}
