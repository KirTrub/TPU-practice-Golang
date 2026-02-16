import { useEffect, useState } from 'react';
import {
  Table,
  Button,
  Group,
  Title,
  Loader,
  Center,
  Text,
  NumberInput,
  Select,
} from '@mantine/core';
import { apiClient } from '../api/client';
import type { HospitalizationReport, Department } from '../types/models';

function HospitalizationReportPage() {
  const currentYear = new Date().getFullYear();

  const [departments, setDepartments] = useState<Department[]>([]);
  const [departmentId, setDepartmentId] = useState<string | null>(null);
  const [year, setYear] = useState<number>(currentYear);

  const [report, setReport] = useState<HospitalizationReport[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchDepartments = async () => {
    const resp = await apiClient.get<Department[]>('/departments');
    setDepartments(resp.data ?? []);
    if (resp.data?.length) {
      setDepartmentId(String(resp.data[0].id));
    }
  };

  const fetchReport = async () => {
    if (!departmentId) return;

    try {
      setLoading(true);
      const resp = await apiClient.get<HospitalizationReport[]>(
        `/reports/hospitalizations?department_id=${departmentId}&year=${year}`
      );
      setReport(resp.data ?? []);
      setError(null);
    } catch {
      setError('Не удалось сформировать отчёт');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchDepartments();
  }, []);

  return (
    <>
      <Group justify="space-between" mb="md">
        <Title order={2}>Отчёт о госпитализациях отделения</Title>

        <Group>
          <Select
            label="Отделение"
            data={departments.map(d => ({ value: String(d.id), label: d.title }))}
            value={departmentId}
            onChange={setDepartmentId}
            w={260}
          />

          <NumberInput
            label="Год"
            value={year}
            onChange={(v) => setYear(Number(v))}
            min={2000}
            max={currentYear}
            w={120}
          />

          <Button onClick={fetchReport}>Сформировать</Button>
        </Group>
      </Group>

      {loading ? (
        <Center><Loader /></Center>
      ) : error ? (
        <Center><Text c="red">{error}</Text></Center>
      ) : (
        <Table highlightOnHover>
          <Table.Thead>
            <Table.Tr>
              <Table.Th>Диагноз</Table.Th>
              <Table.Th>Врач</Table.Th>
              <Table.Th>Кол-во пациентов</Table.Th>
              <Table.Th>Мин. срок</Table.Th>
              <Table.Th>Макс. срок</Table.Th>
            </Table.Tr>
          </Table.Thead>
          <Table.Tbody>
            {report.length ? report.map((r, i) => (
              <Table.Tr key={i}>
                <Table.Td>{r.diagnosis}</Table.Td>
                <Table.Td>{r.doctor_fio}</Table.Td>
                <Table.Td>{r.patient_count}</Table.Td>
                <Table.Td>{r.min_days}</Table.Td>
                <Table.Td>{r.max_days}</Table.Td>
              </Table.Tr>
            )) : (
              <Table.Tr>
                <Table.Td colSpan={5}><Center>Нет данных</Center></Table.Td>
              </Table.Tr>
            )}
          </Table.Tbody>
        </Table>
      )}
    </>
  );
}

export default HospitalizationReportPage;
