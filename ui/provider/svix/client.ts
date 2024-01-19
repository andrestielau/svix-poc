import { Svix } from "svix";

export const svix = new Svix(process.env.SVIX_AUTH_TOKEN!, {
    serverUrl: 'http://localhost:8071'
})