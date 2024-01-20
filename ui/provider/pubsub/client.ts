import { PubSub } from "@google-cloud/pubsub";

export const pubsub = new PubSub({projectId: 'demo', apiEndpoint: 'http://localhost:8085'});
