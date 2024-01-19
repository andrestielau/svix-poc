import { addEventType, listEventTypes, remEventType } from "@/provider/drizzle/queries"
import { useMutation, useQuery, useQueryClient} from "@tanstack/react-query"
import { Button, Divider, Group, TextInput, Tooltip } from "@mantine/core"
import { NewModal, SearchList, WithMenu } from "@/components/components"
import { NewEventType } from "@/provider/drizzle/schema"
import { useForm } from "@mantine/form"

export type EventTypeListProps = {
    search: string
    setSearch: (s: string) => void 
}
export const EventTypeList = ({ search, setSearch }: EventTypeListProps) => { 
    const client = useQueryClient()
    const data = useQuery({queryKey: ['event-types'], queryFn: async () => await listEventTypes() })
    const add = useMutation({ mutationFn: async (i: NewEventType) => await addEventType(i), 
        onSuccess: () => client.invalidateQueries({ queryKey: ['event-types'] }) }) // TODO: replace with grpc
    const rem = useMutation({ mutationFn: async (i: string) => await remEventType(i), 
        onSuccess: () => client.invalidateQueries({ queryKey: ['event-types'] }) })
    const form = useForm({
        name: 'new-event-type',
        initialValues: {
          id: '',
        } as NewEventType,
        validate: {
          id: (value: string) => value ? null : 'Invalid id',
        },
    });
    return <SearchList value={search} setValue={setSearch}
        action={<NewModal title="New Event Type">
            <form onSubmit={form.onSubmit((values) => add.mutate(values))} onReset={() => form.reset()} >
                <TextInput  label='Id' withAsterisk {...form.getInputProps('id')} />
                <Divider my="xs" />
                <Group justify="end">
                <Button variant="light" type='reset' disabled={!form.isDirty}>Reset</Button>
                <Button variant="outline" type='submit' disabled={!form.isValid()}>Create</Button>
                </Group>
            </form>
        </NewModal>}>
        {data.data?.map(({ id, createdAt }) => <WithMenu key={id}>
            <Tooltip label={'Created At: '+ createdAt}>
                <Button variant="default" fullWidth>{id}</Button>
            </Tooltip>
        </WithMenu>)}
    </SearchList> 
}