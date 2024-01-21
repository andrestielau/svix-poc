"use client"
import { EventTypeList } from "@/modules/webhook/event-types"
import { Title } from "@mantine/core"
import { useState } from "react"

const EventTypesPage = () => {
    const [searchApps, setSearchApps] = useState('')
    return <>
        <Title order={1}>Event Types</Title>
        <EventTypeList search={searchApps} setSearch={setSearchApps}/>  
    </>
}

export default EventTypesPage