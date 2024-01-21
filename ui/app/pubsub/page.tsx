"use client"
import { useRouter } from "next/navigation";
import { Grid, Title } from "@mantine/core";
import { useState } from "react";
import { TopicList } from "@/modules/pubsub/topics";
import { SchemaList } from "@/modules/pubsub/schema";


const PubsubPage = () => {
  const router = useRouter()
  const [searchTopics, setSearchTopics] = useState('')
  const [searchSchemas, setSearchSchemas] = useState('')
  return (
    <Grid>
      <Grid.Col span={6}>
        <Title order={1} onDoubleClick={() => router.push('/pubsub/topic')}>Topics</Title>
        <TopicList search={searchTopics} setSearch={setSearchTopics}/>
      </Grid.Col>
      <Grid.Col span={6}>
        <Title order={1} onDoubleClick={() => router.push('/pubsub/schema')}>Schemas</Title>
        <SchemaList search={searchSchemas} setSearch={setSearchSchemas}/>
      </Grid.Col>
    </Grid>
  );
}

export default PubsubPage