"use client"

import { MessageList } from "@/modules/webhook/messages"
import { Title } from "@mantine/core"
import { useState } from "react"

const MessagesPage = ({ params: { appId} }: { params: { appId: string } }) => {
    const [searchMessages, setSearchMessages] = useState('')
    return <>
        <Title order={1}>Messages</Title>
        <MessageList search={searchMessages} setSearch={setSearchMessages} appId={appId} />  
    </>
}
  
export default MessagesPage