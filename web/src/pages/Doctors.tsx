import { useEffect, useState } from 'react';
import { Table, Button, Modal, TextInput, Select, Group, Title, Loader, Center, Text } from '@mantine/core';
import { useForm } from '@mantine/form';
import { useDisclosure } from '@mantine/hooks';
import { apiClient } from '../api/client';
import type { Doctor, DoctorResponse, Department } from '../types/models';

function DoctorsPage() {
  const [doctors, setDoctors] = useState<DoctorResponse[]>([]);
  const [departments, setDepartments] = useState<{ value: string; label: string }[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [editingDoctor, setEditingDoctor] = useState<Doctor | null>(null);
  
  const [opened, { open, close }] = useDisclosure(false);

  const form = useForm<Omit<Doctor, 'id'>>({
    initialValues: {
      first_name: '',
      last_name: '',
      sur_name: '',
      departament_id: 0,
    },
  });

  const fetchDoctors = async () => {
    try {
      setLoading(true);
      const response = await apiClient.get<DoctorResponse[]>('/doctors');
      setDoctors(response.data ?? []);
      setError(null);
    } catch (err) {
      setError('Не удалось загрузить врачей');
    } finally {
      setLoading(false);
    }
  };

  const fetchDepartments = async () => {
    try {
      const response = await apiClient.get<Department[]>('/departments');
      setDepartments(
        (response.data ?? []).map((d) => ({
          value: d.id!.toString(),
          label: d.title,
        }))
      );
    } catch (err) {
      console.error('Ошибка загрузки отделений:', err);
      setDepartments([]);
    }
  };

  useEffect(() => {
    fetchDoctors();
    fetchDepartments();
  }, []);

  const handleSubmit = async (values: typeof form.values) => {
    try {
      if (editingDoctor) {
        await apiClient.put(`/doctors/${editingDoctor.id}`, {
          ...values,
          departament_id: Number(values.departament_id),
        });
      } else {
        await apiClient.post('/doctors', {
          ...values,
          departament_id: Number(values.departament_id),
        });
      }
      await fetchDoctors();
      form.reset();
      close();
    } catch (err) {
      setError(`Ошибка при ${editingDoctor ? 'редактировании' : 'создании'} врача`);
    }
  };

  const openModal = (doctor?: Doctor) => {
    if (doctor) {
      setEditingDoctor(doctor);
      form.setValues({
        first_name: doctor.first_name,
        last_name: doctor.last_name,
        sur_name: doctor.sur_name,
        departament_id: doctor.departament_id,
      });
    } else {
      setEditingDoctor(null);
      form.reset();
    }
    open();
  };

  const closeModal = () => {
    setEditingDoctor(null);
    form.reset();
    close();
  };

  const handleDelete = async (id: number) => {
    try {
      await apiClient.delete(`/doctors/${id}`);
      fetchDoctors();
    } catch (error) {
      alert("Нельзя удалить врача, у которого есть приемы.");
    }
  };

  return (
    <>
      <Modal opened={opened} onClose={closeModal} title={editingDoctor ? 'Редактировать врача' : 'Добавить врача'}>
        <form onSubmit={form.onSubmit(handleSubmit)}>
          <TextInput label="Имя" {...form.getInputProps('first_name')} required />
          <TextInput label="Фамилия" {...form.getInputProps('last_name')} required />
          <TextInput label="Отчество" {...form.getInputProps('sur_name')} />
          <Select
            label="Отделение"
            data={departments}
            {...form.getInputProps('departament_id')}
            required
          />
          <Button type="submit" mt="md">{editingDoctor ? 'Сохранить' : 'Создать'}</Button>
        </form>
      </Modal>

      <Group justify="space-between" mb="md">
        <Title order={2}>Врачи</Title>
        <Button onClick={() => openModal()}>Добавить врача</Button>
      </Group>
      
      {loading ? (
        <Center><Loader /></Center>
      ) : error ? (
        <Center><Text c="red">{error}</Text></Center>
      ) : (
        <Table highlightOnHover>
          <Table.Thead>
            <Table.Tr>
              <Table.Th>ID</Table.Th>
              <Table.Th>ФИО</Table.Th>
              <Table.Th>Отделение</Table.Th>
              <Table.Th>Действия</Table.Th>
            </Table.Tr>
          </Table.Thead>
          <Table.Tbody>
            {doctors.length > 0 ? (
              doctors.map((d) => (
                <Table.Tr key={d.id}>
                  <Table.Td>{d.id}</Table.Td>
                  <Table.Td>{`${d.last_name} ${d.first_name} ${d.sur_name || ''}`}</Table.Td>
                  <Table.Td>{d.departament_title}</Table.Td>
                  <Table.Td>
                    <Button.Group>
                      <Button variant="light" size="xs" onClick={() => openModal(d)}>Ред.</Button>
                      <Button variant="light" color="red" size="xs" onClick={() => handleDelete(d.id!)}>Уд.</Button>
                    </Button.Group>
                  </Table.Td>
                </Table.Tr>
              ))
            ) : (
              <Table.Tr>
                <Table.Td colSpan={4} style={{ textAlign: 'center' }}>
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

export default DoctorsPage;