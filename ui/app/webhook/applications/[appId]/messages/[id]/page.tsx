"use client"
import { getMessage } from "@/provider/svix"
import { JsonInput, Title } from "@mantine/core"
import { useQuery } from "@tanstack/react-query"

const MessagesPage = ({ params: { appId, id } }: { params: { appId: string, id: string } }) => {
    const {data} = useQuery({queryKey: ['webhook','applications', appId], queryFn: async () => await getMessage(appId, id) })
    return <>
        <Title order={1}>Message {id}</Title>
        <JsonInput label='Payload' value={data?.payload} disabled />
    </>
}
  
export default MessagesPage