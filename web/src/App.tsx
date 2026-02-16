// web/src/App.tsx
import { Routes, Route } from 'react-router-dom';
import { AppLayout } from './components/AppLayout';
import PatientsPage from './pages/Patients.tsx';
import DepartmentsPage from './pages/Departments.tsx';
import DiagnosesPage from './pages/Diagnoses';
import DoctorsPage from './pages/Doctors';
import HospitalizationsPage from './pages/Hospitalizations.tsx';
import HospitalizationReportPage from './pages/Report';


function App() {
  return (
    <AppLayout>
        <Routes>
            <Route path="/" element={<HospitalizationsPage />} />
            <Route path="/patients" element={<PatientsPage />} />
            <Route path="/doctors" element={<DoctorsPage />} />
            <Route path="/departments" element={<DepartmentsPage />} />
            <Route path="/diagnoses" element={<DiagnosesPage />} />
            <Route path="/reports/hospitalizations" element={<HospitalizationReportPage />} />
        </Routes>
    </AppLayout>
  );
}

export default App;