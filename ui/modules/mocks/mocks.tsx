import { useMutation, useQuery, useQueryClient} from "@tanstack/react-query"
import { Button, Divider, Group, NumberInput, TextInput, Tooltip } from "@mantine/core"
import { NewModal, SearchList, WithMenu } from "@/components/components"
import { useForm } from "@mantine/form"
import { NewMock } from "@/provider/smocker/client"
import { create, mocks } from "@/provider/smocker"

export type MockListProps = {
    search: string
    setSearch: (s: string) => void 
}
export const MockList = ({ search, setSearch }: MockListProps) => { 
    const client = useQueryClient()
    const data = useQuery({queryKey: ['webhook-mocks'], queryFn: async () => await mocks() })
    const add = useMutation({ mutationFn: async (i: NewMock) => await create([i]), 
        onSuccess: () => client.invalidateQueries({ queryKey: ['webhook-mocks'] }) })
    const form = useForm({
        name: 'new-webhook-mock',
        initialValues: {
            request: {
                method: 'POST',
                path: '/',
            },
            response: { status: 204 },
        } as NewMock,
        validate: {
            request: {
                path: (value: string) => value ? null : 'Invalid path',
            }
        },
    });
    return <SearchList value={search} setValue={setSearch}
        action={<NewModal title="New Mock">
            <form onSubmit={form.onSubmit((values) => add.mutate(values))} onReset={() => form.reset()} >
                <TextInput label='Request Method' withAsterisk {...form.getInputProps('request.method')} />
                <TextInput label='Request Path' {...form.getInputProps('request.path')} />
                <NumberInput label='Response Status'{...form.getInputProps('response.status')} />
                <Divider my="xs" />
                <Group justify="end">
                <Button variant="light" type='reset' disabled={!form.isDirty}>Reset</Button>
                <Button variant="outline" type='submit' disabled={!form.isValid()}>Create</Button>
                </Group>
            </form>
        </NewModal>}>
        {data.data?.map(({ state }) => <WithMenu key={state.id}>
            <Tooltip label={'Created At: '+ state.creation_date}>
                <Button variant="default" fullWidth>{state.id}</Button>
            </Tooltip>
        </WithMenu>)}
    </SearchList> 
}