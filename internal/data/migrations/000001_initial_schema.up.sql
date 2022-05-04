CREATE TABLE "staff" (
  "id" SERIAL PRIMARY KEY,
  "full_name" varchar,
  "username" varchar,
  "password" varchar,
  "email" varchar,
  "created_at" timestamp,
  "role" int
);

CREATE TABLE "roles" (
  "code" int PRIMARY KEY,
  "name" varchar
);

ALTER TABLE "staff" ADD FOREIGN KEY ("role") REFERENCES "roles" ("code");
