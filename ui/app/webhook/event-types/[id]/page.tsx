"use client"
import { getEventType } from "@/provider/svix"
import { Grid, Text, Title } from "@mantine/core"
import { useQuery } from "@tanstack/react-query"

const EventTypePage = ({ params: {id} }: { params: { id: string } }) => {
    const {data} = useQuery({queryKey: ['webhook-event-types'], queryFn: async () => await getEventType(id) })
    return <Grid>
        <Grid.Col span={12}>
            <Title order={1}>Event Type: {data?.name}</Title>
            <Title order={2}>Created At: {data?.createdAt?.toString()}</Title>
            {data?.updatedAt && (data?.updatedAt != data?.createdAt) && 
                <Title order={2}>Updated At: {data?.updatedAt?.toString()}</Title>}
            {data?.description && <Title order={2}>Description: {data?.description}</Title>}
            {data?.schemas?.map((schema: any) => <Text>{JSON.stringify(schema)}</Text>)}
        </Grid.Col>
    </Grid>
}
  

export default EventTypePage