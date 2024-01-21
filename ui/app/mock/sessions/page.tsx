"use client"
import { SessionList } from "@/modules/mock/sessions"
import { Title } from "@mantine/core"
import { useState } from "react"

const SessionsPage = () => {
    const [searchSessions, setSearchSessions] = useState('')
    return <>
        <Title order={1}>Sessions</Title>
        <SessionList search={searchSessions} setSearch={setSearchSessions}/>
    </>
}

export default SessionsPage