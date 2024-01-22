import { health } from "@/provider/svix"
import { JsonInput } from "@mantine/core"
import { useQuery } from "@tanstack/react-query"

export const HealthTooltip = () => {
    const req = useQuery({
        queryKey: ['svix', 'health'],
        queryFn: () => health()
    })
    return <JsonInput value={JSON.stringify(req.data || {}, undefined, '\n')} disabled autosize/>
}