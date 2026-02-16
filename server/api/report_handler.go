package api

import (
    "encoding/json"
    "hospital-app/server/storage"
    "net/http"
    "strconv"
)

type ReportHandler struct {
    repo *storage.ReportRepository
}

func NewReportHandler(repo *storage.ReportRepository) *ReportHandler {
    return &ReportHandler{repo: repo}
}

// GET /api/reports/hospitalizations?department_id=1&year=2024
func (h *ReportHandler) GetHospitalizationReport(w http.ResponseWriter, r *http.Request) {
    deptStr := r.URL.Query().Get("department_id")
    yearStr := r.URL.Query().Get("year")

    if deptStr == "" || yearStr == "" {
        http.Error(w, "department_id and year are required", http.StatusBadRequest)
        return
    }

    deptID, err := strconv.Atoi(deptStr)
    if err != nil {
        http.Error(w, "invalid department_id", http.StatusBadRequest)
        return
    }

    year, err := strconv.Atoi(yearStr)
    if err != nil {
        http.Error(w, "invalid year", http.StatusBadRequest)
        return
    }

    report, err := h.repo.GetHospitalizationReport(r.Context(), deptID, year)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(report)
}
