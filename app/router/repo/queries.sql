-- CreateEventTypes inserts event types into the database
-- name: CreateEventTypes :many
INSERT INTO evnt.event_type (
  name
) 
SELECT 
  name
FROM unnest(pggen.arg('event_types')::evnt.new_event_type[])
ON CONFLICT DO NOTHING
RETURNING 
  id,
  uid,
  name,
  created_at;

-- CreateNotificationTypes inserts notification types into the database
-- name: CreateNotificationTypes :many
INSERT INTO evnt.notification_type (
  name,
  provider_id
) 
SELECT 
  u.name,
  p.id
FROM unnest(pggen.arg('notification_types')::evnt.new_notification_type[]) u
JOIN evnt.provider p ON p.uid = u.provider_uid 
ON CONFLICT DO NOTHING
RETURNING 
  id,
  uid,
  name,
  created_at;

-- CreateProviders inserts providers into the database
-- name: CreateProviders :many
INSERT INTO evnt.provider (
  name
) 
SELECT 
  name
FROM unnest(pggen.arg('providers')::evnt.new_provider[])
ON CONFLICT DO NOTHING
RETURNING 
  id,
  uid,
  name,
  created_at;

-- CreateSubscriptions inserts subscriptions into the database
-- name: CreateSubscriptions :many
INSERT INTO evnt.subscription (
  tenant_id,
  metadata,
  notification_type_id
) 
SELECT 
  u.tenant_id,
  u.metadata,
  n.id
FROM unnest(pggen.arg('subscriptions')::evnt.new_subscription[]) u
JOIN evnt.notification_type n ON n.uid = u.notification_type_uid 
ON CONFLICT DO NOTHING
RETURNING 
  uid,
  tenant_id,
  metadata,
  created_at,
  updated_at,
  notification_type_id;

-- DeleteSubscription deletes subscriptions TODO - soft delete 
-- name: DeleteSubscription :exec
DELETE FROM evnt.subscription
WHERE uid = ANY(pggen.arg('ids')::uuid[]);


-- GetEventTypes fetches a batch of event types by id
-- name: GetEventTypes :many
SELECT 
  id,
  uid,
  name,
  created_at
FROM evnt.event_type
WHERE uid = ANY(pggen.arg('ids')::uuid[]);

-- GetNotificationTypes fetches a batch of event types by id
-- name: GetNotificationTypes :many
SELECT 
  id,
  uid,
  name,
  created_at
FROM evnt.notification_type
WHERE uid = ANY(pggen.arg('ids')::uuid[]);

-- GetProviders fetches a batch of providers by id
-- name: GetProviders :many
SELECT 
  id,
  uid,
  name
FROM evnt.provider
WHERE uid = ANY(pggen.arg('ids')::uuid[]);

-- GetSubscriptions fetches a batch of subscriptions by id
-- name: GetSubscriptions :many
SELECT 
  uid,
  tenant_id,
  metadata,
  created_at,
  updated_at,
  notification_type_id
FROM evnt.subscription
WHERE uid = ANY(pggen.arg('ids')::uuid[]);

-- ListEventTypes returns a page of event types
-- name: ListEventTypes :many
SELECT 
  id,
  uid,
  name,
  created_at
FROM evnt.event_type
WHERE pggen.arg('created_after')::timestamptz IS NULL 
  OR created_at > pggen.arg('created_after')
ORDER BY created_at
LIMIT pggen.arg('limit');

-- ListNotificationTypes returns a page of notification types
-- name: ListNotificationTypes :many
SELECT 
  id,
  uid,
  name,
  created_at
FROM evnt.notification_type
WHERE pggen.arg('created_after')::timestamptz IS NULL 
  OR created_at > pggen.arg('created_after')
ORDER BY created_at
LIMIT pggen.arg('limit');

-- ListProviders returns a page of providers
-- name: ListProviders :many
SELECT 
  id,
  uid,
  name,
  created_at
FROM evnt.provider
WHERE pggen.arg('created_after')::timestamptz IS NULL 
  OR created_at > pggen.arg('created_after')
ORDER BY created_at
LIMIT pggen.arg('limit');


-- ListSubscriptions returns a page of subscriptions
-- name: ListSubscriptions :many
SELECT 
  uid,
  tenant_id,
  metadata,
  created_at,
  updated_at,
  notification_type_id
FROM evnt.subscription
WHERE pggen.arg('created_after')::timestamptz IS NULL 
  OR created_at > pggen.arg('created_after')
ORDER BY created_at
LIMIT pggen.arg('limit');

-- ListEventSubscriptions returns a page of notifications for given event and tenant
-- name: ListEventSubscriptions :many
WITH filtered AS (
  SELECT 
    s.uid,
    en.transform,
    e.id  AS event_type_id,
    e.uid AS event_type_uid,
    n.uid AS notification_type_uid
  FROM evnt.event_type e
  LEFT JOIN evnt.event_notification en
    ON en.event_type_id = e.id
  LEFT JOIN evnt.notification_type n
    ON en.notification_type_id = n.id
  LEFT JOIN evnt.subscription s
    ON n.id = s.notification_type_id
  WHERE e.uid = pggen.arg('event')
    AND (pggen.arg('tenant_id') = '' 
      OR s.tenant_id = pggen.arg('tenant_id'))
) SELECT 
  s.uid,
  s.tenant_id,
  s.metadata,
  s.created_at,
  s.updated_at,
  s.notification_type_id,
  f.transform,
  f.event_type_id,
  f.event_type_uid,
  f.notification_type_uid
FROM evnt.subscription s 
INNER JOIN filtered f
  ON s.uid = f.uid
WHERE s.created_at > pggen.arg('created_after')
ORDER BY s.created_at
LIMIT pggen.arg('limit');