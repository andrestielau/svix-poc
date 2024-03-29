"use server"
import { ISchema, TopicMetadata } from "@google-cloud/pubsub";
import { pubsub } from "./client";
import { MessageOptions } from "@google-cloud/pubsub/build/src/topic";

export const status = async () => ({
    isIdResolved: pubsub.isIdResolved,
    isEmulator: pubsub.isEmulator,
    isOpen: pubsub.isOpen,
})

export const publish = async (topic: string, msg: MessageOptions) => {
    if (msg.json) {
        msg.data = Buffer.from(msg.json)
        delete msg.json
    }
    return pubsub.topic(topic).publishMessage(msg)
}
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
export const deleteSchema = async (nameOrId: string) => pubsub.schema(nameOrId).delete()

export const topic = async (nameOrId: string) => pubsub.topic(nameOrId).getMetadata().then(([t]) => t)
export const createTopic = async (nameOrId: string) => pubsub.createTopic(nameOrId).then(([t]) => t)
export const topics = async () => pubsub.getTopics().then(([topics]) => topics.map(mapMetadata))
export const deleteTopic = async (nameOrId: string) => pubsub.topic(nameOrId).delete()

export const subscription = (nameOrId: string) => pubsub.subscription(nameOrId)?.getMetadata().then(([s]) => s)
export const createSubscription = (topic: string, name: string) => pubsub.createSubscription(topic, name).then(([s]) => s)
export const subscriptions = async () => pubsub.getSubscriptions().then(([s]) => s.map(mapMetadata))
export const deleteSubscription = async (nameOrId: string) => pubsub.subscription(nameOrId).delete()


const mapMetadata = (t: { metadata?: TopicMetadata }) => t?.metadata