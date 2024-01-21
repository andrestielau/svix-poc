import { SearchList, WithMenu } from "@/components/components"
import { listProviders } from "@/provider/drizzle/queries"
import { useQuery } from "@tanstack/react-query"
import { Button, Tooltip } from "@mantine/core"
import { useRouter } from "next/navigation"

export type ProviderProps = {
    search: string
    setSearch: (s: string) => void 
}

export const ProviderList = ({ search, setSearch }: ProviderProps) => { 
    const data = useQuery({queryKey: ['providers'], queryFn: async () => await listProviders()})
    const router = useRouter()
    return <SearchList value={search} setValue={setSearch}>
        {data.data?.map(({ id, createdAt }) => <WithMenu key={id}>
            <Tooltip label={'Created At: '+ createdAt}>
                <Button onDoubleClick={() => router.push('/router/providers/'+id)} variant="default" fullWidth>{id}</Button>
            </Tooltip>
        </WithMenu>)}
    </SearchList> 
}