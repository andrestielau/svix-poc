import { listEventTypes, createEventType, deleteEventType } from "@/provider/svix"
import { useMutation, useQuery, useQueryClient} from "@tanstack/react-query"
import { Button, Divider, Group, TextInput, Tooltip } from "@mantine/core"
import { NewModal, SearchList, WithMenu } from "@/components/components"
import { useForm } from "@mantine/form"
import { EventTypeIn } from "svix"

export type EventTypeListProps = {
    search: string
    setSearch: (s: string) => void 
}
export const EventTypeList = ({ search, setSearch }: EventTypeListProps) => { 
    const client = useQueryClient()
    const data = useQuery({queryKey: ['webhook-event-types'], queryFn: async () => await listEventTypes() })
    const add = useMutation({ mutationFn: async (i: EventTypeIn) => await createEventType(i), 
        onSuccess: () => client.invalidateQueries({ queryKey: ['webhook-event-types'] }) }) // TODO: replace with grpc
    const rem = useMutation({ mutationFn: async (i: string) => await deleteEventType(i), 
        onSuccess: () => client.invalidateQueries({ queryKey: ['webhook-event-types'] }) })
    const form = useForm({
        name: 'new-webhook-event-type',
        initialValues: {
          name: '',
          description: '',
        } as EventTypeIn,
        validate: {
            name: (value: string) => value ? null : 'Invalid name',
        },
    });
    return <SearchList value={search} setValue={setSearch}
        action={<NewModal title="New Event Type">
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
        {data.data?.data?.map(({ name, createdAt }) => <WithMenu key={name}>
            <Tooltip label={'Created At: '+ createdAt}>
                <Button variant="default" fullWidth>{name}</Button>
            </Tooltip>
        </WithMenu>)}
    </SearchList> 
}