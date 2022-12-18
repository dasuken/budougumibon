CREATE TABLE "user" (
  "id" serial PRIMARY KEY NOT NULL,
  "name" varchar(20) UNIQUE NOT NULL,
  "password" varchar(80) NOT NULL,
  "role" varchar(80) NOT NULL,
  "created" timestamptz NOT NULL,
  "modified" timestamptz NOT NULL
);

CREATE TABLE "task" (
  "id" serial PRIMARY KEY NOT NULL,
  "title" varchar(128) NOT NULL,
  "status" varchar(20) NOT NULL,
  "created" timestamptz NOT NULL,
  "modified" timestamptz NOT NULL
);
