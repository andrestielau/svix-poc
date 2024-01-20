"use client"

import { EndpointList } from "@/modules/webhook/endpoints"
import { MessageList } from "@/modules/webhook/messages"
import { getApplication } from "@/provider/svix"
import { Grid, Title } from "@mantine/core"
import { useQuery } from "@tanstack/react-query"
import { useState } from "react"

const AppPage = ({ params: { appId} }: { params: { appId: string } }) => {
    const {data} = useQuery({queryKey: ['webhook-applications'], queryFn: async () => await getApplication(appId) })
    const [searchEndpoints, setSearchEndpoints] = useState('')
    const [searchMessages, setSearchMessages] = useState('')
    return <Grid>
        <Grid.Col span={12}>
            <Title order={1}>Application: {data?.name || data?.uid || data?.id}</Title>
            <Title order={2}>Created At: {data?.createdAt?.toString()}</Title>
            <Title order={2}>Id: {data?.id}</Title>
            {data?.uid && <Title order={2}>Uid: {data?.uid}</Title>}
            {!!Object.keys(data?.metadata||{}).length && <Title order={2}>Metadata: {JSON.stringify(data?.metadata)}</Title>}
            {data?.rateLimit && <Title order={2}>Rate Limit: {data?.rateLimit}</Title>}
        </Grid.Col>
        <Grid.Col span={6}>
            <Title order={1}>Endpoints</Title>
            <EndpointList search={searchEndpoints} setSearch={setSearchEndpoints} appId={appId}/>
        </Grid.Col>
        <Grid.Col span={6}>
            <Title order={1}>Messages</Title>
            <MessageList search={searchMessages} setSearch={setSearchMessages} appId={appId}/>
        </Grid.Col>
    </Grid>
}
  
export default AppPage