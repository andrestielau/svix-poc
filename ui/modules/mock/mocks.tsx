import { NewMock, Mock } from "@/provider/smocker/client"
import { CreationModal, QueryList } from "@/components/crud"
import { mocks, create } from "@/provider/smocker"
import { Button, NumberInput, TextInput, Tooltip } from "@mantine/core"
import { WithMenu } from "@/components/components"
import { useRouter } from "next/navigation"

const queryKey = ["mock", "mocks"]
export type MockListProps = {
    search: string
    setSearch: (s: string) => void 
}
export const MockList = ({ search, setSearch }: MockListProps) => { 
    const router = useRouter()
    return <QueryList<Mock> value={search} setValue={setSearch} 
        queryKey={queryKey} queryFn={async () => await mocks() }
        action={<CreationModal<NewMock> title="New Mock" queryKey={queryKey} 
            name='new-mock-mock' initialValues={{
                request: { method: 'POST', path: '/' },
                response: { status: 204 },
            }}
            mutationFn={async (i) => await create([i])}>{(form) => <>
                <TextInput label='Request Method' withAsterisk {...form.getInputProps('request.method')} />
                <TextInput label='Request Path' {...form.getInputProps('request.path')} />
                <NumberInput label='Response Status'{...form.getInputProps('response.status')} />      
            </>}</CreationModal>}>
        {({ state: { id, creation_date } }) => <WithMenu key={id}>
            <Tooltip label={'Created At: '+ creation_date}>
                <Button onClick={() => 
                    router.push('/mock/mocks/'+id)} variant="default" fullWidth>{id}</Button>
            </Tooltip>
        </WithMenu>}
    </QueryList> 
}