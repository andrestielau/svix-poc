import { pgTable, pgSchema, timestamp, text, foreignKey, jsonb, primaryKey, unique, uuid } from "drizzle-orm/pg-core"
  import { sql } from "drizzle-orm"


export const evnt = pgSchema("evnt");

export const provider = evnt.table("provider", {
	createdAt: timestamp("created_at", { withTimezone: true, mode: 'string' }).defaultNow().notNull(),
	id: text("id").primaryKey().notNull(),
});

export const notificationType = evnt.table("notification_type", {
	provider: text("provider").notNull().references(() => provider.id),
	createdAt: timestamp("created_at", { withTimezone: true, mode: 'string' }).defaultNow().notNull(),
	id: text("id").primaryKey().notNull(),
});

export const eventNotification = evnt.table("event_notification", {
	notificationType: text("notification_type").notNull().references(() => notificationType.id),
	eventType: text("event_type").notNull().references(() => eventType.id),
	transform: text("transform"),
});

export const eventType = evnt.table("event_type", {
	createdAt: timestamp("created_at", { withTimezone: true, mode: 'string' }).defaultNow().notNull(),
	schema: jsonb("schema").notNull(),
	id: text("id").primaryKey().notNull(),
});

export const subscription = evnt.table("subscription", {
	notificationType: text("notification_type").notNull().references(() => notificationType.id),
	uid: uuid("uid").default(sql`uuid_generate_v4()`).notNull(),
	tenantId: text("tenant_id").notNull(),
	metadata: jsonb("metadata").notNull(),
	createdAt: timestamp("created_at", { withTimezone: true, mode: 'string' }).defaultNow().notNull(),
	updatedAt: timestamp("updated_at", { withTimezone: true, mode: 'string' }),
},
(table) => {
	return {
		subscriptionPkey: primaryKey({ columns: [table.notificationType, table.tenantId], name: "subscription_pkey"}),
		subscriptionUidKey: unique("subscription_uid_key").on(table.uid),
	}
});