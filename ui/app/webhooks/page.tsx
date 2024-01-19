"use client"
import { ApplicationList } from "@/modules/webhooks/applications";
import { EventTypeList } from "@/modules/webhooks/event-types";
import { Grid, Title } from "@mantine/core";
import { useState } from "react";


export default function WebhoksPage() {
  const [searchApps, setSearchApps] = useState('')
  const [searchEventTypes, setSearchEventTypes] = useState('')
  return (
    <Grid>
      <Grid.Col span={4}>
        <Title order={1}>Applications</Title>
        <ApplicationList search={searchApps} setSearch={setSearchApps}/>
      </Grid.Col>
      <Grid.Col span={4}>
        <Title order={1}>Event Types</Title>
        <EventTypeList search={searchEventTypes} setSearch={setSearchEventTypes}/>
      </Grid.Col>
    </Grid>
  );
}
