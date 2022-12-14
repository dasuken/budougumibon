CREATE TABLE "user" (
  "id" BIGINT PRIMARY KEY,
  "name" varchar(20) UNIQUE NOT NULL,
  "password" varchar(80) NOT NULL,
  "role" varchar(80) NOT NULL,
  "created" timestamptz NOT NULL,
  "modified" timestamptz NOT NULL
);

CREATE TABLE "task" (
  "id" BIGINT PRIMARY KEY,
  "title" varchar(128) NOT NULL,
  "status" varchar(20) NOT NULL,
  "created" timestamptz NOT NULL,
  "modified" timestamptz NOT NULL
);
