"use server"
import { ApplicationIn, ApplicationListOptions, EventTypeIn, EventTypeListOptions } from "svix";
import { svix } from "./client";

export const listApplications = async (opts?: ApplicationListOptions) => toObj(await svix.application.list(opts))
export const createApplication = async (v: ApplicationIn) => toObj(await svix.application.create(v))
export const deleteApplication = async (id: string) => await svix.application.delete(id)
export const listEventTypes = async (opts?: EventTypeListOptions) => toObj(await svix.eventType.list(opts))
export const createEventType = async (v: EventTypeIn) => toObj(await svix.eventType.create(v))
export const deleteEventType = async (id: string) => await svix.eventType.delete(id)

const toObj = <T extends {}, R extends { [K in keyof T]: T[K]}>(obj: T) => 
    JSON.parse(JSON.stringify(obj)) as R
