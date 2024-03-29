import { defineConfig } from 'drizzle-kit'

export default defineConfig({
    schema: "./provider/drizzle/schema.ts",
    out: "./provider/drizzle",
    driver: 'pg',
    dbCredentials: {
        connectionString: "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable",
    },
    verbose: true,
    strict: true,
    schemaFilter: 'evnt'
})