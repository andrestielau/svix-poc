"use client"

import { useQuery } from "@tanstack/react-query"
import { getEndpoint } from "@/provider/svix"
import { Grid, TagsInput, Title } from "@mantine/core"
import { useRouter } from "next/navigation"
import { useState } from "react"

const EndpointPage = ({ params: { appId, id } }: { params: { appId: string, id: string } }) => {
    const {data} = useQuery({queryKey: ['webhook','applications', appId, 'endpoints', id], queryFn: async () => await getEndpoint(appId, id) })
    const [searchAttempts, setSearchAttempts] = useState('')
    const router = useRouter()
    return <Grid>
        <Grid.Col span={12}>
            <Title order={1}>Endpoint: {data?.uid || data?.url}</Title>
            <Title order={2}>Created At: {data?.createdAt?.toString()}</Title>
            <Title order={2}>Id: {data?.id}</Title>
            {data?.uid && <Title order={2}>Uid: {data?.uid}</Title>}
            {!!Object.keys(data?.metadata||{}).length && <Title order={2}>Metadata: {JSON.stringify(data?.metadata)}</Title>}
            {data?.rateLimit && <Title order={2}>Rate Limit: {data?.rateLimit}</Title>}
            {data?.filterTypes && <TagsInput label={<Title order={2}>Filter Types: </Title>} defaultValue={data?.filterTypes}  disabled />}
        </Grid.Col>
        <Grid.Col span={6}>
            <Title order={1} onDoubleClick={() => router.push('/webhook/applications/'+appId+'/endpoints/'+id+'/attempts')}>Attempts</Title>
        </Grid.Col>
    </Grid>
}
  
export default EndpointPage