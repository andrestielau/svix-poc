import { NewNotificationType, NotificationType } from "@/provider/drizzle"
import { addNotificationType, listNotificationTypes, listProviders, remNotificationType } from "@/provider/drizzle/queries"
import { CreationModal, DeleteMenuItem, QueryList } from "@/components/crud"
import { Button, NativeSelect, TextInput, Tooltip } from "@mantine/core"
import { WithMenu } from "@/components/components"
import { useRouter } from "next/navigation"
import { useQuery } from "@tanstack/react-query"

const queryKey = ["router", "notification-types"]
export type NotificationTypeListProps = {
    search: string
    setSearch: (s: string) => void 
}
export const NotificationTypeList = ({ search, setSearch }: NotificationTypeListProps) => { 
    const router = useRouter()
    const providers = useQuery({queryKey: ['router','providers'], queryFn: async () => await listProviders() })
    return <QueryList<NotificationType> value={search} setValue={setSearch} 
        queryKey={queryKey} queryFn={async () => await listNotificationTypes() }
        action={<CreationModal<NewNotificationType> title="New Notification Type" queryKey={queryKey} 
            name='new-router-notification-type' initialValues={{ provider: providers.data?.at(0)?.id || '', id: '' }} validate={{
                provider: (value: string) => value ? null : 'Invalid provider',
                id: (value: string) => value ? null : 'Invalid id',
            }}
            mutationFn={async (i) => await addNotificationType(i)}>{(form) => <>
                <TextInput  label='Id' withAsterisk {...form.getInputProps('id')} />    
                <NativeSelect  label='Provider' withAsterisk data={['', ...(providers.data?.map(p => p.id) || [])]} {...form.getInputProps('provider')} />    
            </>}</CreationModal>}>
        {({ id, createdAt }) => <WithMenu id={id} key={id}
            danger={<DeleteMenuItem queryKey={queryKey} mutationFn={() => remNotificationType(id)}/>}>
            <Tooltip label={'Created At: '+ createdAt}>
                <Button onDoubleClick={() => router.push('/router/notification-types/'+id)}  variant="default" fullWidth>{id}</Button>
            </Tooltip>
        </WithMenu>}
    </QueryList> 
}