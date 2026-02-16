CREATE TABLE IF NOT EXISTS "patient" (
  "number_patient" SERIAL PRIMARY KEY,
  "first_name" text NOT NULL,
  "second_name" text NOT NULL,
  "sur_name" text,
  "gender" text NOT NULL,
  "date_of_birth" date NOT NULL,
  "patient_address" text NOT NULL
);

CREATE TABLE IF NOT EXISTS "position" (
    "id_position" SERIAL PRIMARY KEY,
    "title_position" TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS "diagnosis" (
  "id_diagnosis" SERIAL PRIMARY KEY,
  "title_diagnosis" text NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS "departament" (
  "id_departament" SERIAL PRIMARY KEY,
  "title_departament" text NOT NULL
);

CREATE TABLE IF NOT EXISTS "doctor" (
  "id_doctor" SERIAL PRIMARY KEY,
  "first_name" text NOT NULL,
  "second_name" text NOT NULL,
  "sur_name" text,
  "id_position" integer REFERENCES "position" ("id_position"),
  "id_departament" integer NOT NULL REFERENCES "departament" ("id_departament")
);

CREATE TABLE IF NOT EXISTS "hospitalization" (
  "id_hospitalization" SERIAL PRIMARY KEY,
  "number_patient" integer NOT NULL REFERENCES "patient" ("number_patient"),
  "id_doctor" integer NOT NULL REFERENCES "doctor" ("id_doctor"),
  "id_diagnosis" integer NOT NULL REFERENCES "diagnosis" ("id_diagnosis"),
  "id_departament" integer NOT NULL REFERENCES "departament" ("id_departament"),
  "start_hospitalization" date NOT NULL,
  "finish_hospitalization" date NOT NULL,
  "hospitalization_date" timestamp DEFAULT CURRENT_TIMESTAMP
);