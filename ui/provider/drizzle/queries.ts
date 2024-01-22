"use server"
import { eq } from "drizzle-orm"
import { NewEventType, NewNotificationType, NewSubscription, db } from "."
import { eventType, notificationType, subscription } from "./schema"

export const listProviders = async () => db.query.provider.findMany()
export const listEventTypes = async () => db.query.eventType.findMany()
export const addEventType = async (i: NewEventType) => db.insert(eventType).values(i).then(_ => true)
export const remEventType = async (id: string) =>  db.delete(eventType).where(eq(eventType.id, id)).then(_ => true)
export const listNotificationTypes = async () => db.query.notificationType.findMany()
export const addNotificationType = async (i: NewNotificationType) => db.insert(notificationType).values(i).then(_ => true)
export const remNotificationType = async (id: string) => db.delete(notificationType).where(eq(notificationType.id, id)).then(_ => true)
export const listSubscriptions = async () => db.query.subscription.findMany()
export const addSubscription = async (i: NewSubscription) => db.insert(subscription).values(i).then(_ => true)
export const remSubscription = async (id: string) => db.delete(subscription).where(eq(subscription.uid, id)).then(_ => true)
