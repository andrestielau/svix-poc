import { CreationModal, QueryList } from "@/components/crud"
import { Button, TextInput, Tooltip } from "@mantine/core"
import { WithMenu } from "@/components/components"
import { useRouter } from "next/navigation"
import { createSubscription, subscriptions } from "@/provider/pubsub"
import { SubscriptionMetadata } from "@google-cloud/pubsub"

const queryKey = ["pubsub", "subscriptions"]
export type SubscriptionListProps = {
    search: string
    setSearch: (s: string) => void 
}
export const SubscriptionList = ({ search, setSearch }: SubscriptionListProps) => { 
    const router = useRouter()
    return <QueryList<SubscriptionMetadata> value={search} setValue={setSearch} 
        queryKey={queryKey} queryFn={async () => await subscriptions() }
        action={<CreationModal<{ name: string, topic: string }> title="New Subscription" queryKey={queryKey} 
            name='new-pubsub-subscription' initialValues={{ name: '', topic: '' }} validate={{
                name: (value: string) => value ? null : 'Invalid name',
                topic: (value: string) => value ? null : 'Invalid topic',
            }}
            mutationFn={async (i) => await createSubscription(i.topic, i.name)}>{(form) => <>
                <TextInput label='Name' withAsterisk {...form.getInputProps('name')} />     
            </>}</CreationModal>}>
        {({ name }) => <WithMenu key={name}>
            <Tooltip label={''}>
                <Button  variant="default" fullWidth onClick={() => 
                    router.push('/pubsub/subscription/'+name?.replaceAll('projects/demo/subscriptions/', ''))}>
                        {name?.replaceAll('projects/demo/subscriptions/', '')}
                    </Button>
            </Tooltip>
        </WithMenu>}
    </QueryList> 
}