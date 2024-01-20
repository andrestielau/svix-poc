import { listEndpoints, createEndpoint } from "@/provider/svix"
import { CreationModal, QueryList } from "@/components/crud"
import { Button, TextInput, Tooltip } from "@mantine/core"
import {  WithMenu } from "@/components/components"
import { EndpointIn, EndpointOut } from "svix"
import { useRouter } from "next/navigation"

export type EndpointListProps = {
    appId: string
    search: string
    setSearch: (s: string) => void 
}
export const EndpointList = ({ appId, search, setSearch }: EndpointListProps) => { 
    const router = useRouter()
    const queryKey = ["webhook", appId, "endpoints"]
    return <QueryList<EndpointOut> value={search} setValue={setSearch} 
        queryKey={queryKey} queryFn={async () => await listEndpoints(appId) }
        action={<CreationModal<EndpointIn> title="New Endpoint" queryKey={queryKey} 
            name='new-webhook-endpoint' initialValues={{ url: '' }} validate={{
                url: (value: string) => value ? null : 'Invalid url',
            }}
            mutationFn={async (i) => await createEndpoint(appId, i)}>{(form) => <>
                <TextInput label='Url' withAsterisk {...form.getInputProps('name')} />
                <TextInput label='Description' {...form.getInputProps('description')} />        
            </>}</CreationModal>}>
        {({ id, createdAt }) => <WithMenu key={id}>
            <Tooltip label={'Created At: '+ createdAt}>
                <Button onClick={() => 
                    router.push('/webhook/'+appId+'/endpoints/'+id)} variant="default" fullWidth>{id}</Button>
            </Tooltip>
        </WithMenu>}
    </QueryList> 
}