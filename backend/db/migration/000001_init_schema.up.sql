CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "full_name" varchar,
  "phone_number" varchar,
  "latitude" bigint NOT NULL,
  "longitude" varchar NOT NULL,
  "address" varchar,
  "is_family" boolean,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX users_full_name_idx ON users (full_name);