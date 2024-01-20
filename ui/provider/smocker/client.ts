import axios from "axios";
export type Version = {
    app_name: string
    build_version: string
    build_commit: string
    build_date: string
}
export type Message = { message: string }

export type State = {
    id: string
    times_count: number
    creation_date: string
    locked?: boolean
}
export type Context = {
    times?: number
}
export type RequestDef = {
    path: string
    method: string
}
export type ResponseDef = {
    status: number
    body?: string
    delay?: Delay
    proxy?: Proxy
    headers?: Map<string, string>
}
export type Mock = {
  request:  RequestDef
  response?: ResponseDef
  dynamic_response?: DynamicResponseDef
  context: Context
  state: State
}
export type Delay = string | {
    min: string
    max: string
}
export type Proxy = {
    host: any
    follow_redirect: any
    skip_verify_tls: any
    keep_host: any
    headers: Map<string, string>
}
export type DynamicResponseDef = {
    engine: string
    script: string
}
export type NewMock = {
  request: RequestDef
  response?: ResponseDef
  dynamic_response?: DynamicResponseDef
}

export type NewMockParams = {
    reset?: boolean
    session?: string
}
export type NewSession = { name?: string }
export type Session = { id: string, name: string }
export type SessionSummary = Session & { date: string }
export type SessionDetails = SessionSummary & { history: any[] }

export type Request = { path: string, method: string, body: string, headers: Map<string, string>, date: string }
export type Body = { message: string, nearest: Mock[], request: Request }
export type Response = { status: number, body: Body, headers: Map<string, string>, date: string }
export type Failure = { request: Request, response: Response }
export type HistoryParams = { session?: string, filter?: string }
export type History = { verified: boolean, message: string, failures: Failure[] }
export type HistorySummaryParams = { session?: string, src?: string, dest?: string }
export type HistorySummary = { type: 'request' | 'response', message: string, from: string, to: string, date: string }

export type Verification = {
    mocks: {
        verified: boolean
        all_used: boolean
        message: string
        failures: Mock[]
        unused: Mock[]
      }
    history: History
  }
export const smocker = axios.create({ baseURL: 'http://localhost:8081' })