import { drizzle } from "drizzle-orm/node-postgres";
import * as schema  from './schema'
import { Client } from "pg";

const client = new Client({
    connectionString: "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable",
});

client.connect().then(console.log);

export const db = drizzle(client, { schema });

export type Provider = typeof schema.provider.$inferSelect
export type NotificationType = typeof schema.notificationType.$inferSelect
export type NewNotificationType = typeof schema.notificationType.$inferInsert
export type EventType = typeof schema.eventType.$inferSelect
export type NewEventType = typeof schema.eventType.$inferInsert
export type Subscription = typeof schema.subscription.$inferSelect
export type NewSubscription = typeof schema.subscription.$inferInsert