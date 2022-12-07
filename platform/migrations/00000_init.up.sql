CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "users" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "email" VARCHAR UNIQUE NOT NULL,
  "password" VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS "todos" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "user_id" uuid NOT NULL,
  "title" VARCHAR(30) NOT NULL,
  "completed" BOOLEAN NOT NULL DEFAULT false
);

ALTER TABLE
  "todos"
ADD
  FOREIGN KEY ("user_id") REFERENCES "users" ("id")