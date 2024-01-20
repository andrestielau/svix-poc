"use client"
import { Title } from "@mantine/core";
import { useState } from "react";
import { SchemaList } from "@/modules/pubsub/schema";


const SchemasPage = () => {
    const [searchSchemas, setSearchSchemas] = useState('')
    return <>
        <Title order={1}>Schemas</Title>
        <SchemaList search={searchSchemas} setSearch={setSearchSchemas}/>
    </>
}

export default SchemasPage