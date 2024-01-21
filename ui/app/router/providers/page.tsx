"use client"
import { Title } from "@mantine/core";
import { ProviderList } from "@/modules/router/providers";
import { useState } from "react";

const ProvidersPage = () => {
    const [searchProviders, setSearchProviders] = useState('')
    return <>
        <Title order={1}>Providers</Title>
        <ProviderList search={searchProviders} setSearch={setSearchProviders}/>
    </>
}
export default ProvidersPage