"use client"
import '@mantine/core/styles.css';
import React, { useState } from 'react';
import { MantineProvider, ColorSchemeScript, AppShell, Burger, ScrollArea, Flex, NavLink } from '@mantine/core';
import { theme } from '../theme';
import { useDisclosure } from '@mantine/hooks';
import { IconActivity, IconChevronRight, IconFingerprint, IconGauge, IconTestPipe } from '@tabler/icons-react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { useRouter } from 'next/router';
import { usePathname } from 'next/navigation';
import Link from 'next/link';

const queryClient = new QueryClient()
const data = [
  { 
    href: '/',
    label: 'Dashboard', 
    leftSection: <IconChevronRight size="1.5rem" stroke={1.5} />,
    rightSection: <IconGauge size="1.5rem" stroke={1.5} />,
  },
  {
    href: '/webhooks',
    label: 'WebHooks',
    description: 'Manage WebHooks', 
    rightSection: <IconFingerprint size="1.5rem" stroke={1.5} />,
  },
  { 
    href: '/router',
    label: 'Router',
    description: 'Manage Subscriptions', 
    rightSection: <IconActivity size="1.5rem" stroke={1.5} />,
  },
  { 
    href: '/mocks',
    label: 'Mocks',
    description: 'Manage mock server', 
    rightSection: <IconTestPipe size="1.5rem" stroke={1.5} />,
  },
];

export default function RootLayout({ children }: { children: any  }) {
  const [mobileOpened, { toggle: toggleMobile }] = useDisclosure();
  const [desktopOpened, { toggle: toggleDesktop }] = useDisclosure();
  const pathname = usePathname()
  const [active, setActive] = useState(data.findIndex(({ href }) => href === pathname ))
  
  return (
    <html lang="en">
      <head>
        <ColorSchemeScript />
        <link rel="shortcut icon" href="/favicon.svg" />
        <meta
          name="viewport"
          content="minimum-scale=1, initial-scale=1, width=device-width, user-scalable=no"
        />
      </head>
      <body>
        <QueryClientProvider client={queryClient}>
          <MantineProvider theme={theme}>
            <AppShell
              header={{ height: 60 }}
              navbar={{
                width: { base: 300 },
                breakpoint: 'sm',
                collapsed: { mobile: !mobileOpened, desktop: !desktopOpened },
              }}
              padding="md"
            >
              <AppShell.Header>
                <Burger
                  opened={mobileOpened}
                  onClick={toggleMobile}
                  hiddenFrom="sm"
                  size="sm"
                />
                <div>Logo</div>
              </AppShell.Header>
              <AppShell.Navbar p="xs" style={{width: 360}}>
                <AppShell.Section>Navbar header</AppShell.Section>
                <AppShell.Section grow component={ScrollArea}>
                  {data.map((item, index) => <NavLink
                    component={Link}
                    href={item.href}
                    key={item.label}
                    active={index === active}
                    label={item.label}
                    description={item.description}
                    leftSection={item.leftSection}
                    rightSection={item.rightSection}
                    onClick={() => setActive(index)}
                  />)}
                </AppShell.Section>
                <AppShell.Section>
                  <Flex
                    p="xs"
                    mih={50}
                    gap="md"
                    justify="flex-end"
                    align="center"
                    direction="row"
                    wrap="wrap"
                  >
                    <Burger
                      visibleFrom='sm'
                      opened={desktopOpened}
                      onClick={toggleDesktop}
                      size="sm"
                    />
                  </Flex>
                </AppShell.Section>
              </AppShell.Navbar>
              <AppShell.Main p={80}>
                {children}
              </AppShell.Main>
            </AppShell>
          </MantineProvider>
        </QueryClientProvider>
      </body>
    </html>
  );
}