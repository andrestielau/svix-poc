import { NewEventType, EventType } from "@/provider/drizzle/schema"
import { addEventType, listEventTypes } from "@/provider/drizzle/queries"
import { CreationModal, DeleteMenuItem, QueryList } from "@/components/crud"
import { Button, TextInput, Tooltip } from "@mantine/core"
import { WithMenu } from "@/components/components"
import { useRouter } from "next/navigation"

const queryKey = ["router", "event-types"]
export type EventTypeListProps = {
    search: string
    setSearch: (s: string) => void 
}
export const EventTypeList = ({ search, setSearch }: EventTypeListProps) => { 
    const router = useRouter()
    return <QueryList<EventType> value={search} setValue={setSearch} 
        queryKey={queryKey} queryFn={async () => await listEventTypes() }
        action={<CreationModal<NewEventType> title="New Event Type" queryKey={queryKey} 
            name='new-router-event-type' initialValues={{ schema: '', id: '' }} validate={{
                schema: (value: string) => value ? null : 'Invalid schema',
                id: (value: string) => value ? null : 'Invalid id',
            }}
            mutationFn={async (i) => await addEventType(i)}>{(form) => <>
                <TextInput  label='Id' withAsterisk {...form.getInputProps('id')} />    
            </>}</CreationModal>}>
        {({ id, createdAt }) => <WithMenu id={id} key={id}
            danger={<DeleteMenuItem queryKey={queryKey} mutationFn={() => deleteEventType(id)}/>}>
            <Tooltip label={'Created At: '+ createdAt}>
                <Button onDoubleClick={() => router.push('/router/event-types/'+id)}  variant="default" fullWidth>{id}</Button>
            </Tooltip>
        </WithMenu>}
    </QueryList> 
}