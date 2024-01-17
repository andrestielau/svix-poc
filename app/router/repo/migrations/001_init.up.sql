create schema evnt;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create type evnt.new_event_type AS (
    name text,
    key text,
    schema JSONB
);
-- event_type is an internal event type that producers send
create table evnt.event_type (
    uid uuid unique not null default uuid_generate_v4(),
    created_at timestamptz not null default now(),
    name text unique not null,
    key text unique not null,
    schema JSONB not null,
    id serial primary key
);
create type evnt.new_provider AS (
    name text
);
-- provider is a service that sends notifications on request
create table evnt.provider (
    uid uuid unique not null default uuid_generate_v4(),
    created_at timestamptz not null default now(),
    name text unique not null,
    key text unique not null,
    id serial primary key
);
create type evnt.new_notification_type AS (
    provider_uid uuid,
    metadata JSONB,
    name text
);
-- notification_type is an external event type that providers allow to send
create table evnt.notification_type (
    provider_id int not null references evnt.provider(id),
    uid uuid unique not null default uuid_generate_v4(),
    created_at timestamptz not null default now(),
    name text unique not null,
    id serial primary key
);
create type evnt.new_subscription AS (
    notification_type_uid uuid,
    tenant_id text,
    metadata JSONB
);
-- subscription is a mapping between the event_notifications 
-- and the targets to call on the provider when internal events happen
create table evnt.subscription (
    notification_type_id int not null references evnt.notification_type(id),
    uid uuid unique not null default uuid_generate_v4(),
    tenant_id text not null,
    metadata JSONB not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz, 
    primary key(tenant_id, notification_type_id)
);

create table evnt.event_notification(
    notification_type_id int not null references evnt.notification_type(id),
    event_type_id int not null references evnt.event_type(id),
    transform text null
);