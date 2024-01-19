"use server"
import { smocker, Mock, Version, Message, NewMock, HistoryParams, HistorySummaryParams, SessionSummary, SessionDetails, Session, HistorySummary, NewMockParams, NewSession } from "./client"

export const version = () => smocker.get<Version>('/version').then(({data}) => data)
export const reset = () => smocker.post<Message>('/reset').then(({data}) => data)
export const mock = (id: string) => smocker.get<Mock[]>('/mocks', { params: { id }}).then(({data}) => data ? data[0] : data)
export const mocks = (session?: string) => smocker.get<Mock[]>('/mocks', { params: { session }}).then(({data}) => data)
export const lock = (ids: string[]) => smocker.post<Mock[]>('/mocks/lock', ids).then(({data}) => data)
export const unlock = (ids: string[]) => smocker.post<Mock[]>('/mocks/unlock', ids).then(({data}) => data)
export const create = (mocks: NewMock[], params?: NewMockParams) => smocker.post<NewMock>('/mocks', mocks, {params}).then(({data}) => data)
export const history = (params: HistoryParams) => smocker.get<History[]>('/history', { params }).then(({data}) => data)
export const historySummary = (params: HistorySummaryParams) => smocker.get<HistorySummary>('/history/summary').then(({data}) => data)
export const sessions = () => smocker.get<SessionDetails[]>('/sessions').then(({data}) => data)
export const sessionsSummary = () => smocker.get<SessionSummary[]>('/sessions/summary').then(({data}) => data)
export const startSession = (params: NewSession) => smocker.post<Session>('/sessions', undefined, {params}).then(({data}) => data)
export const updateSession = (session: Session) => smocker.put<SessionSummary>('/sessions', session).then(({data}) => data)
export const verifySession = (session?: string) => smocker.post('/sessions/verify', undefined, { params: { session } }).then(({data}) => data)