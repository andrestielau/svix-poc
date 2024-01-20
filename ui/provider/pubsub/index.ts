import { ISchema, Subscription, Topic } from "@google-cloud/pubsub";
import { pubsub } from "./client";
import { Buffer } from "buffer";

export const isEmulator = () => pubsub.isEmulator

export const schemas = async () => {
    const ret: ISchema[] = []
    for await (const schema of pubsub.listSchemas()) { // Node doesn't support Array.fromAsync: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/fromAsync
        ret.push(schema)
    }
    return ret
}
export const topics = () => pubsub.getTopics()
export const subscriptions = () => pubsub.getSubscriptions()

export const topic = (nameOrId: string) => pubsub.topic(nameOrId)
export const schema = (nameOrId: string) => pubsub.schema(nameOrId)
export const subscription = (nameOrId: string) => pubsub.subscription(nameOrId)

export const createTopic = (nameOrId: string) => pubsub.createTopic(nameOrId)

export const subscriptionMetadata = (subscription: Subscription) => subscription.getMetadata() 
export const deleteSubscription = (subscription: Subscription) => subscription.delete()

// TODO: Map options as types
export const subscribe = (topic: Topic, name: string) => topic.createSubscription(name)
export const publish = (topic: Topic, data: Buffer) => topic.publishMessage({data})
export const deleteTopic = (topic: Topic) => topic.delete()
