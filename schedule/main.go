package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	db, err := pgxpool.New(
		ctx,
		"postgres://postgres:postgres@localhost:5432/schedule",
	)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		log.Fatal("Ошибка подключения к базе:", err)
	}

	log.Println("Подключение к PostgreSQL успешно!")

	store := New(db)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /groups", getGroupsHandler(store))
	mux.HandleFunc("GET /teachers", getTeachersHandler(store))
	mux.HandleFunc("GET /subjects", getSubjectsHandler(store))
	mux.HandleFunc("GET /lessons", getLessonsHandler(store))

	mux.HandleFunc("POST /lessons", createLessonHandler(store))

	mux.HandleFunc("DELETE /lessons/{id}", deleteLessonHandler(store))

	log.Println("Сервер запущен на http://localhost:8080")

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

err = http.ListenAndServe(":8080", withCORS(mux))