package handlers

import (
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
)

func HandleGetModuleProgress(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO:
	}
}
