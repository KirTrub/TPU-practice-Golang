import { useEffect, useState } from 'react';
import { Table, Button, Modal, TextInput, Group, Title, Loader, Center, Text } from '@mantine/core';
import { useForm } from '@mantine/form';
import { useDisclosure } from '@mantine/hooks';
import { apiClient } from '../api/client';
import type { Diagnosis } from '../types/models';

function DiagnosesPage() {
  const [diagnoses, setDiagnoses] = useState<Diagnosis[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const [opened, { open, close }] = useDisclosure(false);
  const [editingDiagnosis, setEditingDiagnosis] = useState<Diagnosis | null>(null);
  
  const form = useForm({ initialValues: { title: '' } });

  const fetchDiagnoses = async () => {
    try {
      setLoading(true);
      const response = await apiClient.get<Diagnosis[]>('/diagnoses');
      setDiagnoses(response.data ?? []);
      setError(null);
    } catch (err) {
      setError('Не удалось загрузить диагнозы');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchDiagnoses();
  }, []);

  const handleSubmit = async (values: typeof form.values) => {
    try {
      if (editingDiagnosis) {
        await apiClient.put(`/diagnoses/${editingDiagnosis.id}`, values);
      } else {
        await apiClient.post('/diagnoses', values);
      }
      await fetchDiagnoses();
      form.reset();
      close();
    } catch (err) {
      setError(`Ошибка при ${editingDiagnosis ? 'редактировании' : 'создании'} диагноза`);
    }
  };

  const openModal = (diagnosis?: Diagnosis) => {
    if (diagnosis) {
      setEditingDiagnosis(diagnosis);
      form.setValues(diagnosis);
    } else {
      setEditingDiagnosis(null);
      form.reset();
    }
    open();
  };

  const closeModal = () => {
    setEditingDiagnosis(null);
    form.reset();
    close();
  };
 
  const handleDelete = async (id: number) => {
    try {
      await apiClient.delete(`/diagnoses/${id}`);
      fetchDiagnoses();
    } catch (error) {
      alert("Нельзя удалить диагноз, который используется.");
    }
  }

  return (
    <>
      <Modal opened={opened} onClose={closeModal} title={editingDiagnosis ? 'Редактировать диагноз' : 'Добавить диагноз'}>
        <form onSubmit={form.onSubmit(handleSubmit)}>
          <TextInput label="Название диагноза" {...form.getInputProps('title')} required />
          <Button type="submit" mt="md">{editingDiagnosis ? 'Сохранить' : 'Создать'}</Button>
        </form>
      </Modal>

      <Group justify="space-between" mb="md">
        <Title order={2}>Диагнозы</Title>
        <Button onClick={() => openModal()}>Добавить диагноз</Button>
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
              <Table.Th>Название</Table.Th>
              <Table.Th>Действия</Table.Th>
            </Table.Tr>
          </Table.Thead>
          <Table.Tbody>
            {diagnoses.length > 0 ? (
              diagnoses.map((d) => (
                <Table.Tr key={d.id}>
                  <Table.Td>{d.id}</Table.Td>
                  <Table.Td>{d.title}</Table.Td>
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
                <Table.Td colSpan={3}><Center>Нет данных</Center></Table.Td>
              </Table.Tr>
            )}
          </Table.Tbody>
        </Table>
      )}
    </>
  );
}

export default DiagnosesPage;