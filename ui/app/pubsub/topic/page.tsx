"use client"
import { Title } from "@mantine/core";
import { useState } from "react";
import { TopicList } from "@/modules/pubsub/topics";


const TopicsPage = () => {
    const [searchTopics, setSearchTopics] = useState('')
    return <>
        <Title order={1}>Topics</Title>
        <TopicList search={searchTopics} setSearch={setSearchTopics}/>
    </>
}

export default TopicsPage