"use client"
import { Grid, GridCol, Title } from "@mantine/core";
import { ProviderList } from "@/modules/router/providers";
import { EventTypeList } from "@/modules/router/event-types";
import { NotificationTypeList } from "@/modules/router/notification-types";
import { useState } from "react";

const RouterPage = () => {
  const [searchProviders, setSearchProviders] = useState('')
  const [searchEventTypes, setSearchEventTypes] = useState('')
  const [searchNotificationTypes, setSearchNotificationTypes] = useState('')
  return (
    <Grid>
      <GridCol span={4}>
        <Title order={1}>Providers</Title>
        <ProviderList search={searchProviders} setSearch={setSearchProviders}/>
      </GridCol>
      <GridCol span={4}>
        <Title order={1}>Event Types</Title>
        <EventTypeList search={searchEventTypes} setSearch={setSearchEventTypes}/>
      </GridCol>
      <GridCol span={4}>
        <Title order={1}>Notification Types</Title>
        <NotificationTypeList search={searchNotificationTypes} setSearch={setSearchNotificationTypes}/>
      </GridCol>
    </Grid>
  );
}
export default RouterPage