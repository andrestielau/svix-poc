"use client"
import { Accordion, AppShell, Button, Group, Text, Tooltip } from "@mantine/core"
import { SearchBox } from "../components"
import { ReactNode, useState } from "react"
import { ResetButton, VersionTooltip } from "@/modules/mock/utils"
import { PublishModal, StatusTooltip } from "@/modules/pubsub/utils"
import { HealthTooltip } from "@/modules/webhook/utils"

export const Aside = ({}) => {
    const [search, setSearch] = useState()
    return  <AppShell.Aside>
        <AppShell.Section>
        <SearchBox value={search} setValue={setSearch}/>
        </AppShell.Section>
        <Accordion variant="contained" defaultValue="Apples">
            {services.map((item, i) => <Accordion.Item key={item.title} value={item.title}>
                <Tooltip label={item.tooltip}>
                    <Accordion.Control>
                        <Group wrap="nowrap">
                            <Text>{item.title}</Text>
                            <Text size="sm" c="dimmed" fw={400}>{item.description}</Text>
                        </Group>
                    </Accordion.Control>
                </Tooltip>
                <Accordion.Panel>{item.panel}</Accordion.Panel>
            </Accordion.Item>)}
        </Accordion>
    </AppShell.Aside>
}

type AccordionLabelProps = {
    tooltip: ReactNode
    description: string
    title: string
    panel: ReactNode
}

const services: AccordionLabelProps[] = [
    {
        title: 'PubSub',
        tooltip: <StatusTooltip/>,
        description: 'PubSub Actions',
        panel: <>
            <PublishModal />
        </>,
    },
    {
        title: 'Router',
        tooltip: '',
        description: 'Router Actions',
        panel:<>
            <Button>Do Something</Button>
        </>,
    },
    {
        title: 'WebHooks',
        tooltip: '',
        description: 'WebHooks Actions',
        panel: <>
            <Button>Do Something</Button>
        </>,
    },
    {
        title: 'Svix',
        tooltip: <HealthTooltip/>,
        description: 'Svix Actions',
        panel:<>
            <Button>Do Something</Button>
        </>,
    },
    {
        title: 'Smocker',
        description: 'Smocker Actions',
        tooltip: <VersionTooltip/>,
        panel: <>
            <ResetButton />
        </>,
    },
];