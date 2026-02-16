package api

import (
	"encoding/json"
	"hospital-app/server/models"
	"hospital-app/server/storage"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DoctorHandler struct {
	repo *storage.DoctorRepository
}

func NewDoctorHandler(repo *storage.DoctorRepository) *DoctorHandler {
	return &DoctorHandler{repo: repo}
}

// POST /api/doctors
func (h *DoctorHandler) CreateDoctorHandler(w http.ResponseWriter, r *http.Request) {
	var doc models.Doctor
	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := h.repo.CreateDoctor(r.Context(), &doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int32{"id": id})
}

// GET /api/doctors
func (h *DoctorHandler) GetAllDoctorsHandler(w http.ResponseWriter, r *http.Request) {
	docs, err := h.repo.GetAllDoctors(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(docs)
}

// PUT /api/doctors/{id}
func (h *DoctorHandler) UpdateDoctorHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var doc models.Doctor
	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	doc.ID = int32(id)
	if err := h.repo.UpdateDoctor(r.Context(), &doc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DELETE /api/doctors/{id}
func (h *DoctorHandler) DeleteDoctorHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := h.repo.DeleteDoctor(r.Context(), int32(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
