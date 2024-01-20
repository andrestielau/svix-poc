"use client"
import { MantineProvider, ColorSchemeScript, AppShell, Burger, ScrollArea, Flex, em, Breadcrumbs, Anchor } from '@mantine/core';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { useDisclosure, useMediaQuery } from '@mantine/hooks';
import { usePathname, useRouter } from 'next/navigation';
import { WalkNavTree, routes } from './routes';
import { theme } from '../theme';
import '@mantine/core/styles.css';
import React from 'react';

const queryClient = new QueryClient()

export default function RootLayout({ children }: { children: any  }) {
  const [mobileOpened, { toggle: toggleMobile }] = useDisclosure();
  const [desktopOpened, { toggle: toggleDesktop }] = useDisclosure();
  const isMobile = useMediaQuery(`(max-width: ${em(750)})`);
  const pathname = usePathname()
  const router = useRouter()
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
                  <WalkNavTree pathname={pathname} items={routes} open={(isMobile && mobileOpened) || (!isMobile && desktopOpened)} to={router.push}/>
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
              <AppShell.Main p={80}>{pathname !== '/' &&
                <Breadcrumbs>{pathname.split('/').map((v, i, a) => 
                  <Anchor key={i} href={a.slice(0,i).join("/")+"/"+v}>{v}</Anchor>)}
                </Breadcrumbs>}
                {children}
              </AppShell.Main>
            </AppShell>
          </MantineProvider>
        </QueryClientProvider>
      </body>
    </html>
  );
}