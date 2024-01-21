"use client"
import { schema } from "@/provider/pubsub"
import { Code, Grid, Title } from "@mantine/core"
import { useQuery } from "@tanstack/react-query"

const SchemaPage = ({ params: { id } }: { params: { id: string } }) => {
    const {data} = useQuery({queryKey: ['pubsub', 'schema', id], queryFn: async () => await schema(id) })
    return <Grid>
        <Grid.Col span={12}>
            <Title order={1}>Schema: {data?.name}</Title>
            <Title order={2}>Created At: {data?.revisionCreateTime?.nanos}</Title>
            <Code block>{data?.definition}</Code>
        </Grid.Col>
    </Grid>
}

export default SchemaPage