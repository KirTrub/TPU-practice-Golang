package api

import (
	"encoding/json"
	"hospital-app/server/models"
	"hospital-app/server/storage"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DepartamentHandler struct {
	repo *storage.DepartamentRepository
}

func NewDepartamentHandler(repo *storage.DepartamentRepository) *DepartamentHandler {
	return &DepartamentHandler{repo: repo}
}

// POST /api/departments
func (h *DepartamentHandler) CreateDepartamentHandler(w http.ResponseWriter, r *http.Request) {
	var dep models.Departament
	if err := json.NewDecoder(r.Body).Decode(&dep); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := h.repo.CreateDepartament(r.Context(), &dep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int32{"id": id})
}

// GET /api/departments
func (h *DepartamentHandler) GetAlldepartmentsHandler(w http.ResponseWriter, r *http.Request) {
	deps, err := h.repo.GetAlldepartments(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(deps)
}

// PUT /api/departaments/{id}
func (h *DepartamentHandler) UpdateDepartamentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var dep models.Departament
	if err := json.NewDecoder(r.Body).Decode(&dep); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dep.ID = int32(id)
	if err := h.repo.UpdateDepartament(r.Context(), &dep); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DELETE /api/departaments/{id}
func (h *DepartamentHandler) DeleteDepartamentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := h.repo.DeleteDepartament(r.Context(), int32(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
