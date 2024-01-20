export const toObj = <T extends {}, R extends { [K in keyof T]: T[K]}>(obj: T) => 
    JSON.parse(JSON.stringify(obj)) as R // TODO: refactor this idiocy
