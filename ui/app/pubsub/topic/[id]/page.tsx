"use client"
import { topic } from "@/provider/pubsub"
import { Grid, Title } from "@mantine/core"
import { useQuery } from "@tanstack/react-query"

const TopicPage = ({ params: { id } }: { params: { id: string } }) => {
    const {data} = useQuery({queryKey: ['pubsub', 'topics', id], queryFn: async () => await topic(id) })
    return <Grid>
        <Grid.Col span={12}>
            <Title order={1}>Topic: {data?.name}</Title>
            {data?.kmsKeyName && <Title order={2}>Kms Key Name: {data?.kmsKeyName}</Title>}
            {data?.schemaSettings?.schema && <Title order={2}>Schema: {data?.schemaSettings?.schema}</Title>}
            {!!Object.keys(data?.labels||{}).length && <Title order={2}>Labels: {JSON.stringify(data?.labels)}</Title>}
            {!!data?.messageRetentionDuration?.nanos && <Title order={2}>Retention Duration: {data?.messageRetentionDuration?.nanos}ns</Title>}
            {!!data?.messageStoragePolicy?.allowedPersistenceRegions && <Title order={2}>Allowed Persistence Regions: {data?.messageStoragePolicy?.allowedPersistenceRegions}</Title>}
        </Grid.Col>
    </Grid>
}

export default TopicPage