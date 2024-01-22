"use client"
import { reset, version } from "@/provider/smocker"
import { Button, JsonInput, Tooltip } from "@mantine/core"
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
import { notifications } from '@mantine/notifications';

export const VersionTooltip = () => {
    const req = useQuery({
        queryKey: ['mock', 'health'],
        queryFn: () => version()
    })
    return <JsonInput value={JSON.stringify(req.data || {}, undefined, ' ')} disabled autosize/>
}
export const ResetButton = () => {
    const client = useQueryClient()
    const req = useMutation({
        mutationFn: () => reset(),
        onSuccess: () => {
            client.invalidateQueries({queryKey: ['mock']})
            notifications.show({ title: 'Smocker Reset', message: 'Your instance of smocker has been successfully reset' })
        }
    })
    return <Tooltip label='Clear the mocks and the history of calls.'>
        <Button color='yellow' fullWidth onClick={() => req.mutate()}>Reset</Button>
    </Tooltip>
}