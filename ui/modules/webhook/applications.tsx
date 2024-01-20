import { listApplications, createApplication } from "@/provider/svix"
import { CreationModal, QueryList } from "@/components/crud"
import { Button, TextInput, Tooltip } from "@mantine/core"
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
    return <QueryList<ApplicationOut> value={search} setValue={setSearch} 
        queryKey={queryKey} queryFn={async () => await listApplications() }
        action={<CreationModal<ApplicationIn> title="New Application" queryKey={queryKey} 
            name='new-webhook-application' initialValues={{ name: '' }} validate={{
                name: (value: string) => value ? null : 'Invalid name',
            }}
            mutationFn={async (i) => await createApplication(i)}>{(form) => <>
                <TextInput label='Name' withAsterisk {...form.getInputProps('name')} />
                <TextInput label='Description' {...form.getInputProps('description')} />        
            </>}</CreationModal>}>
        {({ id, createdAt }) => <WithMenu key={id}>
            <Tooltip label={'Created At: '+ createdAt}>
                <Button onClick={() => 
                    router.push('/webhook/applications/'+id)} variant="default" fullWidth>{id}</Button>
            </Tooltip>
        </WithMenu>}
    </QueryList> 
}