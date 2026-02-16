// web/src/types/models.ts
export interface Patient {
  id?: number;
  first_name: string;
  last_name: string;
  sur_name?: string;
  gender: string;
  birth_date: string;
  address: string;
}

export interface Department {
  id?: number;
  title: string;
}

export interface Diagnosis {
  id?: number;
  title: string;
}

export interface Doctor {
  id?: number;
  first_name: string;
  last_name: string;
  sur_name?: string;
  departament_id: number;
}

export interface DoctorResponse extends Doctor {
  departament_title: string;
}

export interface Hospitalization {
    id?: number;
    patient_id: number;
    doctor_id: number;
    diagnosis_id: number;
    departament_id: number;
    start_date: string;
    finish_date: string;
}

export interface HospitalizationResponse {
    id: number;
    start_date: string;
    finish_date: string;
    patient: { id: number; first_name: string; last_name: string; };
    doctor: { id: number; first_name: string; last_name: string; };
    diagnosis: { id: number; title: string; };
    departament: { id: number; title: string; };
}

export interface Report {
  diagnosis: string;
  doctor_fio: string;
  patient_count: number;
  min_days: number;
  max_days: number;
}