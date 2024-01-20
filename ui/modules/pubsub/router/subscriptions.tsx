import { NewSubscription, Subscription } from "@/provider/drizzle/schema"
import { addSubscription, listSubscriptions } from "@/provider/drizzle/queries"
import { CreationModal, QueryList } from "@/components/crud"
import { Button, TextInput, Tooltip } from "@mantine/core"
import { WithMenu } from "@/components/components"
import { useRouter } from "next/navigation"

const queryKey = ["router", "subscriptions"]
export type SubscriptionListProps = {
    search: string
    setSearch: (s: string) => void 
}
export const SubscriptionList = ({ search, setSearch }: SubscriptionListProps) => { 
    const router = useRouter()
    return <QueryList<Subscription> value={search} setValue={setSearch} 
        queryKey={queryKey} queryFn={async () => await listSubscriptions() }
        action={<CreationModal<NewSubscription> title="New Subscription" queryKey={queryKey} 
            name='new-router-subscription' initialValues={{ notificationType: '', tenantId: '', metadata: '' }} validate={{
                notificationType: (value: string) => value ? null : 'Invalid notificationType',
            }}
            mutationFn={async (i) => await addSubscription(i)}>{(form) => <>
                <TextInput label='Name' withAsterisk {...form.getInputProps('name')} />
                <TextInput label='Description' {...form.getInputProps('description')} />        
            </>}</CreationModal>}>
        {({ uid, createdAt, ...rest }) => <WithMenu key={uid}>
            <Tooltip label={<span>Created At: {createdAt}<br/>{JSON.stringify(rest)}</span>} >
                <Button onClick={() => router.push('/router/subscriptions/'+uid)} variant="default" fullWidth>{uid}</Button>
            </Tooltip>
        </WithMenu>}
    </QueryList> 
}