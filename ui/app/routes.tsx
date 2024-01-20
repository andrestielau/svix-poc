import Link from 'next/link';
import React, { ReactNode } from 'react';
import { IconActivity, IconApiApp, IconCalendarEvent, IconDashboard, IconGauge, IconHistory, IconHistoryToggle, IconPresentationAnalytics, IconRouter, IconSchema, IconTestPipe, IconWebhook } from '@tabler/icons-react';
import { NavLink, Tooltip } from '@mantine/core';

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
      label: 'Home', 
      rightSection: <IconGauge size="1.5rem" stroke={1.5} />,
    },
    {
      href: '/webhook',
      label: 'WebHooks',
      description: 'Manage WebHooks', 
      rightSection: <IconWebhook size="1.5rem" stroke={1.5} />,
      children: [
        { 
          href: '/',
          label: 'Dash',
          rightSection: <IconDashboard size="1.5rem" stroke={1.5} />
        },
        { 
          href: '/applications',
          label: 'Applications',
          rightSection: <IconApiApp size="1.5rem" stroke={1.5} />
        },
        { 
          href: '/event-types',
          label: 'EventTypes',
          rightSection: <IconCalendarEvent size="1.5rem" stroke={1.5} />
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
          rightSection: <IconDashboard size="1.5rem" stroke={1.5} />
        },
        { 
          href: '/history',
          label: 'History',
          rightSection: <IconHistory size="1.5rem" stroke={1.5} />
        },
        { 
          href: '/sessions',
          label: 'Sessions',
          rightSection: <IconHistoryToggle size="1.5rem" stroke={1.5} />
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
          rightSection: <IconDashboard size="1.5rem" stroke={1.5} />
        },
        { 
          href: '/router',
          label: 'Router',
          rightSection: <IconRouter size="1.5rem" stroke={1.5} />
        }, 
        { 
          href: '/schema',
          label: 'Schema',
          rightSection: <IconSchema size="1.5rem" stroke={1.5} />
        },
        { 
          href: '/topic',
          label: 'Topics',
          rightSection: <IconPresentationAnalytics size="1.5rem" stroke={1.5} />
        }]
    },
];

export type WalkNavTreeProps = {pathname: string, open: boolean, items?: Route[], to?: (href: string) => void}
export const WalkNavTree = ({ pathname, open, items, to }: WalkNavTreeProps) => items?.map((item) => {
const children = item.children && <WalkNavTree pathname={pathname} open={open} items={item.children?.
  map(({ href, ...rest }) => ({ href: item.href + href, ...rest }))} to={to}/>
const ret = (<NavLink
  component={Link}
  href={item.href}
  key={item.href}
  active={item.href === pathname}
  label={item.label}
  description={item.description}
  leftSection={item.leftSection}
  rightSection={item.rightSection}
  onDoubleClick={() => item.children && to && to(item.href)}
>{children}</NavLink>)
  return open ? ret : <Tooltip label={item.label}>{ret}</Tooltip>
})