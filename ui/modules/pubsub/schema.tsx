import { CreationModal, QueryList } from "@/components/crud"
import { Button, TextInput, Tooltip } from "@mantine/core"
import { WithMenu } from "@/components/components"
import { useRouter } from "next/navigation"
import { createSchema, schemas } from "@/provider/pubsub"
import { ISchema, SchemaType } from "@google-cloud/pubsub"

const queryKey = ["pubsub", "schemas"]
export type SchemaListProps = {
    search: string
    setSearch: (s: string) => void 
}
export const SchemaList = ({ search, setSearch }: SchemaListProps) => { 
    const router = useRouter()
    return <QueryList<ISchema> value={search} setValue={setSearch} 
        queryKey={queryKey} queryFn={async () => await schemas() }
        action={<CreationModal<{ id: string, def: string, type?: SchemaType }> title="New Schema" queryKey={queryKey} 
            name='new-pubsub-schema' initialValues={{ id: '', def: '' }} validate={{
                def: (value: string) => value ? null : 'Invalid def',
                id: (value: string) => value ? null : 'Invalid id',
            }}
            mutationFn={async (i) => await createSchema(i.id, i.def, i.type)}>{(form) => <>
                <TextInput label='Id' withAsterisk {...form.getInputProps('id')} />
                <TextInput label='Definition' withAsterisk {...form.getInputProps('def')} />
            </>}</CreationModal>}>
        {({ name_, id }) => <WithMenu key={id}>
            <Tooltip label={name_}>
                <Button onClick={() => 
                    router.push('/pubsub/schemas/'+id)} variant="default" fullWidth>{id} {name_}</Button>
            </Tooltip>
        </WithMenu>}
    </QueryList> 
}