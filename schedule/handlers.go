package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getGroupsHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		groups, err := store.GetGroups(r.Context())
		if err != nil {
			http.Error(
				w,
				"Ошибка получения групп",
				http.StatusInternalServerError,
			)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(groups)
	}
}

func getTeachersHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		teachers, err := store.GetTeachers(r.Context())
		if err != nil {
			http.Error(
				w,
				"Ошибка получения преподавателей",
				http.StatusInternalServerError,
			)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(teachers)
	}
}

func getSubjectsHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		subjects, err := store.GetSubjects(r.Context())
		if err != nil {
			http.Error(
				w,
				"Ошибка получения предметов",
				http.StatusInternalServerError,
			)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(subjects)
	}
}

func getLessonsHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		lessons, err := store.GetLessons(r.Context())
		if err != nil {
			http.Error(
				w,
				"Ошибка получения занятий",
				http.StatusInternalServerError,
			)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(lessons)
	}
}

func createLessonHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var lesson Lesson

		err := json.NewDecoder(r.Body).Decode(&lesson)
		if err != nil {
			http.Error(
				w,
				"Неверный JSON",
				http.StatusBadRequest,
			)
			return
		}

		if lesson.GroupID <= 0 ||
			lesson.TeacherID <= 0 ||
			lesson.SubjectID <= 0 ||
			lesson.Room == "" ||
			lesson.StartsAt.IsZero() ||
			lesson.EndsAt.IsZero() {

			http.Error(
				w,
				"Некорректные данные занятия",
				http.StatusBadRequest,
			)
			return
		}

		if lesson.EndsAt.Before(lesson.StartsAt) ||
			lesson.EndsAt.Equal(lesson.StartsAt) {

			http.Error(
				w,
				"Время окончания должно быть позже времени начала",
				http.StatusBadRequest,
			)
			return
		}

		if lesson.Status == "" {
			lesson.Status = "planned"
		}

		if lesson.Status != "planned" &&
			lesson.Status != "done" &&
			lesson.Status != "cancelled" {

			http.Error(
				w,
				"Некорректный статус",
				http.StatusBadRequest,
			)
			return
		}

		id, err := store.CreateLesson(
			r.Context(),
			lesson,
		)

		if err != nil {
			http.Error(
				w,
				"Ошибка создания занятия",
				http.StatusInternalServerError,
			)
			return
		}

		lesson.ID = id

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(lesson)
	}
}

func deleteLessonHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := r.PathValue("id")

		var id int64

		_, err := fmt.Sscanf(idStr, "%d", &id)
		if err != nil || id <= 0 {
			http.Error(
				w,
				"Некорректный ID",
				http.StatusBadRequest,
			)
			return
		}

		err = store.DeleteLesson(r.Context(), id)
		if err != nil {
			http.Error(
				w,
				"Ошибка удаления занятия",
				http.StatusInternalServerError,
			)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}