-- Current sql file was generated after introspecting the database
-- If you want to run this migration please uncomment this code before executing migrations
/*
CREATE SCHEMA "evnt";
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "evnt"."provider" (
	"created_at" timestamp with time zone DEFAULT now() NOT NULL,
	"id" text PRIMARY KEY NOT NULL
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "evnt"."notification_type" (
	"provider" text NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL,
	"id" text PRIMARY KEY NOT NULL
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "evnt"."event_notification" (
	"notification_type" text NOT NULL,
	"event_type" text NOT NULL,
	"transform" text
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "evnt"."event_type" (
	"created_at" timestamp with time zone DEFAULT now() NOT NULL,
	"schema" jsonb NOT NULL,
	"id" text PRIMARY KEY NOT NULL
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "evnt"."subscription" (
	"notification_type" text NOT NULL,
	"uid" uuid DEFAULT uuid_generate_v4() NOT NULL,
	"tenant_id" text NOT NULL,
	"metadata" jsonb NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL,
	"updated_at" timestamp with time zone,
	CONSTRAINT "subscription_pkey" PRIMARY KEY("notification_type","tenant_id"),
	CONSTRAINT "subscription_uid_key" UNIQUE("uid")
);
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "evnt"."notification_type" ADD CONSTRAINT "notification_type_provider_fkey" FOREIGN KEY ("provider") REFERENCES "evnt"."provider"("id") ON DELETE no action ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "evnt"."event_notification" ADD CONSTRAINT "event_notification_notification_type_fkey" FOREIGN KEY ("notification_type") REFERENCES "evnt"."notification_type"("id") ON DELETE no action ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "evnt"."event_notification" ADD CONSTRAINT "event_notification_event_type_fkey" FOREIGN KEY ("event_type") REFERENCES "evnt"."event_type"("id") ON DELETE no action ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "evnt"."subscription" ADD CONSTRAINT "subscription_notification_type_fkey" FOREIGN KEY ("notification_type") REFERENCES "evnt"."notification_type"("id") ON DELETE no action ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;

*/