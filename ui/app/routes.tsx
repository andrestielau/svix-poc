import Link from 'next/link';
import React, { ReactNode } from 'react';
import { IconActivity, IconFingerprint, IconGauge, IconHistory, IconTestPipe } from '@tabler/icons-react';
import { NavLink } from '@mantine/core';

export type Route = {
    href: string
    label: string
    description?: string
    leftSection?: ReactNode
    rightSection?: ReactNode
    children?: Route[]
  }
export const routes: Route[] = [
    { 
      href: '/',
      label: 'Dashboard', 
      rightSection: <IconGauge size="1.5rem" stroke={1.5} />,
    },
    {
      href: '/webhook',
      label: 'WebHooks',
      description: 'Manage WebHooks', 
      rightSection: <IconFingerprint size="1.5rem" stroke={1.5} />,
      children: [
        { 
          href: '/',
          label: 'Dash',
          rightSection: <IconHistory size="1.5rem" stroke={1.5} />
        },
        { 
          href: '/applications',
          label: 'Applications',
          rightSection: <IconHistory size="1.5rem" stroke={1.5} />
        },
        { 
          href: '/event-types',
          label: 'EventTypes',
          rightSection: <IconHistory size="1.5rem" stroke={1.5} />
        }
      ]
    },
    { 
      href: '/mock',
      label: 'Mocks',
      description: 'Manage Mock Server', 
      rightSection: <IconTestPipe size="1.5rem" stroke={1.5} />,
      children: [
        { 
          href: '/',
          label: 'Dash',
          rightSection: <IconHistory size="1.5rem" stroke={1.5} />
        },
        { 
          href: '/history',
          label: 'History',
          rightSection: <IconHistory size="1.5rem" stroke={1.5} />
        },
        { 
          href: '/sessions',
          label: 'Sessions',
          rightSection: <IconHistory size="1.5rem" stroke={1.5} />
        }
      ]
    },
    { 
      href: '/pubsub',
      label: 'PubSub',
      description: 'Manage Messages', 
      rightSection: <IconActivity size="1.5rem" stroke={1.5} />,
      children: [
        { 
          href: '/',
          label: 'Dash',
          rightSection: <IconHistory size="1.5rem" stroke={1.5} />
        },
        { 
          href: '/router',
          label: 'Router',
          rightSection: <IconHistory size="1.5rem" stroke={1.5} />
        }, 
        { 
          href: '/schema',
          label: 'Schema',
          rightSection: <IconHistory size="1.5rem" stroke={1.5} />
        },
        { 
          href: '/topic',
          label: 'Topics',
          rightSection: <IconHistory size="1.5rem" stroke={1.5} />
        }]
    },
];

export type WalkNavTreeProps = {pathname: string, items?: Route[]}
export const WalkNavTree = ({ pathname, items }: WalkNavTreeProps) => items?.map((item) => <NavLink
  component={Link}
  href={item.href}
  key={item.label}
  active={item.href === pathname}
  label={item.label}
  description={item.description}
  leftSection={item.leftSection}
  rightSection={item.rightSection}
>
  {item.children && <WalkNavTree pathname={pathname} items={item.children?.
    map(({ href, ...rest }) => ({ href: item.href + href, ...rest }))}/>}
</NavLink>)