"use client"
import { Grid, GridCol, Title } from "@mantine/core";
import { ProviderList } from "@/modules/router/providers";
import { EventTypeList } from "@/modules/router/event-types";
import { NotificationTypeList } from "@/modules/router/notification-types";
import { useState } from "react";
import { useRouter } from "next/navigation";
import { SubscriptionList } from "@/modules/router/subscriptions";

const RouterPage = () => {
  const router = useRouter()
  const [searchProviders, setSearchProviders] = useState('')
  const [searchEventTypes, setSearchEventTypes] = useState('')
  const [searchSubscriptions, setSearchSubscriptions] = useState('')
  const [searchNotificationTypes, setSearchNotificationTypes] = useState('')
  return (
    <Grid>
      <GridCol span={3}>
        <Title order={1} onDoubleClick={() => router.push('/router/providers')}>Providers</Title>
        <ProviderList search={searchProviders} setSearch={setSearchProviders}/>
      </GridCol>
      <GridCol span={3}>
        <Title order={1} onDoubleClick={() => router.push('/router/event-types')}>Event Types</Title>
        <EventTypeList search={searchEventTypes} setSearch={setSearchEventTypes}/>
      </GridCol>
      <GridCol span={3}>
        <Title order={1} onDoubleClick={() => router.push('/router/notification-types')}>Notification Types</Title>
        <NotificationTypeList search={searchNotificationTypes} setSearch={setSearchNotificationTypes}/>
      </GridCol>
      <GridCol span={3}>
        <Title order={1} onDoubleClick={() => router.push('/router/subscriptions')}>Subscriptions</Title>
        <SubscriptionList search={searchSubscriptions} setSearch={setSearchSubscriptions}/>
      </GridCol>
    </Grid>
  );
}
export default RouterPage