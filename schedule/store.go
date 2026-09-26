package main

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetGroups(ctx context.Context) ([]Group, error) {
	rows, err := s.db.Query(ctx, `
		SELECT id, name, year
		FROM groups
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []Group

	for rows.Next() {
		var group Group

		if err := rows.Scan(
			&group.ID,
			&group.Name,
			&group.Year,
		); err != nil {
			return nil, err
		}

		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

func (s *Store) GetTeachers(ctx context.Context) ([]Teacher, error) {
	rows, err := s.db.Query(ctx, `
		SELECT id, full_name, email
		FROM teachers
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teachers []Teacher

	for rows.Next() {
		var teacher Teacher

		if err := rows.Scan(
			&teacher.ID,
			&teacher.FullName,
			&teacher.Email,
		); err != nil {
			return nil, err
		}

		teachers = append(teachers, teacher)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return teachers, nil
}

func (s *Store) GetSubjects(ctx context.Context) ([]Subject, error) {
	rows, err := s.db.Query(ctx, `
		SELECT id, name
		FROM subjects
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subjects []Subject

	for rows.Next() {
		var subject Subject

		if err := rows.Scan(
			&subject.ID,
			&subject.Name,
		); err != nil {
			return nil, err
		}

		subjects = append(subjects, subject)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return subjects, nil
}

func (s *Store) GetLessons(ctx context.Context) ([]Lesson, error) {
	rows, err := s.db.Query(ctx, `
		SELECT
			id,
			group_id,
			teacher_id,
			subject_id,
			room,
			starts_at,
			ends_at,
			status,
			created_at
		FROM lessons
		ORDER BY starts_at
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []Lesson

	for rows.Next() {
		var lesson Lesson

		if err := rows.Scan(
			&lesson.ID,
			&lesson.GroupID,
			&lesson.TeacherID,
			&lesson.SubjectID,
			&lesson.Room,
			&lesson.StartsAt,
			&lesson.EndsAt,
			&lesson.Status,
			&lesson.CreatedAt,
		); err != nil {
			return nil, err
		}

		lessons = append(lessons, lesson)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lessons, nil
}

func (s *Store) CreateLesson(
	ctx context.Context,
	lesson Lesson,
) (int64, error) {

	var id int64

	err := s.db.QueryRow(ctx, `
		INSERT INTO lessons (
			group_id,
			teacher_id,
			subject_id,
			room,
			starts_at,
			ends_at,
			status
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`,
		lesson.GroupID,
		lesson.TeacherID,
		lesson.SubjectID,
		lesson.Room,
		lesson.StartsAt,
		lesson.EndsAt,
		lesson.Status,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Store) DeleteLesson(ctx context.Context, id int64) error {
	_, err := s.db.Exec(ctx, `
		DELETE FROM lessons
		WHERE id = $1
	`, id)

	return err
}