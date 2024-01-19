import { SearchList, WithMenu } from "@/components/components"
import { listNotificationTypes, listProviders } from "@/provider/drizzle/queries"
import { Button, Tooltip } from "@mantine/core"
import { useQuery } from "@tanstack/react-query"

export type ProviderProps = {
    search: string
    setSearch: (s: string) => void 
}

export const ProviderList = ({ search, setSearch }: ProviderProps) => { 
    const data = useQuery({queryKey: ['providers'], queryFn: async () => await listProviders()})
    return <SearchList value={search} setValue={setSearch}>
        {data.data?.map(({ id, createdAt }) => <WithMenu key={id}>
            <Tooltip label={'Created At: '+ createdAt}>
                <Button variant="default" fullWidth>{id}</Button>
            </Tooltip>
        </WithMenu>)}
    </SearchList> 
}