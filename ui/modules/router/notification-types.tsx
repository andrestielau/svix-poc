import { NewNotificationType, NotificationType } from "@/provider/drizzle/schema"
import { addNotificationType, listNotificationTypes } from "@/provider/drizzle/queries"
import { CreationModal, QueryList } from "@/components/crud"
import { Button, TextInput, Tooltip } from "@mantine/core"
import { WithMenu } from "@/components/components"
import { useRouter } from "next/navigation"

const queryKey = ["router", "notification-types"]
export type NotificationTypeListProps = {
    search: string
    setSearch: (s: string) => void 
}
export const NotificationTypeList = ({ search, setSearch }: NotificationTypeListProps) => { 
    const router = useRouter()
    return <QueryList<NotificationType> value={search} setValue={setSearch} 
        queryKey={queryKey} queryFn={async () => await listNotificationTypes() }
        action={<CreationModal<NewNotificationType> title="New Notification Type" queryKey={queryKey} 
            name='new-router-notification-type' initialValues={{ provider: '', id: '' }} validate={{
                provider: (value: string) => value ? null : 'Invalid provider',
                id: (value: string) => value ? null : 'Invalid id',
            }}
            mutationFn={async (i) => await addNotificationType(i)}>{(form) => <>
                <TextInput  label='Id' withAsterisk {...form.getInputProps('id')} />    
            </>}</CreationModal>}>
        {({ id, createdAt }) => <WithMenu key={id}>
            <Tooltip label={'Created At: '+ createdAt}>
                <Button onDoubleClick={() => router.push('/router/notification-types/'+id)}  variant="default" fullWidth>{id}</Button>
            </Tooltip>
        </WithMenu>}
    </QueryList> 
}