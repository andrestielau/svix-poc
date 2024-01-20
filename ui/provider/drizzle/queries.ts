"use server"
import { eq } from "drizzle-orm"
import { db } from "."
import { NewEventType, NewNotificationType, NewSubscription, eventType, notificationType, subscription } from "./schema"

export const listProviders = async () => {
    return db.query.provider.findMany()
}
export const listEventTypes = async () => {
    return db.query.eventType.findMany()
}
export const addEventType = async (i: NewEventType) => {
    return db.insert(eventType).values(i)
}
export const remEventType = async (id: string) => {
    return db.delete(eventType).where(eq(eventType.id, id))
}
export const listNotificationTypes = async () => {
    return db.query.notificationType.findMany()
}
export const addNotificationType = async (i: NewNotificationType) => {
    return db.insert(notificationType).values(i)
}
export const remNotificationType = async (id: string) => {
    return db.delete(notificationType).where(eq(notificationType.id, id))
}
export const listSubscriptions = async () => {
    return db.query.subscription.findMany()
}
export const addSubscription = async (i: NewSubscription) => {
    return db.insert(subscription).values(i)
}
export const remSubscription = async (id: string) => {
    return db.delete(subscription).where(eq(subscription.uid, id))
}