package api

import (
	"encoding/json"
	"fmt"
	"hospital-app/server/models"
	"hospital-app/server/storage"
	"net/http"

	"github.com/gorilla/mux"
)

type HospitalizationHandler struct {
	repo *storage.HospitalizationRepository
}

func NewHospitalizationHandler(repo *storage.HospitalizationRepository) *HospitalizationHandler {
	return &HospitalizationHandler{repo: repo}
}

// POST /api/hospitalizations
func (h *HospitalizationHandler) CreateHospitalizationHandler(w http.ResponseWriter, r *http.Request) {
	var hosp models.Hospitalization
	if err := json.NewDecoder(r.Body).Decode(&hosp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := h.repo.CreateHospitalization(r.Context(), &hosp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int32{"id": id})
}

// GET /api/hospitalizations
func (h *HospitalizationHandler) GetAllHospitalizationsHandler(w http.ResponseWriter, r *http.Request) {
	hosps, err := h.repo.GetAllHospitalizations(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(hosps)
}

// DELETE /api/hospitalizations/{id}
func (h *HospitalizationHandler) DeleteHospitalizationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}
	var id int32
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		http.Error(w, "invalid id parameter", http.StatusBadRequest)
		return
	}
	err := h.repo.DeleteHospitalization(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
