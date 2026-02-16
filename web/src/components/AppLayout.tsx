// web/src/components/AppLayout.tsx
import { AppShell, Burger, Group, NavLink } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import { Link, useLocation } from 'react-router-dom';

const navLinks = [
    { to: '/', label: 'Госпитализации' },
    { to: '/patients', label: 'Пациенты' },
    { to: '/doctors', label: 'Врачи' },
    { to: '/departments', label: 'Отделения' },
    { to: '/diagnoses', label: 'Диагнозы' },
    { to: '/reports/hospitalizations', label: "Отчет"},
]

export function AppLayout({ children }: { children: React.ReactNode }) {
  const [opened, { toggle }] = useDisclosure();
  const location = useLocation();

  return (
    <AppShell
      header={{ height: 60 }}
      navbar={{ width: 300, breakpoint: 'sm', collapsed: { mobile: !opened } }}
      padding="md"
    >
      <AppShell.Header>
        <Group h="100%" px="md">
          <Burger opened={opened} onClick={toggle} hiddenFrom="sm" size="sm" />
          ИС "Больница"
        </Group>
      </AppShell.Header>

      <AppShell.Navbar p="md">
        {navLinks.map((link) => (
            <NavLink
                key={link.to}
                component={Link}
                to={link.to}
                label={link.label}
                active={location.pathname === link.to}
            />
        ))}
      </AppShell.Navbar>

      <AppShell.Main>{children}</AppShell.Main>
    </AppShell>
  );
}