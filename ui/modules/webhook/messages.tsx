import { listMessages, createMessage } from "@/provider/svix"
import { CreationModal, QueryList } from "@/components/crud"
import { Button, JsonInput, TextInput, Tooltip } from "@mantine/core"
import {  WithMenu } from "@/components/components"
import { MessageIn, MessageOut } from "svix"
import { useRouter } from "next/navigation"

export type MessageListProps = {
    appId: string
    search: string
    setSearch: (s: string) => void 
}
export const MessageList = ({ appId, search, setSearch }: MessageListProps) => { 
    const router = useRouter()
    const queryKey = ["webhook", appId, "messages"]
    return <QueryList<MessageOut> value={search} setValue={setSearch} 
        queryKey={queryKey} queryFn={async () => await listMessages(appId) }
        action={<CreationModal<MessageIn> title="New Message" queryKey={queryKey} 
            name='new-webhook-message' initialValues={{ eventType: '', payload: '{}'}} validate={{
                eventType: (value: string) => value ? null : 'Invalid eventType',
            }}
            mutationFn={async ({ payload, ...rest }) => await createMessage(appId, { ...rest, payload: JSON.parse(payload) })}>{(form) => <>
                <TextInput label='EventType' withAsterisk {...form.getInputProps('eventType')} />
                <JsonInput label='Payload' {...form.getInputProps('payload')}/>        
            </>}</CreationModal>}>
        {({ id, timestamp }) => <WithMenu key={id}>
            <Tooltip label={'Created At: '+ timestamp}>
                <Button onClick={() => 
                    router.push('/webhook/applications/'+appId+'/messages/'+id)} variant="default" fullWidth>{id}</Button>
            </Tooltip>
        </WithMenu>}
    </QueryList> 
}