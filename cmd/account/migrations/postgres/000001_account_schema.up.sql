CREATE TABLE "accounts" (
    "id" uuid PRIMARY KEY NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "password" varchar NOT NULL CHECK (char_length(password) <= 10),
    "active" boolean NOT NULL,
    "last_login_at" timestamp NULL,
    "created_at" timestamp NOT NULL DEFAULT NOW(),
    "updated_at" timestamp NOT NULL DEFAULT NOW()
);
