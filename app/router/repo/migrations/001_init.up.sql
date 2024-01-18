create schema evnt;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create type evnt.new_event_type AS (
    id text,
    schema JSONB
);
-- event_type is an internal event type that producers send
create table evnt.event_type (
    created_at timestamptz not null default now(),
    schema JSONB not null,
    id text primary key
);
create type evnt.new_provider AS (
    id text
);
-- provider is a service that sends notifications on request
create table evnt.provider (
    created_at timestamptz not null default now(),
    id text primary key
);
create type evnt.new_notification_type AS (
    provider text,
    metadata JSONB,
    id text
);
-- notification_type is an external event type that providers allow to send
create table evnt.notification_type (
    provider text not null references evnt.provider(id),
    created_at timestamptz not null default now(),
    id text primary key
);
create type evnt.new_subscription AS (
    notification_type text,
    tenant_id text,
    metadata JSONB
);
-- subscription is a mapping between the event_notifications 
-- and the targets to call on the provider when internal events happen
create table evnt.subscription (
    notification_type text not null references evnt.notification_type(id),
    uid uuid unique not null default uuid_generate_v4(),
    tenant_id text not null,
    metadata JSONB not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz, 
    primary key(tenant_id, notification_type)
);

create table evnt.event_notification(
    notification_type text not null references evnt.notification_type(id),
    event_type text not null references evnt.event_type(id),
    transform text null
);