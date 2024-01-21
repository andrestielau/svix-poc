"use client"
import { MantineProvider, ColorSchemeScript, AppShell, Burger, ScrollArea, Flex, Breadcrumbs, Anchor, Divider, Image, Group } from '@mantine/core';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { usePathname, useRouter } from 'next/navigation';
import { useDisclosure, useColorScheme } from '@mantine/hooks';
import { WalkNavTree, routes } from './routes';
import NextImage from 'next/image';
import { theme } from '../theme';
import '@mantine/core/styles.css';
import React from 'react';
import Link from 'next/link';

const queryClient = new QueryClient()

export default function RootLayout({ children }: { children: any  }) {
  const [opened, { toggle }] = useDisclosure();
  const pathname = usePathname()
  const router = useRouter()
  return (
    <html lang="en">
      <head>
        <ColorSchemeScript />
        <link rel="shortcut icon" href="/favicon.png" />
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
                width: 300,
                breakpoint: 'sm',
                collapsed: { mobile: !opened, desktop: !opened },
              }}
              padding="xs"
            >
              <AppShell.Header>
                <Group justify="space-between">
                  <Flex
                    p="2.5"
                    mih={50}
                    gap="md"
                    justify="flex-start"
                    align="center"
                    direction="row"
                    wrap="wrap"
                  >
                    <Link href='/'>
                      <Image component={NextImage} src={useColorScheme() === 'light' ?'/logo.svg':'/logo-dark.svg'} height={50} width={50} alt="Logo" />
                    </Link>
                  </Flex>
                  <Flex
                    p="5"
                    pr='md'
                    mih={50}
                    gap="md"
                    justify="flex-end"
                    align="center"
                    direction="row"
                    wrap="wrap"
                  > 
                    <Burger opened={opened} onClick={toggle} size="sm" hiddenFrom='sm' />
                  </Flex>
                </Group>
              </AppShell.Header>
              <AppShell.Navbar p="xs" style={{minWidth: 360}}>
                <AppShell.Section>Navbar header</AppShell.Section>
                <AppShell.Section grow component={ScrollArea}>
                  <WalkNavTree pathname={pathname} items={routes} open={opened} to={router.push}/>
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
                    <Burger opened={opened} onClick={toggle} size="sm" visibleFrom='sm'/>
                  </Flex>
                </AppShell.Section>
              </AppShell.Navbar>
              <AppShell.Main p={80}>{pathname !== '/' && <>
                <Breadcrumbs>{pathname.split('/').map((v, i, a) => 
                  <Anchor key={i} href={a.slice(0,i).join("/")+"/"+v}>{v}</Anchor>)}
                </Breadcrumbs>
                <Divider mt='md'/></>}
                {children}
              </AppShell.Main>
            </AppShell>
          </MantineProvider>
        </QueryClientProvider>
      </body>
    </html>
  );
}
