"use client"

import { MockList } from "@/modules/mocks/mocks";
import { SessionList } from "@/modules/mocks/sessions";
import { Grid, Title } from "@mantine/core";
import { useState } from "react";

const MocksPage = async () => {
  const [searchMocks, setSearchMocks] = useState('')
  const [searchSessions, setSearchSessions] = useState('')
  return (
    <Grid>
      <Grid.Col span={4}>
        <Title order={1}>Mocks</Title>
        <MockList search={searchMocks} setSearch={setSearchMocks}/>
      </Grid.Col>
      <Grid.Col span={4}>
        <Title order={1}>Sessions</Title>
        <SessionList search={searchSessions} setSearch={setSearchSessions}/>
      </Grid.Col>
    </Grid>
  );
}
export default MocksPage