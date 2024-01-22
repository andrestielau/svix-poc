import { UndefinedInitialDataOptions, UseMutationOptions, UseQueryOptions, UseQueryResult, useMutation, useQuery, useQueryClient} from "@tanstack/react-query"
import { NewModal, SearchList, SearchListProps } from "@/components/components"
import { Button, Divider, Group, MenuItem, rem } from "@mantine/core"
import { UseFormInput, UseFormReturnType, useForm } from "@mantine/form"
import { ReactNode } from "react"
import { IconTrash } from "@tabler/icons-react"

export type QueryListProps<T, F> = {
    accessor: (value: T) => string 
    children: (value: T) => ReactNode
} & SearchListProps & UndefinedInitialDataOptions<F, Error, T[]>

export const QueryList = <T, F = any>({ action, children, accessor, value, setValue, ...rest }: QueryListProps<T, F>) => { 
    const query = useQuery<F, Error, T[]>(rest)
    return <SearchList value={value} setValue={setValue} action={action}>
        {query.data?.filter((data) => !value || accessor(data).includes(value)).map(children)}
    </SearchList> 
}
export type CreationModalProps<F,T> = {
    title: string
    queryKey: string[]
    onReset?: () => void
    children: (form: UseFormReturnType<F>) => ReactNode
} & UseFormInput<F> & UseMutationOptions<T, Error, F>

export const CreationModal = <F,T = any>({ children, title, queryKey, onSuccess, onReset, ...rest }: CreationModalProps<F,T>) => {
    const client = useQueryClient()
    const add = useMutation<T, Error, F>({ ...rest,
        onSuccess: (d,v,c) => client.invalidateQueries({ queryKey }).
            then(() => onSuccess && onSuccess(d,v,c))
    })
    const form = useForm<F>(rest);
    return <NewModal title={title}>
        <form onSubmit={form.onSubmit((values) => add.mutate(values))} onReset={() => { form.reset(); return onReset && onReset() }} >
            {children(form)}
            <Divider my="xs" />
            <Group justify="end">
                <Button variant="light" type='reset' disabled={!form.isDirty}>Reset</Button>
                <Button variant="outline" type='submit'>Create</Button>
            </Group>
        </form>
    </NewModal>
}

export const DeleteMenuItem = ({queryKey, onSuccess, ...rest }: UseMutationOptions & { queryKey: string[] }) => {
    const client = useQueryClient()
    const del = useMutation({
      ...rest,
      onSuccess: (d,v,c) => client.invalidateQueries({ queryKey }).
        then(() => onSuccess && onSuccess(d,v,c))
    })
    return <MenuItem
      color="red"
      leftSection={<IconTrash style={{ width: rem(14), height: rem(14) }} />}
      onClick={() => del.mutate()}
    >
      Delete
    </MenuItem>
  }