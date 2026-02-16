import { useEffect, useState } from 'react';
import { Table, Button, Modal, Select, Group, Title, TextInput, Loader, Center, Text } from '@mantine/core';
import { useForm } from '@mantine/form';
import { useDisclosure } from '@mantine/hooks';
import { apiClient } from '../api/client';
import type { Hospitalization, HospitalizationResponse, Patient, DoctorResponse, Diagnosis, Department } from '../types/models';

function HospitalizationsPage() {
  const [hospitalizations, setHospitalizations] = useState<HospitalizationResponse[]>([]);
  const [patients, setPatients] = useState<{ value: string; label: string }[]>([]);
  const [doctors, setDoctors] = useState<{ value: string; label: string }[]>([]);
  const [diagnoses, setDiagnoses] = useState<{ value: string; label: string }[]>([]);
  const [departments, setDepartments] = useState<{ value: string; label: string }[]>([]);
  const [opened, { open, close }] = useDisclosure(false);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [editingHosp, setEditingHosp] = useState<HospitalizationResponse | null>(null);


  const form = useForm<Omit<Hospitalization, 'id'>>({
    initialValues: {
      patient_id: 0,
      doctor_id: 0,
      diagnosis_id: 0,
      departament_id: 0,
      start_date: '',
      finish_date: ''
    }
  });

  const fetchData = async () => {
    try {
      setLoading(true);
      const [hospRes, patRes, docRes, diagRes, depRes] = await Promise.all([
        apiClient.get<HospitalizationResponse[]>('/hospitalizations'),
        apiClient.get<Patient[]>('/patients'),
        apiClient.get<DoctorResponse[]>('/doctors'),
        apiClient.get<Diagnosis[]>('/diagnoses'),
        apiClient.get<Department[]>('/departments'),
      ]);

      setHospitalizations(hospRes.data ?? []);
      setPatients((patRes.data ?? []).map(p => ({
        value: p.id!.toString(),
        label: `${p.last_name} ${p.first_name}`
      })));
      setDoctors((docRes.data ?? []).map(d => ({
        value: d.id!.toString(),
        label: `${d.last_name} ${d.first_name}`
      })));
      setDiagnoses((diagRes.data ?? []).map(d => ({
        value: d.id!.toString(),
        label: d.title
      })));
      setDepartments((depRes.data ?? []).map(d => ({
        value: d.id!.toString(),
        label: d.title
      })));
      setError(null);
    } catch (err) {
      setError("Ошибка загрузки данных");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  const handleSubmit = async (values: typeof form.values) => {
    const payload = {
      ...values,
      patient_id: Number(values.patient_id),
      doctor_id: Number(values.doctor_id),
      diagnosis_id: Number(values.diagnosis_id),
      departament_id: Number(values.departament_id),
    };
    try {
      if (editingHosp) {
        await apiClient.put(`/hospitalizations/${editingHosp.id}`, payload);
      } else {
        await apiClient.post('/hospitalizations', payload);
      }
      await fetchData();
      form.reset();
      closeModal();
    } catch (err) {
      setError(`Ошибка при ${editingHosp ? 'редактировании' : 'создании'} записи`);
    }
  };

  const handleDelete = async (id: number) => {
    try {
      await apiClient.delete(`/hospitalizations/${id}`);
      fetchData();
    } catch (err) {
      alert("Нельзя удалить госпитализацию, если она связана с другими данными.");
    }
  };

  const openModal = (h?: HospitalizationResponse) => {
    if (h) {
      setEditingHosp(h);
      form.setValues({
        patient_id: h.patient.id,
        doctor_id: h.doctor.id,
        diagnosis_id: h.diagnosis.id,
        departament_id: h.departament.id,
        start_date: h.start_date,
        finish_date: h.finish_date,
      });
    } else {
      setEditingHosp(null);
      form.reset();
    }
    open();
  };

  const closeModal = () => {
    setEditingHosp(null);
    form.reset();
    close();
  };

  return (
    <>
      <Modal opened={opened} onClose={closeModal} title={editingHosp ? 'Редактировать госпитализацию' : 'Оформить госпитализацию'}>
        <form onSubmit={form.onSubmit(handleSubmit)}>
          <Select label="Пациент" data={patients} {...form.getInputProps('patient_id')} required searchable />
          <Select label="Врач" data={doctors} {...form.getInputProps('doctor_id')} required searchable />
          <Select label="Диагноз" data={diagnoses} {...form.getInputProps('diagnosis_id')} required />
          <Select label="Отделение" data={departments} {...form.getInputProps('departament_id')} required />
          <TextInput label="Дата начала (YYYY-MM-DD)" {...form.getInputProps('start_date')} required />
          <TextInput label="Дата окончания (YYYY-MM-DD)" {...form.getInputProps('finish_date')} required />
          <Button type="submit" mt="md">{editingHosp ? 'Сохранить' : 'Создать'}</Button>
        </form>
      </Modal>

      <Group justify="space-between" mb="md">
        <Title order={2}>Госпитализации</Title>
        <Button onClick={() => openModal()}>Оформить госпитализацию</Button>
      </Group>

      {loading ? (
        <Center><Loader /></Center>
      ) : error ? (
        <Center><Text c="red">{error}</Text></Center>
      ) : (
        <Table highlightOnHover>
          <Table.Thead>
            <Table.Tr>
              <Table.Th>Пациент</Table.Th>
              <Table.Th>Врач</Table.Th>
              <Table.Th>Диагноз</Table.Th>
              <Table.Th>Период</Table.Th>
              <Table.Th>Действия</Table.Th>
            </Table.Tr>
          </Table.Thead>
          <Table.Tbody>
            {hospitalizations.length > 0 ? (
              hospitalizations.map((h) => (
                <Table.Tr key={h.id}>
                  <Table.Td>{h.patient ? `${h.patient.last_name} ${h.patient.first_name}` : "—"}</Table.Td>
                  <Table.Td>{h.doctor ? `${h.doctor.last_name} ${h.doctor.first_name}` : "—"}</Table.Td>
                  <Table.Td>{h.diagnosis?.title ?? "—"}</Table.Td>
                  <Table.Td>{`${h.start_date} - ${h.finish_date}`}</Table.Td>
                  <Table.Td>
                    <Button.Group>
                      <Button variant="light" size="xs" onClick={() => openModal(h)}>Ред.</Button>
                      <Button variant="light" color="red" size="xs" onClick={() => handleDelete(h.id!)}>Уд.</Button>
                    </Button.Group>
                  </Table.Td>
                </Table.Tr>
              ))
            ) : (
              <Table.Tr>
                <Table.Td colSpan={5} style={{ textAlign: 'center' }}>
                  Нет данных
                </Table.Td>
              </Table.Tr>
            )}
          </Table.Tbody>
        </Table>
      )}
    </>
  );
}

export default HospitalizationsPage;
