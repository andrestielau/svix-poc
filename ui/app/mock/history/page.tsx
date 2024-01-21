"use client"
import { HistoryList } from "@/modules/mock/history"
import { Title } from "@mantine/core"
import { useState } from "react"

const HistoryPage = () => {
    const [searchHistory, setSearchHistory] = useState('')
    return <>
        <Title order={1}>Sessions</Title>
        <HistoryList search={searchHistory} setSearch={setSearchHistory}/>
    </>
}

export default HistoryPage