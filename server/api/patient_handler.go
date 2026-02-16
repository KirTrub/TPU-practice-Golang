package api

import (
	"encoding/json"
	"hospital-app/server/models"
	"hospital-app/server/storage"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PatientHandler struct {
	repo *storage.PatientRepository
}

func NewPatientHandler(repo *storage.PatientRepository) *PatientHandler {
	return &PatientHandler{repo: repo}
}

// POST /api/patients
func (h *PatientHandler) CreatePatientHandler(w http.ResponseWriter, r *http.Request) {
	var patient models.Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.repo.CreatePatient(r.Context(), &patient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int32{"id": id})
}

// GET /api/patients
func (h *PatientHandler) GetAllPatientsHandler(w http.ResponseWriter, r *http.Request) {
	patients, err := h.repo.GetAllPatients(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patients)
}

// GET /api/patients/{id}
func (h *PatientHandler) GetPatientByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	patient, err := h.repo.GetPatientByID(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Patient not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(patient)
}

// PUT /api/patients/{id}
func (h *PatientHandler) UpdatePatientHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	var patient models.Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	patient.ID = int32(id)
	if err := h.repo.UpdatePatient(r.Context(), &patient); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DELETE /api/patients/{id}
func (h *PatientHandler) DeletePatientHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	if err := h.repo.DeletePatient(r.Context(), int32(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
