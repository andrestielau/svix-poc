import type { Config } from "drizzle-kit";
export default {
  schema: "./src/drizzle/schema.ts",
  out: "./src/drizzle",
  driver: 'pg',
  dbCredentials: {
    connectionString: "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable",
  },
  schemaFilter: ['evnt']
} satisfies Config;