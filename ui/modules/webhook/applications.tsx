import { listApplications, createApplication, deleteApplication } from "@/provider/svix"
import { CreationModal, DeleteMenuItem, QueryList } from "@/components/crud"
import { Button, NumberInput, TextInput, Tooltip } from "@mantine/core"
import { ApplicationIn, ApplicationOut } from "svix"
import { WithMenu } from "@/components/components"
import { useRouter } from "next/navigation"

const queryKey = ["webhook", "applications"]
export type ApplicationListProps = {
    search: string
    setSearch: (s: string) => void 
}
export const ApplicationList = ({ search, setSearch }: ApplicationListProps) => { 
    const router = useRouter()
    return <QueryList<ApplicationOut> value={search} setValue={setSearch} accessor={({ id }) => id}
        queryKey={queryKey} queryFn={async () => await listApplications() }
        action={<CreationModal<ApplicationIn> title="New Application" queryKey={queryKey} 
            name='new-webhook-application' initialValues={{ name: '' }} validate={{
                name: (value: string) => value ? null : 'Invalid name',
            }}
            mutationFn={async (i) => await createApplication(i)}>{(form) => <>
                <TextInput label='Name' withAsterisk {...form.getInputProps('name')} />
                <TextInput label='Uid' {...form.getInputProps('uid')} />
                <TextInput label='Description' {...form.getInputProps('description')} />     
                <NumberInput label='Rate Limit' {...form.getInputProps('rateLimit')} />    
            </>}</CreationModal>}>
        {({ id, createdAt }) => <WithMenu id={id} key={id}
            danger={<DeleteMenuItem queryKey={queryKey} mutationFn={() => deleteApplication(id)}/>}>
            <Tooltip label={'Created At: '+ createdAt}>
                <Button onClick={() => 
                    router.push('/webhook/applications/'+id)} variant="default" fullWidth>{id}</Button>
            </Tooltip>
        </WithMenu>}
    </QueryList> 
}