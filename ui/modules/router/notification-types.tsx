import { addNotificationType, listNotificationTypes, remNotificationType } from "@/provider/drizzle/queries"
import { useMutation, useQuery, useQueryClient} from "@tanstack/react-query"
import { Button, Divider, Group, TextInput, Tooltip } from "@mantine/core"
import { NewModal, SearchList, WithMenu } from "@/components/components"
import { NewNotificationType } from "@/provider/drizzle/schema"
import { useForm } from "@mantine/form"

export type NotificationTypeListProps = {
    search: string
    setSearch: (s: string) => void 
}
export const NotificationTypeList = ({ search, setSearch }: NotificationTypeListProps) => { 
    const client = useQueryClient()
    const data = useQuery({queryKey: ['notification-types'], queryFn: async () => await listNotificationTypes() })
    const add = useMutation({ mutationFn: async (i: NewNotificationType) => await addNotificationType(i), 
        onSuccess: () => client.invalidateQueries({ queryKey: ['notification-types'] }) })
    const rem = useMutation({ mutationFn: async (i: string) => await remNotificationType(i), 
        onSuccess: () => client.invalidateQueries({ queryKey: ['notification-types'] }) })
    const form = useForm({
        name: 'new-notification-type',
        initialValues: {
          id: '',
        } as NewNotificationType,
        validate: {
          id: (value: string) => value ? null : 'Invalid id',
        },
    });
    return <SearchList value={search} setValue={setSearch}
        action={<NewModal title="New Notification Type">
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