import { db } from '@/drizzle'

const Home = async () => {
  const providers = await db.query.provider.findMany()
  const eventTypes = await db.query.eventType.findMany()
  const notificationTypes = await db.query.notificationType.findMany()
  const subscriptions = await db.query.subscription.findMany()
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <h1>Providers: </h1>
      <ul>
        {providers.map(({id, createdAt}) => <li>{id}: {createdAt}</li>)}
      </ul>
      <h1>EventTypes: </h1>
      <ul>
        {eventTypes.map(({id, createdAt}) => <li>{id}: {createdAt}</li>)}
      </ul>
      <h1>NotificationTypes: </h1>
      <ul>
        {notificationTypes.map(({id, createdAt}) => <li>{id}: {createdAt}</li>)}
      </ul>
      <h1>Subscriptions: </h1>
      <ul>
        {subscriptions.map(({uid, tenantId, createdAt, notificationType}) => <li>
          {tenantId}-{uid}: {createdAt} for {notificationType}
        </li>)}
      </ul>
    </main>
  )
}
export default Home