"use client"

import { ApplicationList } from "@/modules/webhook/applications"
import { useRouter } from "next/navigation"
import { Title } from "@mantine/core"
import { useState } from "react"

const AppsPage = () => {
    const router = useRouter()
    const [searchApps, setSearchApps] = useState('')
    return <>
        <Title order={1} onClick={() => router.push('/webhook/applications')}>Applications</Title>
        <ApplicationList search={searchApps} setSearch={setSearchApps}/>  
    </>
}
  
export default AppsPage