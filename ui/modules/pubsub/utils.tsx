import { publish, status, topics } from "@/provider/pubsub"
import { MessageOptions } from "@google-cloud/pubsub/build/src/topic"
import { ActionIcon, Button, CloseButton, Divider, Flex, Group, InputWrapper, JsonInput, Modal, NativeSelect, Stack, TextInput } from "@mantine/core"
import { useForm } from "@mantine/form"
import { useDisclosure } from "@mantine/hooks"
import { useMutation, useQuery } from "@tanstack/react-query"
import { IconDotsVertical, IconX } from "@tabler/icons-react"
import { UseFormReturnType } from "@mantine/form/lib/types"
import { v4 } from "uuid"

export const StatusTooltip = () => {
    const req = useQuery({
        queryKey: ['pubsub', 'status'],
        queryFn: () => status()
    })
    return <JsonInput value={JSON.stringify(req.data || {}, undefined, ' ')} disabled autosize/>
}
type Publish = {
    topic: string
    msg: MessageOptions
}&{attributes: [string, string, string][]}
export const PublishModal = () => {
    const t = useQuery({queryKey: ['pubsub','topics'], queryFn: async () => await topics() })
    const [opened, { open, close }] = useDisclosure(false)
    const mut = useMutation<unknown, Error, Publish>({
        mutationFn: ({ topic, msg, attributes }) => publish(topic, { attributes: Object.fromEntries(attributes), ...msg })
    })
    const form = useForm<Publish>({
        initialValues: {
            topic: '',
            msg: {},
            attributes: []
        },
        validate: {
            topic: (value) => value ? null : 'Invalid topic',
        }
    })
    return <>
    <Button color="yellow" size="sm" onClick={open} fullWidth>Publish</Button>
    <Modal opened={opened} onClose={close} title={"Publish"} centered>
        <form onSubmit={form.onSubmit((values) => mut.mutate(values))} onReset={() => form.reset()} >
            <NativeSelect label='Topic' withAsterisk data={['', ...(t.data?.map(p => p!.name!) || [])]} {...form.getInputProps('topic')} />
            <JsonInput label='Payload' withAsterisk {...form.getInputProps('msg.json')}  />
            <InputWrapper><AttributeList form={form} reorder={(from: number, to: number) => form.reorderListItem('attributes', {from, to})}/></InputWrapper>
            <Divider my="xs" />
            <Group justify="end">
                <Button variant="light" type='reset' disabled={!form.isDirty}>Reset</Button>
                <Button variant="outline" type='submit'>Publish</Button>
            </Group>
        </form>
    </Modal>
</>
}
const AttributeList = ({ form, reorder }: { form: UseFormReturnType<Publish>, reorder: (from: number, to: number) => void }) => {
    const { attributes } = form.values
    return <Stack gap='0'>{ attributes.map(([k, v, key], i) => <Flex key={key}
        onDrop={(ev) => drop(ev, i, attributes.length, reorder)} onDragOver={onDragOver} 
        onDragEnter={onDragEnter} onDragLeave={onDragLeave} onDragEnd={onDragEnd}  className="attribute-item">
        <ActionIcon variant='default' size='lg' style={{ pointerEvents: 'all' }} onDragStart={(ev) => drag(ev, i)} draggable><IconDotsVertical/></ActionIcon>
        <TextInput placeholder='Key' rightSectionPointerEvents="all"
            rightSection={<CloseButton 
                aria-label="Clear key" 
                onClick={() => form.setFieldValue('attributes.'+i+'.0', '')} 
                style={{ display: attributes[i][0] ? undefined : 'none' }} />}
                {...form.getInputProps('attributes.'+i+'.0')}/>
        <TextInput placeholder='Value' rightSectionPointerEvents="all"
            rightSection={<CloseButton
                aria-label="Clear value" 
                onClick={() => form.setFieldValue('attributes.'+i+'.1', '')} 
                style={{ display: attributes[i][1] ? undefined : 'none' }} />}
                {...form.getInputProps('attributes.'+i+'.1')} />
        <ActionIcon variant='default' size='lg' onClick={() => form.removeListItem('attributes', i)}><IconX/></ActionIcon>
    </Flex>)}
    <Button variant='default' fullWidth onClick={() => 
        form.insertListItem('attributes', ['', '', v4()], attributes.length)}>Add Attribute</Button>
        {JSON.stringify(Object.fromEntries(attributes))}
</Stack>
}
var prev: any
function onDragOver(ev: any) {
    ev.preventDefault()
}
function onDragEnter(ev: any) {
    ev.preventDefault()
    const c = ev.target.closest(".attribute-item")
    if (c === prev) return
    if (prev) prev.style.border = "";
    prev = c
    if (c) c.style.border = "3px dotted red";
}
function onDragLeave(ev: any) {
    ev.preventDefault()
}
function onDragEnd() { 
    if (prev) prev.style.border = "";
    prev = undefined 
}
function drag(ev: any, id: number) {
    ev.dataTransfer.setData("text/plain", id.toString());
    ev.dataTransfer.dropEffect = "move";
}

function drop(ev: any, id: number, size: number, reorder: (from: number, to: number) => void) {
    var data = ev.dataTransfer?.getData("text/plain");
    const from = parseInt(data)
    if (id === from) return 
    reorder(from, id)
}