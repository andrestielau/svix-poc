import { NewSession, SessionDetails } from "@/provider/smocker/client"
import { CreationModal, QueryList } from "@/components/crud"
import { sessions, startSession } from "@/provider/smocker"
import { Button, TextInput, Tooltip } from "@mantine/core"
import { WithMenu } from "@/components/components"
import { useRouter } from "next/navigation"

const queryKey = ["mock", "sessions"]
export type SessionListProps = {
    search: string
    setSearch: (s: string) => void 
}
export const SessionList = ({ search, setSearch }: SessionListProps) => { 
    const router = useRouter()
    return <QueryList<SessionDetails> value={search} setValue={setSearch} accessor={({ id }) => id}
        queryKey={queryKey} queryFn={async () => await sessions() }
        action={<CreationModal<NewSession> title="New Session" queryKey={queryKey} 
            name='new-mock-session'
            mutationFn={async (i) => await startSession(i)}>{(form) => <>
                <TextInput label='Name' withAsterisk {...form.getInputProps('name')} />
                <TextInput label='Description' {...form.getInputProps('description')} />        
            </>}</CreationModal>}>
        {({ id, name, date }) => <WithMenu id={id} key={id}>
            <Tooltip label={'Created At: '+ date}>
                <Button onClick={() => 
                    router.push('/mock/sessions/'+id)} variant="default" fullWidth>{name || id}</Button>
            </Tooltip>
        </WithMenu>}
    </QueryList> 
}