import { useEffect, useState } from 'react';
import { Table, Button, Modal, TextInput, Group, Title, Loader, Center, Text } from '@mantine/core';
import { useForm } from '@mantine/form';
import { useDisclosure } from '@mantine/hooks';
import { apiClient } from '../api/client';
import type { Department } from '../types/models';

function DepartmentsPage() {
  const [departments, setDepartments] = useState<Department[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const [opened, { open, close }] = useDisclosure(false);
  const [editingDep, setEditingDep] = useState<Department | null>(null);
  
  const form = useForm({ initialValues: { title: '' } });

  const fetchDepartments = async () => {
    try {
      setLoading(true);
      const response = await apiClient.get<Department[]>('/departments');
      setDepartments(response.data ?? []);
      setError(null);
    } catch (err) {
      setError('Не удалось загрузить отделения');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchDepartments();
  }, []);

  const handleSubmit = async (values: typeof form.values) => {
    try {
      if (editingDep) {
        await apiClient.put(`/departments/${editingDep.id}`, values);
      } else {
        await apiClient.post('/departments', values);
      }
      await fetchDepartments();
      form.reset();
      close();
    } catch (err) {
      setError(`Ошибка при ${editingDep ? 'редактировании' : 'создании'} отделения`);
    }
  };

  const openModal = (dep?: Department) => {
    if (dep) {
      setEditingDep(dep);
      form.setValues(dep);
    } else {
      setEditingDep(null);
      form.reset();
    }
    open();
  };

  const closeModal = () => {
    setEditingDep(null);
    form.reset();
    close();
  };
 
  const handleDelete = async (id: number) => {
    try {
      await apiClient.delete(`/departments/${id}`);
      fetchDepartments();
    } catch (error) {
      alert("Нельзя удалить отделение, которое используется врачами.");
    }
  };

  return (
    <>
      <Modal opened={opened} onClose={closeModal} title={editingDep ? 'Редактировать отделение' : 'Добавить отделение'}>
        <form onSubmit={form.onSubmit(handleSubmit)}>
          <TextInput label="Название отделения" {...form.getInputProps('title')} required />
          <Button type="submit" mt="md">{editingDep ? 'Сохранить' : 'Создать'}</Button>
        </form>
      </Modal>

      <Group justify="space-between" mb="md">
        <Title order={2}>Отделения</Title>
        <Button onClick={() => openModal()}>Добавить отделение</Button>
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
            {departments.length > 0 ? (
              departments.map((d) => (
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

export default DepartmentsPage;