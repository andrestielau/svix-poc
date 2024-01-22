import { Button, TextInput,  Menu, rem, Text, Modal, Stack, Flex, CloseButton } from "@mantine/core"
import { useDisclosure } from "@mantine/hooks"
import { IconDotsVertical, IconMessageCircle, IconPhoto, IconPlus, IconSearch, IconSettings, IconTrash } from "@tabler/icons-react"
import { PropsWithChildren, ReactNode } from "react"


export const WithMenu = ({ children, id, danger }: PropsWithChildren<{ id?: string, danger?: ReactNode }>) => <Button.Group>
  {children}
  <Menu shadow="md" width={200}>
      <Menu.Target>
        <Button variant="default">
          <IconDotsVertical style={{ width: rem(20), height: rem(20) }}/>
        </Button>
      </Menu.Target>

      <Menu.Dropdown>
        <Menu.Label>Application</Menu.Label>
        <Menu.Item leftSection={<IconSettings style={{ width: rem(14), height: rem(14) }} />}>
          Settings
        </Menu.Item>
        <Menu.Item leftSection={<IconMessageCircle style={{ width: rem(14), height: rem(14) }} />}>
          Messages
        </Menu.Item>
        <Menu.Item leftSection={<IconPhoto style={{ width: rem(14), height: rem(14) }} />}>
          Gallery
        </Menu.Item>
        <Menu.Item
          leftSection={<IconSearch style={{ width: rem(14), height: rem(14) }} />}
          rightSection={
            <Text size="xs" c="dimmed">
              ⌘K
            </Text>
          }
        >
          Search
        </Menu.Item>
        <Menu.Divider />
        {danger && <Menu.Label>Danger zone</Menu.Label>}
        {danger}
      </Menu.Dropdown>
    </Menu>
</Button.Group>

export const NewModal = ({ title, children }: PropsWithChildren<{title: string}>) => {
  const [opened, { open, close }] = useDisclosure(false);
  return <>
    <Modal opened={opened} onClose={close} title={title} centered>
      {children}
    </Modal>
    <Button variant="default" size="sm" onClick={open}>
      <IconPlus style={{ width: rem(20), height: rem(20) }}/>
    </Button>
  </>
}

export type SearchListProps = { 
  value: string
  action?: ReactNode
  setValue: (v: string) => void
}
export const SearchList = ({ action, children, value, setValue }: PropsWithChildren<SearchListProps>) => {
  return <Stack h={300} bg="var(--mantine-color-body)" gap="xs" pt={10}>
      <Flex justify="flex-end" align="center" direction="row">
          <SearchBox value={value} setValue={setValue} />
          {action}
      </Flex>
      {children}
  </Stack>
}

export const SearchBox = ({ value, setValue }: any) => 
  <TextInput placeholder="Search"
      style={{width: '100%'}}
      value={value}
      onChange={(event) => setValue(event.currentTarget.value)}
      rightSectionPointerEvents="all"
      rightSection={<CloseButton 
          aria-label="Clear input" 
          onClick={() => setValue('')} 
          style={{ display: value ? undefined : 'none' }} 
      />}
  />