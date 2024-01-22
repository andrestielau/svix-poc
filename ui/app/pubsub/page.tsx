"use client"
import { useRouter } from "next/navigation";
import { Grid, Title } from "@mantine/core";
import { useState } from "react";
import { TopicList } from "@/modules/pubsub/topics";
import { SchemaList } from "@/modules/pubsub/schema";
import { SubscriptionList } from "@/modules/pubsub/subscriptions";


const PubsubPage = () => {
  const router = useRouter()
  const [searchTopics, setSearchTopics] = useState('')
  const [searchSchemas, setSearchSchemas] = useState('')
  const [searchSubscriptions, setSearchSubscriptions] = useState('')
  return (
    <Grid>
      <Grid.Col span={4}>
        <Title order={1} onDoubleClick={() => router.push('/pubsub/topic')}>Topics</Title>
        <TopicList search={searchTopics} setSearch={setSearchTopics}/>
      </Grid.Col>
      <Grid.Col span={4}>
        <Title order={1} onDoubleClick={() => router.push('/pubsub/schema')}>Schemas</Title>
        <SchemaList search={searchSchemas} setSearch={setSearchSchemas}/>
      </Grid.Col>
      <Grid.Col span={4}>
        <Title order={1} onDoubleClick={() => router.push('/pubsub/schema')}>Subscriptions</Title>
        <SubscriptionList search={searchSubscriptions} setSearch={setSearchSubscriptions}/>
      </Grid.Col>
    </Grid>
  );
}

export default PubsubPage