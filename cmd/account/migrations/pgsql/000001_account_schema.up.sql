DROP TABLE IF EXISTS account;

CREATE TABLE "account" (
    "id" uuid NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "password" varchar NOT NULL,
    "active" boolean NOT NULL,
    "last_login_at" timestamp NULL,
    "created_at" timestamp NOT NULL DEFAULT NOW(),
    "updated_at" timestamp NOT NULL DEFAULT NOW(),
    CONSTRAINT account_pkey PRIMARY KEY (id),
    CONSTRAINT account_password_ck CHECK (char_length(password) >= 10)
);
