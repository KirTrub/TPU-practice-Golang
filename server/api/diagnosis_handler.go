package api

import (
	"encoding/json"
	"hospital-app/server/models"
	"hospital-app/server/storage"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DiagnosisHandler struct {
	repo *storage.DiagnosisRepository
}

func NewDiagnosisHandler(repo *storage.DiagnosisRepository) *DiagnosisHandler {
	return &DiagnosisHandler{repo: repo}
}

// POST /api/diagnoses
func (h *DiagnosisHandler) CreateDiagnosisHandler(w http.ResponseWriter, r *http.Request) {
	var diag models.Diagnosis
	if err := json.NewDecoder(r.Body).Decode(&diag); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := h.repo.CreateDiagnosis(r.Context(), &diag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int32{"id": id})
}

// GET /api/diagnoses
func (h *DiagnosisHandler) GetAllDiagnosesHandler(w http.ResponseWriter, r *http.Request) {
	diags, err := h.repo.GetAllDiagnoses(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(diags)
}

// PUT /api/Diagnosis/{id}
func (h *DiagnosisHandler) UpdateDiagnoseHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var dep models.Diagnosis
	if err := json.NewDecoder(r.Body).Decode(&dep); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dep.ID = int32(id)
	if err := h.repo.UpdateDiagnose(r.Context(), &dep); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DELETE /api/diagnoses/{id}
func (h *DiagnosisHandler) DeleteDiagnoseHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := h.repo.DeleteDiagnose(r.Context(), int32(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
