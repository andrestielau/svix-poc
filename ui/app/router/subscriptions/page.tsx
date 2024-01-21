"use client"
import { Title } from "@mantine/core";
import { SubscriptionList } from "@/modules/router/subscriptions";
import { useState } from "react";

const SubscriptionPage = () => {
    const [searchSubscriptions, setSearchSubscriptions] = useState('')
    return <>
        <Title order={1}>Subscriptions</Title>
        <SubscriptionList search={searchSubscriptions} setSearch={setSearchSubscriptions}/>
    </>
}
export default SubscriptionPage