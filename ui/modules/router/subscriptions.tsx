import { NewSubscription, Subscription } from "@/provider/drizzle"
import { addSubscription, listNotificationTypes, listSubscriptions, remSubscription } from "@/provider/drizzle/queries"
import { CreationModal, DeleteMenuItem, QueryList } from "@/components/crud"
import { Button, JsonInput, NativeSelect, TextInput, Tooltip } from "@mantine/core"
import { WithMenu } from "@/components/components"
import { useRouter } from "next/navigation"
import { useQuery } from "@tanstack/react-query"

const queryKey = ["router", "subscriptions"]
export type SubscriptionListProps = {
    search: string
    setSearch: (s: string) => void 
}
export const SubscriptionList = ({ search, setSearch }: SubscriptionListProps) => { 
    const router = useRouter()
    const notificationTypes = useQuery({queryKey: ['router','notification-types'], queryFn: async () => await listNotificationTypes() })
    return <QueryList<Subscription> value={search} setValue={setSearch} 
        queryKey={queryKey} queryFn={async () => await listSubscriptions() }
        action={<CreationModal<NewSubscription> title="New Subscription" queryKey={queryKey} 
            name='new-router-subscription' initialValues={{ notificationType: '', tenantId: '', metadata: '{}' }} validate={{
                notificationType: (value: string) => value ? null : 'Invalid notificationType',
            }}
            mutationFn={async (i) => await addSubscription(i)}>{(form) => <>
                <NativeSelect label='Notification Type' withAsterisk data={['', ...(notificationTypes.data?.map(p => p.id) || [])]} {...form.getInputProps('notificationType')} />   
                <TextInput label='Tenant' withAsterisk {...form.getInputProps('tenantId')} />
                <JsonInput label='Metadata' {...form.getInputProps('metadata')} />        
            </>}</CreationModal>}>
        {({ uid, createdAt, ...rest }) => <WithMenu id={uid} key={uid}
            danger={<DeleteMenuItem queryKey={queryKey} mutationFn={() => remSubscription(uid)}/>}>
            <Tooltip label={<span>Created At: {createdAt}<br/>{JSON.stringify(rest)}</span>} >
                <Button onDoubleClick={() => router.push('/router/subscriptions/'+uid)} variant="default" fullWidth>{uid}</Button>
            </Tooltip>
        </WithMenu>}
    </QueryList> 
}