"use client"
import { subscription } from "@/provider/pubsub"
import { Grid, Title } from "@mantine/core"
import { useQuery } from "@tanstack/react-query"

const SubscriptionPage = ({ params: { id } }: { params: { id: string } }) => {
    const {data} = useQuery({queryKey: ['pubsub', 'subscriptions', id], queryFn: async () => await subscription(id) })
    return <Grid>
        <Grid.Col span={12}>
            <Title order={1}>Subscription: {data?.name}</Title>
            <Title order={2}>Topic: {data?.topic}</Title>
        </Grid.Col>
    </Grid>
}

export default SubscriptionPage