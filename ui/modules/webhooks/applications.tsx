import { listApplications, createApplication, deleteApplication } from "@/provider/svix"
import { useMutation, useQuery, useQueryClient} from "@tanstack/react-query"
import { Button, Divider, Group, TextInput, Tooltip } from "@mantine/core"
import { NewModal, SearchList, WithMenu } from "@/components/components"
import { useForm } from "@mantine/form"
import { ApplicationIn } from "svix"

export type ApplicationListProps = {
    search: string
    setSearch: (s: string) => void 
}
export const ApplicationList = ({ search, setSearch }: ApplicationListProps) => { 
    const client = useQueryClient()
    const data = useQuery({queryKey: ['webhook-applications'], queryFn: async () => await listApplications() })
    const add = useMutation({ mutationFn: async (i: ApplicationIn) => await createApplication(i), 
        onSuccess: () => client.invalidateQueries({ queryKey: ['webhook-applications'] }) }) // TODO: replace with grpc
    const rem = useMutation({ mutationFn: async (i: string) => await deleteApplication(i), 
        onSuccess: () => client.invalidateQueries({ queryKey: ['webhook-applications'] }) })
    const form = useForm({
        name: 'new-webhook-application',
        initialValues: {
          name: '',
          description: '',
        } as ApplicationIn,
        validate: {
            name: (value: string) => value ? null : 'Invalid name',
        },
    });
    return <SearchList value={search} setValue={setSearch}
        action={<NewModal title="New Application">
            <form onSubmit={form.onSubmit((values) => add.mutate(values))} onReset={() => form.reset()} >
                <TextInput label='Name' withAsterisk {...form.getInputProps('name')} />
                <TextInput label='Description' {...form.getInputProps('description')} />
                <Divider my="xs" />
                <Group justify="end">
                <Button variant="light" type='reset' disabled={!form.isDirty}>Reset</Button>
                <Button variant="outline" type='submit' disabled={!form.isValid()}>Create</Button>
                </Group>
            </form>
        </NewModal>}>
        {data.data?.data?.map(({ id, createdAt }) => <WithMenu key={id}>
            <Tooltip label={'Created At: '+ createdAt}>
                <Button variant="default" fullWidth>{id}</Button>
            </Tooltip>
        </WithMenu>)}
    </SearchList> 
}