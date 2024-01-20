"use client"
import { mock } from "@/provider/smocker"
import { Grid, Title } from "@mantine/core"
import { useQuery } from "@tanstack/react-query"

const MockPage = ({ params: { id } }: { params: { id: string } }) => {
    const {data} = useQuery({queryKey: ['webhook-mocks'], queryFn: async () => await mock(id) })
    return <Grid>
        <Grid.Col span={12}>
            <Title order={1}>Mock: {data?.state?.id}</Title>
            <Title order={2}>Created At: {data?.state?.creation_date}</Title>
            {data?.request && <Title order={2}>Request: {JSON.stringify(data?.request)}</Title>}
            {data?.dynamic_response && <Title order={2}>Response: {JSON.stringify(data?.dynamic_response)}</Title>}
        </Grid.Col>
    </Grid>
}

export default MockPage