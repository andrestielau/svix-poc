"use client"
import '@mantine/core/styles.css';
import React from 'react';
import { MantineProvider, ColorSchemeScript, AppShell, Burger, ScrollArea, Flex } from '@mantine/core';
import { theme } from '../theme';
import { useDisclosure } from '@mantine/hooks';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { usePathname } from 'next/navigation';
import { WalkNavTree, routes } from './routes';

const queryClient = new QueryClient()


export default function RootLayout({ children }: { children: any  }) {
  const [mobileOpened, { toggle: toggleMobile }] = useDisclosure();
  const [desktopOpened, { toggle: toggleDesktop }] = useDisclosure();
  const pathname = usePathname()
  
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
                  <WalkNavTree pathname={pathname} items={routes} />
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