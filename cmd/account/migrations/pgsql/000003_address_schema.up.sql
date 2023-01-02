DROP TABLE IF EXISTS address;

CREATE TABLE "address" (
    "id" uuid NOT NULL,
    "person_id" uuid NOT NULL,
    "street" varchar NOT NULL,
    "unit" varchar NULL,
    "city" varchar NOT NULL,
    "district" varchar NOT NULL,
    "state" varchar NOT NULL,
    "country" varchar NOT NULL,
    "postal_code" varchar NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT NOW(),
    "updated_at" timestamp NOT NULL DEFAULT NOW(),
    CONSTRAINT address_pkey PRIMARY KEY (id),
    CONSTRAINT address_person_id_fkey FOREIGN KEY(person_id) REFERENCES person(id) ON DELETE CASCADE
);

comment on column address.person_id is 'Many-to-one relationship with person table.';