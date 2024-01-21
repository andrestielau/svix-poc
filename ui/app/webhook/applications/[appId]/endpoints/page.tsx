"use client"

import { EndpointList } from "@/modules/webhook/endpoints"
import { Title } from "@mantine/core"
import { useState } from "react"

const EndpointsPage = ({ params: { appId} }: { params: { appId: string } }) => {
    const [searchEndpoints, setSearchEndpoints] = useState('')
    return <>
        <Title order={1}>Endpoints</Title>
        <EndpointList search={searchEndpoints} setSearch={setSearchEndpoints} appId={appId} />  
    </>
}
  
export default EndpointsPage