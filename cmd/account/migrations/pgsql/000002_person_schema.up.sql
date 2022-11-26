DROP TABLE IF EXISTS person;

CREATE TABLE "person" (
    "id" uuid NOT NULL,
    "account_id" uuid UNIQUE NOT NULL,
    "first_name" varchar NOT NULL,
    "last_name" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "phone" varchar NOT NULL,
    "date_of_birth" timestamp NOT NULL,
    "avatar" varchar NULL,
    "created_at" timestamp NOT NULL DEFAULT NOW(),
    "updated_at" timestamp NOT NULL DEFAULT NOW(),
    CONSTRAINT person_pkey PRIMARY KEY (id),
    CONSTRAINT person_account_id_fkey FOREIGN KEY(account_id) REFERENCES account(id) ON DELETE CASCADE
);

comment on column person.account_id is 'One-to-one relationship with account table.';
