import { Button, InputWrapper, NativeSelect, TextInput, Tooltip } from "@mantine/core"
import { CreationModal, QueryList } from "@/components/crud"
import { ISchema, SchemaType } from "@google-cloud/pubsub"
import { createSchema, schemas } from "@/provider/pubsub"
import { WithMenu } from "@/components/components"
import { useRouter } from "next/navigation"
import Editor, { useMonaco } from '@monaco-editor/react';
import { useEffect } from "react"

const protoDefault = `syntax = "proto3";
message {Name} {
    string id = 1;
}`
const queryKey = ["pubsub", "schemas"]
export type SchemaListProps = {
    search: string
    setSearch: (s: string) => void 
}

export const SchemaList = ({ search, setSearch }: SchemaListProps) => { 
    const router = useRouter()
    const monaco = useMonaco()

    return <QueryList<ISchema> value={search} setValue={setSearch} 
        queryKey={queryKey} queryFn={async () => await schemas() }
        action={<CreationModal<{ id: string, def: string, type?: SchemaType }> title="New Schema" queryKey={queryKey} 
            name='new-pubsub-schema' initialValues={{ id: '', type: "PROTOCOL_BUFFER", def: protoDefault }} validate={{
                type: (value) => !(value && schemaTypes.findIndex((v) => v.value === value) != -1) && 'invalid type',
                def: (value) => value ? null : 'Invalid def',
                id: (value) => value ? null : 'Invalid id',
            }}
            onReset={() => monaco?.editor.getEditors()[0].setValue(protoDefault)}
            mutationFn={async (i) => await createSchema(i.id, i.def, i.type)}>{(form) => <>
                <TextInput label='Id' withAsterisk {...form.getInputProps('id')} />
                <NativeSelect label="Type" data={schemaTypes}  {...form.getInputProps('type')} />
                <InputWrapper label="Definition" withAsterisk error={!!monaco?.editor?.getModelMarkers({}).length && 'Invalid definition'}>
                    <Editor height="30vh" theme='vs-dark' defaultLanguage="proto" {...form.getInputProps('def')} onValidate={console.log}/>
                </InputWrapper>
            </>}</CreationModal>}>
        {({ type, name, revisionCreateTime }) => <WithMenu key={name}>
            <Tooltip label={<label>Created At: {revisionCreateTime?.nanos}</label>}>
                <Button onClick={() => 
                    router.push('/pubsub/schema/'+name?.replaceAll('projects/demo/schemas/', ''))} variant="default" fullWidth>
                        {type}: {name?.replaceAll('projects/demo/schemas/', '')}
                </Button>
            </Tooltip>
        </WithMenu>}
    </QueryList> 
}
const schemaTypes = [
    { value: "PROTOCOL_BUFFER", label: 'Proto' },
    { value: "AVRO", label: 'Avro' },
]