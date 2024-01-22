import { status } from "@/provider/pubsub"
import { JsonInput } from "@mantine/core"
import { useQuery } from "@tanstack/react-query"

export const StatusTooltip = () => {
    const req = useQuery({
        queryKey: ['pubsub', 'status'],
        queryFn: () => status()
    })
    return <JsonInput value={JSON.stringify(req.data || {}, undefined, ' ')} disabled autosize/>
}