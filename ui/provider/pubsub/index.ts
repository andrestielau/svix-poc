"use server"
import { ISchema, TopicMetadata } from "@google-cloud/pubsub";
import { pubsub } from "./client";

export const isEmulator = async () => pubsub.isEmulator
const mapMetadata = (t: { metadata?: TopicMetadata }) => t?.metadata

export type SchemaView = "SCHEMA_VIEW_UNSPECIFIED" | "BASIC" | "FULL"
export type SchemaType = "TYPE_UNSPECIFIED" | "PROTOCOL_BUFFER" | "AVRO"

export const validateSchema = async (definition: string, type: SchemaType = "TYPE_UNSPECIFIED") => 
    pubsub.validateSchema({ type, definition }).then(() => true).catch((reason) => reason)
export const schema = async (nameOrId: string, view: SchemaView = "SCHEMA_VIEW_UNSPECIFIED") => pubsub.schema(nameOrId)?.get(view)
export const createSchema = async (id: string, def: string, type: SchemaType = "TYPE_UNSPECIFIED") => 
    pubsub.createSchema(id, type, def).then(schema => schema.get("SCHEMA_VIEW_UNSPECIFIED"))
export const schemas = async () => {
    const ret: ISchema[] = []
    for await (const schema of pubsub.listSchemas()) { // Node doesn't support Array.fromAsync: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/fromAsync
        ret.push(schema)
    }
    return ret
}

export const topic = async (nameOrId: string) => pubsub.topic(nameOrId).getMetadata().then(([topic]) => topic)
export const createTopic = async (nameOrId: string) => pubsub.createTopic(nameOrId).then(([topic]) => topic)
export const topics = async () => pubsub.getTopics().then(([topics]) => topics.map(mapMetadata))

export const subscription = (nameOrId: string) => pubsub.subscription(nameOrId)?.getMetadata().then(([s]) => s)
export const createSubscription = (topic: string, name: string) => pubsub.createSubscription(topic, name).then(([s]) => s)
export const subscriptions = async () => pubsub.getSubscriptions().then(([subscriptions]) => subscriptions.map(mapMetadata))

