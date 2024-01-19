import { drizzle } from "drizzle-orm/node-postgres";
import * as schema  from './schema'
import { Client } from "pg";

const client = new Client({
    connectionString: "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable",
});

client.connect().then(console.log);

export const db = drizzle(client, { schema });