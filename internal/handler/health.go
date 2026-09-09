package handler

import (
	"net/http"
)

func Health(w http.ResponseWriter, _ *http.Request) {
	respondJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func Ready(w http.ResponseWriter, _ *http.Request) {
	// TODO: проверить готовность зависимостей: БД, кеш, внешние сервисы.
	respondJSON(w, http.StatusOK, map[string]string{"status": "ready"})
}
