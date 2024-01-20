import { CreationModal, QueryList } from "@/components/crud"
import { Button, TextInput, Tooltip } from "@mantine/core"
import { WithMenu } from "@/components/components"
import { useRouter } from "next/navigation"
import { createTopic, topics } from "@/provider/pubsub"
import { Topic } from "@google-cloud/pubsub"

const queryKey = ["pubsub", "topics"]
export type TopicListProps = {
    search: string
    setSearch: (s: string) => void 
}
export const TopicList = ({ search, setSearch }: TopicListProps) => { 
    const router = useRouter()
    return <QueryList<Topic> value={search} setValue={setSearch} 
        queryKey={queryKey} queryFn={async () => await topics() }
        action={<CreationModal<{ name: string }> title="New Topic" queryKey={queryKey} 
            name='new-pubsub-topic' initialValues={{ name: '' }} validate={{
                name: (value: string) => value ? null : 'Invalid name',
            }}
            mutationFn={async (i) => await createTopic(i.name)}>{(form) => <>
                <TextInput label='Name' withAsterisk {...form.getInputProps('name')} />     
            </>}</CreationModal>}>
        {({ name }) => <WithMenu key={name}>
            <Tooltip label={''}>
                <Button onClick={() => 
                    router.push('/pubsub/topics/'+name)} variant="default" fullWidth>{name}</Button>
            </Tooltip>
        </WithMenu>}
    </QueryList> 
}