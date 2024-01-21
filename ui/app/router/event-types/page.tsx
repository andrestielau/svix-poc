"use client"
import { Title } from "@mantine/core";
import { EventTypeList } from "@/modules/router/event-types";
import { useState } from "react";

const EventTypesPage = () => {
    const [searchEventTypes, setSearchEventTypes] = useState('')
    return <>
        <Title order={1}>Event Types</Title>
        <EventTypeList search={searchEventTypes} setSearch={setSearchEventTypes}/>
    </>
}
export default EventTypesPage