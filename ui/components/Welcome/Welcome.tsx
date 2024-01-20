import { Title, Text, Anchor } from '@mantine/core';
import classes from './Welcome.module.css';

export function Welcome() {
  return (
    <>
      <Title className={classes.title} ta="center" mt={100}>
        Welcome to{' '}
        <Text inherit variant="gradient" component="span" gradient={{ from: 'red', to: 'blue' }}>
          Webhooks API
        </Text>
      </Title>
      <Text c="dimmed" ta="center" size="lg" maw={580} mx="auto" mt="xl">
        This is a Proof Of Concept for integration between published events and svix.
      </Text>
    </>
  );
}
