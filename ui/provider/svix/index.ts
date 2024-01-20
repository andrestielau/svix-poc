"use server"
import { ApplicationIn, ApplicationListOptions, EndpointIn, EndpointListOptions, EventTypeIn, EventTypeListOptions, MessageIn, MessageListOptions } from "svix";
import { svix } from "./client";

export const listEventTypes = async (opts?: EventTypeListOptions) => toObj(await svix.eventType.list(opts))?.data
export const createEventType = async (v: EventTypeIn) => toObj(await svix.eventType.create(v))
export const deleteEventType = async (id: string) => await svix.eventType.delete(id)
export const getEventType = async (id: string) => toObj(await svix.eventType.get(id))

export const listApplications = async (opts?: ApplicationListOptions) => toObj(await svix.application.list(opts))?.data
export const createApplication = async (v: ApplicationIn) => toObj(await svix.application.create(v))
export const deleteApplication = async (id: string) => await svix.application.delete(id)
export const getApplication = async (id: string) => toObj(await svix.application.get(id))

export const listEndpoints = async (appId: string, opts?: EndpointListOptions) => toObj(await svix.endpoint.list(appId, opts))?.data
export const createEndpoint = async (appId: string, v: EndpointIn) => toObj(await svix.endpoint.create(appId, v))
export const deleteEndpoint = async (appId: string, id: string) => await svix.endpoint.delete(appId, id)

export const listMessages = async (appId: string, opts?: MessageListOptions) => toObj(await svix.message.list(appId, opts))?.data
export const createMessage = async (appId: string, v: MessageIn) => toObj(await svix.message.create(appId, v))
export const deleteMessage = async (appId: string, id: string) => await svix.message.expungeContent(appId, id)

const toObj = <T extends {}, R extends { [K in keyof T]: T[K]}>(obj: T) => 
    JSON.parse(JSON.stringify(obj)) as R
