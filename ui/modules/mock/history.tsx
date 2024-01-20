import { SearchList, WithMenu } from "@/components/components"
import { historySummary } from "@/provider/smocker"
import { Button, Tooltip } from "@mantine/core"
import { useQuery } from "@tanstack/react-query"

export type HistoryProps = {
    search: string
    setSearch: (s: string) => void 
}

export const HistoryList = ({ search, setSearch }: HistoryProps) => { 
    const data = useQuery({queryKey: ['providers'], queryFn: async () => await historySummary()})
    return <SearchList value={search} setValue={setSearch}>
        {data.data?.map(({ date, from, to, type, message }) => <WithMenu key={from+date}>
            <Tooltip label={<span>{message}<br/>Created At: {date}</span>}>
                <Button variant="default" fullWidth>{type} from {from} to {to}</Button>
            </Tooltip>
        </WithMenu>)}
    </SearchList> 
}