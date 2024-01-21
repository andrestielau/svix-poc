"use client"
import { ApplicationList } from "@/modules/webhook/applications";
import { EventTypeList } from "@/modules/webhook/event-types";
import { useRouter } from "next/navigation";
import { Grid, Title } from "@mantine/core";
import { useState } from "react";


const WebhooksPage = () => {
  const router = useRouter()
  const [searchApps, setSearchApps] = useState('')
  const [searchEventTypes, setSearchEventTypes] = useState('')
  return (
    <Grid>
      <Grid.Col span={6}>
        <Title order={1} onDoubleClick={() => router.push('/webhook/applications')}>Applications</Title>
        <ApplicationList search={searchApps} setSearch={setSearchApps}/>
      </Grid.Col>
      <Grid.Col span={6}>
        <Title order={1} onDoubleClick={() => router.push('/webhook/event-types')}>Event Types</Title>
        <EventTypeList search={searchEventTypes} setSearch={setSearchEventTypes}/>
      </Grid.Col>
    </Grid>
  );
}

export default WebhooksPage