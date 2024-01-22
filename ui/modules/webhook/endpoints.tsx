import { listEndpoints, createEndpoint, deleteEndpoint } from "@/provider/svix"
import { CreationModal, DeleteMenuItem, QueryList } from "@/components/crud"
import { Button, TextInput, Tooltip } from "@mantine/core"
import { WithMenu } from "@/components/components"
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
    return <QueryList<EndpointOut> value={search} setValue={setSearch} accessor={({ id }) => id}
        queryKey={queryKey} queryFn={async () => await listEndpoints(appId) }
        action={<CreationModal<EndpointIn> title="New Endpoint" queryKey={queryKey} 
            name='new-webhook-endpoint' initialValues={{ url: '' }} validate={{
                url: (value: string) => value ? null : 'Invalid url',
            }}
            mutationFn={async (i) => await createEndpoint(appId, i)}>{(form) => <>
                <TextInput label='Url' withAsterisk {...form.getInputProps('url')} />
                <TextInput label='Uid' {...form.getInputProps('uid')} />
                <TextInput label='Name' {...form.getInputProps('name')} />
                <TextInput label='Description' {...form.getInputProps('description')} />        
            </>}</CreationModal>}>
        {({ id, uid, url, createdAt }) => <WithMenu id={id} key={id}
            danger={<DeleteMenuItem queryKey={queryKey} mutationFn={() => deleteEndpoint(appId, id)}/>}>
            <Tooltip label={<label>Created At: {createdAt.toString()}<br/>{url}</label>}>
                <Button onClick={() => 
                    router.push('/webhook/applications/'+appId+'/endpoints/'+id)} variant="default" fullWidth>{uid || id}</Button>
            </Tooltip>
        </WithMenu>}
    </QueryList> 
}