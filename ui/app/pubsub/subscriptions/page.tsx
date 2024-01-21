"use client"
import { Title } from "@mantine/core";
import { useState } from "react";
import { SubscriptionList } from "@/modules/pubsub/subscriptions";


const SubscriptionsPage = () => {
    const [searchSubscriptions, setSearchSubscriptions] = useState('')
    return <>
        <Title order={1}>Subscriptions</Title>
        <SubscriptionList search={searchSubscriptions} setSearch={setSearchSubscriptions}/>
    </>
}

export default SubscriptionsPage