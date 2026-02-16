// web/src/pages/Patients.tsx
import { useEffect, useState } from 'react';
import { Table, Button, Modal, TextInput, Group, Title } from '@mantine/core';
import { useForm } from '@mantine/form';
import { useDisclosure } from '@mantine/hooks';
import { apiClient } from '../api/client';
import type { Patient } from '../types/models';

function PatientsPage() {
  const [patients, setPatients] = useState<Patient[]>([]);
  const [opened, { open, close }] = useDisclosure(false);
  const [editingPatient, setEditingPatient] = useState<Patient | null>(null);

  const form = useForm<Omit<Patient, 'id'>>({
    initialValues: {
      first_name: '',
      last_name: '',
      sur_name: '',
      gender: '',
      birth_date: '',
      address: '',
    },
  });

  const fetchPatients = async () => {
    try {
      const response = await apiClient.get<Patient[]>('/patients');
      // если API вернул null или что-то не то → ставим пустой массив
      setPatients(Array.isArray(response.data) ? response.data : []);
    } catch (err) {
      console.error('Ошибка загрузки пациентов', err);
      setPatients([]); // чтобы точно не было null
    }
  };

  useEffect(() => {
    fetchPatients();
  }, []);

  const handleSubmit = async (values: typeof form.values) => {
    if (editingPatient) {
      await apiClient.put(`/patients/${editingPatient.id}`, values);
    } else {
      await apiClient.post('/patients', values);
    }
    await fetchPatients();
    closeModal();
  };

  const openModal = (patient?: Patient) => {
    if (patient) {
      setEditingPatient(patient);
      form.setValues(patient);
    } else {
      setEditingPatient(null);
      form.reset();
    }
    open();
  };

  const closeModal = () => {
    setEditingPatient(null);
    form.reset();
    close();
  };
  
  const handleDelete = async (id: number) => {
    await apiClient.delete(`/patients/${id}`);
    fetchPatients();
  }

  return (
    <>
      <Modal 
        opened={opened} 
        onClose={closeModal} 
        title={editingPatient ? 'Редактировать пациента' : 'Добавить пациента'}
      >
        <form onSubmit={form.onSubmit(handleSubmit)}>
          <TextInput label="Имя" {...form.getInputProps('first_name')} required />
          <TextInput label="Фамилия" {...form.getInputProps('last_name')} required />
          <TextInput label="Отчество" {...form.getInputProps('sur_name')} />
          <TextInput label="Пол" {...form.getInputProps('gender')} required />
          <TextInput label="Дата рождения (YYYY-MM-DD)" {...form.getInputProps('birth_date')} required />
          <TextInput label="Адрес" {...form.getInputProps('address')} required />
          <Button type="submit" mt="md">{editingPatient ? 'Сохранить' : 'Создать'}</Button>
        </form>
      </Modal>

      <Group justify="space-between" mb="md">
        <Title order={2}>Пациенты</Title>
        <Button onClick={() => openModal()}>Добавить пациента</Button>
      </Group>

      <Table>
        <Table.Thead>
          <Table.Tr>
            <Table.Th>ID</Table.Th>
            <Table.Th>ФИО</Table.Th>
            <Table.Th>Дата рождения</Table.Th>
            <Table.Th>Адрес</Table.Th>
            <Table.Th>Действия</Table.Th>
          </Table.Tr>
        </Table.Thead>
        <Table.Tbody>
          {patients.length > 0 ? (
            patients.map((p) => (
              <Table.Tr key={p.id}>
                <Table.Td>{p.id}</Table.Td>
                <Table.Td>{`${p.last_name} ${p.first_name} ${p.sur_name || ''}`}</Table.Td>
                <Table.Td>{p.birth_date}</Table.Td>
                <Table.Td>{p.address}</Table.Td>
                <Table.Td>
                  <Button.Group>
                    <Button variant="light" size="xs" onClick={() => openModal(p)}>Ред.</Button>
                    <Button variant="light" color="red" size="xs" onClick={() => handleDelete(p.id!)}>Уд.</Button>
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
    </>
  );
}

export default PatientsPage;
