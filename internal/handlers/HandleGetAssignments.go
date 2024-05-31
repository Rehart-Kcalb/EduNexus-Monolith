package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleGetAssignments(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		course_id, err := DB.GetCourseId(context.Background(), r.PathValue("course_name"))
		if err != nil {
			log.Println(err)
		}
		assignments, err := DB.GetAssignments(context.Background(), course_id)
		if err != nil {
			log.Println(err)
		}
		types.NewJsonResponse(struct {
			Assignments any `json:"assignments"`
		}{assignments}, http.StatusOK).Respond(w)
	}
}