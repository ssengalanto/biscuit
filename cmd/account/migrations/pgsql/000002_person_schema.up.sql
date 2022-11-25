CREATE TABLE "person" (
    "id" uuid PRIMARY KEY NOT NULL,
    "account_id" uuid REFERENCES account(id) UNIQUE NOT NULL,
    "first_name" varchar NOT NULL,
    "last_name" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "phone" varchar NOT NULL,
    "date_of_birth" timestamp NOT NULL,
    "avatar" varchar NULL,
    "created_at" timestamp NOT NULL DEFAULT NOW(),
    "updated_at" timestamp NOT NULL DEFAULT NOW()
);

comment on column person.account_id is 'One-to-one relationship with account table.';
