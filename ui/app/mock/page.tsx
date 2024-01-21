"use client"

import { HistoryList } from "@/modules/mock/history";
import { MockList } from "@/modules/mock/mocks";
import { SessionList } from "@/modules/mock/sessions";
import { Grid, Title } from "@mantine/core";
import { useRouter } from "next/navigation";
import { useState } from "react";

const MocksPage = () => {
  const router = useRouter()
  const [searchMocks, setSearchMocks] = useState('')
  const [searchHistory, setSearchHistory] = useState('')
  const [searchSessions, setSearchSessions] = useState('')
  return (
    <Grid>
      <Grid.Col span={4}>
        <Title order={1} onDoubleClick={() => router.push('/mock/mocks')}>Mocks</Title>
        <MockList search={searchMocks} setSearch={setSearchMocks}/>
      </Grid.Col>
      <Grid.Col span={4}>
        <Title order={1} onDoubleClick={() => router.push('/mock/mocks')}>Sessions</Title>
        <SessionList search={searchSessions} setSearch={setSearchSessions}/>
      </Grid.Col>
      <Grid.Col span={4}>
        <Title order={1} onDoubleClick={() => router.push('/mock/history')}>History</Title>
        <HistoryList search={searchHistory} setSearch={setSearchHistory}/>
      </Grid.Col>
    </Grid>
  );
}
export default MocksPage