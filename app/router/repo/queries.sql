-- CreateEventTypes inserts event types into the database
-- name: CreateEventTypes :many
INSERT INTO evnt.event_type (
  id
) 
SELECT 
  id
FROM unnest(pggen.arg('event_types')::evnt.new_event_type[])
ON CONFLICT DO NOTHING
RETURNING 
  id,
  created_at;

-- CreateNotificationTypes inserts notification types into the database
-- name: CreateNotificationTypes :many
INSERT INTO evnt.notification_type (
  id,
  provider
) 
SELECT 
  id,
  provider
FROM unnest(pggen.arg('notification_types')::evnt.new_notification_type[])
ON CONFLICT DO NOTHING
RETURNING 
  id,
  created_at;

-- CreateSubscriptions inserts subscriptions into the database
-- name: CreateSubscriptions :many
INSERT INTO evnt.subscription (
  tenant_id,
  metadata,
  notification_type
) 
SELECT 
  tenant_id,
  metadata,
  notification_type
FROM unnest(pggen.arg('subscriptions')::evnt.new_subscription[]) u
ON CONFLICT DO NOTHING
RETURNING 
  uid,
  tenant_id,
  metadata,
  created_at,
  updated_at,
  notification_type;

-- DeleteSubscription deletes subscriptions TODO - soft delete 
-- name: DeleteSubscription :exec
DELETE FROM evnt.subscription
WHERE uid = ANY(pggen.arg('ids')::uuid[]);

-- DeleteEventTypes deletes eventTypes TODO - soft delete 
-- name: DeleteEventTypes :exec
DELETE FROM evnt.event_type
WHERE id = ANY(pggen.arg('ids')::text[]);

-- DeleteNotificationTypes deletes notificationTypes TODO - soft delete 
-- name: DeleteNotificationTypes :exec
DELETE FROM evnt.notification_type
WHERE id = ANY(pggen.arg('ids')::text[]);


-- GetEventTypes fetches a batch of event types by id
-- name: GetEventTypes :many
SELECT 
  id,
  created_at
FROM evnt.event_type
WHERE id = ANY(pggen.arg('ids')::text[]);

-- GetNotificationTypes fetches a batch of event types by id
-- name: GetNotificationTypes :many
SELECT 
  id,
  created_at
FROM evnt.notification_type
WHERE id = ANY(pggen.arg('ids')::text[]);

-- GetProviders fetches a batch of providers by id
-- name: GetProviders :many
SELECT 
  id
FROM evnt.provider
WHERE id = ANY(pggen.arg('ids')::text[]);

-- GetSubscriptions fetches a batch of subscriptions by id
-- name: GetSubscriptions :many
SELECT 
  uid,
  tenant_id,
  metadata,
  created_at,
  updated_at,
  notification_type
FROM evnt.subscription
WHERE uid = ANY(pggen.arg('ids')::uuid[]);

-- ListEventTypes returns a page of event types
-- name: ListEventTypes :many
SELECT 
  id,
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
  notification_type
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
    e.id AS event_type
  FROM evnt.event_type e
  LEFT JOIN evnt.event_notification en
    ON en.event_type = e.id
  LEFT JOIN evnt.notification_type n
    ON en.notification_type = n.id
  LEFT JOIN evnt.subscription s
    ON n.id = s.notification_type
  WHERE e.id = pggen.arg('event')
    AND (pggen.arg('tenant_id') = '' 
      OR s.tenant_id = pggen.arg('tenant_id'))
) SELECT 
  s.uid,
  s.tenant_id,
  s.metadata,
  s.created_at,
  s.updated_at,
  s.notification_type,
  f.transform,
  f.event_type
FROM evnt.subscription s 
INNER JOIN filtered f
  ON s.uid = f.uid
WHERE s.created_at > pggen.arg('created_after')
ORDER BY s.created_at
LIMIT pggen.arg('limit');