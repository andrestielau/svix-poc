"use server"
import { ISchema, TopicMetadata } from "@google-cloud/pubsub";
import { pubsub } from "./client";

export const isEmulator = async () => pubsub.isEmulator

export const schemas = async () => {
    const ret: ISchema[] = []
    for await (const schema of pubsub.listSchemas()) { // Node doesn't support Array.fromAsync: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/fromAsync
        ret.push(schema)
    }
    return ret
}
const mapMetadata = (t: { metadata?: TopicMetadata }) => t?.metadata
export const topics = async () => pubsub.getTopics().then(([topics]) => topics.map(mapMetadata))
export const subscriptions = async () => pubsub.getSubscriptions().then(([subscriptions]) => subscriptions.map(mapMetadata))

export const topic = async (nameOrId: string) => pubsub.topic(nameOrId)?.metadata
export type SchemaView = "SCHEMA_VIEW_UNSPECIFIED" | "BASIC" | "FULL"
export const schema = async (nameOrId: string, view: SchemaView = "SCHEMA_VIEW_UNSPECIFIED") => pubsub.schema(nameOrId)?.get(view)
export const subscription = (nameOrId: string) => pubsub.subscription(nameOrId)?.metadata
export type SchemaType = "TYPE_UNSPECIFIED" | "PROTOCOL_BUFFER" | "AVRO"
export const createSchema = async (id: string, def: string, type: SchemaType = "TYPE_UNSPECIFIED") => pubsub.createSchema(id, type, def).then(schema => schema.get("SCHEMA_VIEW_UNSPECIFIED"))
export const createTopic = async (nameOrId: string) => pubsub.createTopic(nameOrId).then(_ => _)
