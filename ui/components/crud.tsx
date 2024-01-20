import { UndefinedInitialDataOptions, UseMutationOptions, UseQueryOptions, UseQueryResult, useMutation, useQuery, useQueryClient} from "@tanstack/react-query"
import { NewModal, SearchList, SearchListProps } from "@/components/components"
import { Button, Divider, Group } from "@mantine/core"
import { UseFormInput, UseFormReturnType, useForm } from "@mantine/form"
import { ReactNode } from "react"

export type QueryListProps<T, F> = {
    children: (value: T) => ReactNode
} & SearchListProps & UndefinedInitialDataOptions<F, Error, T[]>

export const QueryList = <T, F = any>({ action, children, value, setValue, ...rest }: QueryListProps<T, F>) => { 
    const query = useQuery<F, Error, T[]>(rest)
    return <SearchList value={value} setValue={setValue} action={action}>
        {query.data?.map(children)}
    </SearchList> 
}
export type CreationModalProps<F,T> = {
    title: string
    queryKey: string[]
    children: (form: UseFormReturnType<F>) => ReactNode
} & UseFormInput<F> & UseMutationOptions<T, Error, F>

export const CreationModal = <F,T = any>({ children, title, queryKey, onSuccess, ...rest }: CreationModalProps<F,T>) => {
    const client = useQueryClient()
    const add = useMutation<T, Error, F>({ ...rest,
        onSuccess: (d,v,c) => client.invalidateQueries({ queryKey }).
            then(() => onSuccess && onSuccess(d,v,c))
    })
    const form = useForm<F>(rest);
    return <NewModal title={title}>
        <form onSubmit={form.onSubmit((values) => add.mutate(values))} onReset={() => form.reset()} >
            {children(form)}
            <Divider my="xs" />
            <Group justify="end">
                <Button variant="light" type='reset' disabled={!form.isDirty}>Reset</Button>
                <Button variant="outline" type='submit' disabled={!form.isValid()}>Create</Button>
            </Group>
        </form>
    </NewModal>
}