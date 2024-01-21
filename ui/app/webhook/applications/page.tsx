"use client"

import { ApplicationList } from "@/modules/webhook/applications"
import { Title } from "@mantine/core"
import { useState } from "react"

const AppsPage = () => {
    const [searchApps, setSearchApps] = useState('')
    return <>
        <Title order={1}>Applications</Title>
        <ApplicationList search={searchApps} setSearch={setSearchApps}/>  
    </>
}
  
export default AppsPage