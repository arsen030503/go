package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go_crud/database"
	"go_crud/models"
	"net/http"
	"strconv"
)

func GetEntries(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, user_id, content, created_at FROM entries")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var entries []models.Entry
	for rows.Next() {
		var entry models.Entry
		err := rows.Scan(&entry.ID, &entry.UserID, &entry.Content, &entry.CreatedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		entries = append(entries, entry)
	}
	json.NewEncoder(w).Encode(entries)
}

func CreateEntry(w http.ResponseWriter, r *http.Request) {
	var entry models.Entry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := database.DB.QueryRow("INSERT INTO entries (user_id, content) VALUES ($1, $2) RETURNING id, created_at",
		entry.UserID, entry.Content).Scan(&entry.ID, &entry.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(entry)
}
func UpdateEntry(w http.ResponseWriter, r *http.Request) {
	// Извлекаем ID записи из URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Декодируем данные из запроса
	var entry models.Entry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Обновляем запись в базе данных
	_, err = database.DB.Exec("UPDATE entries SET content=$1 WHERE id=$2", entry.Content, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем успешный ответ
	entry.ID = id
	json.NewEncoder(w).Encode(entry)
}
func DeleteEntry(w http.ResponseWriter, r *http.Request) {
	// Извлекаем ID записи из URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Удаляем запись из базы данных
	_, err = database.DB.Exec("DELETE FROM entries WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusNoContent)
}
