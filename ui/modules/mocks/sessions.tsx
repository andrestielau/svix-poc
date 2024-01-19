import { useMutation, useQuery, useQueryClient} from "@tanstack/react-query"
import { Button, Divider, Group, NumberInput, TextInput, Tooltip } from "@mantine/core"
import { NewModal, SearchList, WithMenu } from "@/components/components"
import { useForm } from "@mantine/form"
import { NewSession } from "@/provider/smocker/client"
import { create, sessions, startSession } from "@/provider/smocker"

export type SessionListProps = {
    search: string
    setSearch: (s: string) => void 
}
export const SessionList = ({ search, setSearch }: SessionListProps) => { 
    const client = useQueryClient()
    const data = useQuery({queryKey: ['webhook-sessions'], queryFn: async () => await sessions() })
    const add = useMutation({ mutationFn: async (i: NewSession) => await startSession(i), 
        onSuccess: () => client.invalidateQueries({ queryKey: ['webhook-sessions'] }) })
    const form = useForm({
        name: 'new-webhook-session',
        initialValues: {
            name: '',
        } as NewSession,
        validate: {},
    });
    return <SearchList value={search} setValue={setSearch}
        action={<NewModal title="New Session">
            <form onSubmit={form.onSubmit((values) => add.mutate(values))} onReset={() => form.reset()} >
                <TextInput label='Name' withAsterisk {...form.getInputProps('name')} />
                <Divider my="xs" />
                <Group justify="end">
                <Button variant="light" type='reset' disabled={!form.isDirty}>Reset</Button>
                <Button variant="outline" type='submit' disabled={!form.isValid()}>Create</Button>
                </Group>
            </form>
        </NewModal>}>
        {data.data?.map(({ id, name, date }) => <WithMenu key={id}>
            <Tooltip label={'Created At: '+ date}>
                <Button variant="default" fullWidth>{name || id}</Button>
            </Tooltip>
        </WithMenu>)}
    </SearchList> 
}