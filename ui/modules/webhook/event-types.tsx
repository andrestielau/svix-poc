import { listEventTypes, createEventType, deleteEventType } from "@/provider/svix"
import { CreationModal, DeleteMenuItem, QueryList } from "@/components/crud"
import { Button, InputWrapper, TextInput, Tooltip } from "@mantine/core"
import { EventTypeIn, EventTypeOut } from "svix"
import { WithMenu } from "@/components/components"
import { useRouter } from "next/navigation"
import { Editor } from "@monaco-editor/react"

const queryKey = ["webhook", "event-types"]
export type EventTypeListProps = {
    search: string
    setSearch: (s: string) => void 
}
export const EventTypeList = ({ search, setSearch }: EventTypeListProps) => { 
    const router = useRouter()
    return <QueryList<EventTypeOut> value={search} setValue={setSearch} accessor={({ name }) => name}
        queryKey={queryKey} queryFn={async () => await listEventTypes() }
        action={<CreationModal<EventTypeIn> title="New EventType" queryKey={queryKey} 
            name='new-webhook-event-type' initialValues={{ name: '', description: '', schemas: { 0: {}} }} validate={{
                name: (value: string) => value ? null : 'Invalid name',
            }}
            mutationFn={async (i) => await createEventType(i)}>{(form) => <>
                <TextInput label='Name' withAsterisk {...form.getInputProps('name')} />
                <TextInput label='Description' {...form.getInputProps('description')} /> 
                <InputWrapper label="Schema" withAsterisk>
                    <Editor height="30vh" theme='vs-dark' defaultLanguage="json" {...form.getInputProps('schemas.0')} onValidate={console.log}/>
                </InputWrapper>       
            </>}</CreationModal>}>
        {({ name, createdAt }) => <WithMenu id={name} key={name}
            danger={<DeleteMenuItem queryKey={queryKey} mutationFn={() => deleteEventType(name)}/>}>
            <Tooltip label={'Created At: '+ createdAt}>
                <Button onDoubleClick={() => router.push('/webhook/event-types/'+name)} variant="default" fullWidth>{name}</Button>
            </Tooltip>
        </WithMenu>}
    </QueryList> 
}

