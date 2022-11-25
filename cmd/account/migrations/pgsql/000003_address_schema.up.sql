CREATE TABLE "address" (
    "id" uuid PRIMARY KEY NOT NULL,
    "person_id" uuid REFERENCES person(id) UNIQUE NOT NULL,
    "place_id" varchar NOT NULL,
    "address_line1" json NOT NULL,
    "address_line2" json NULL,
    "city" json NOT NULL,
    "state" json NOT NULL,
    "country" json NOT NULL,
    "postal_code" json NOT NULL,
    "formatted_address" varchar NOT NULL,
    "lat" decimal NOT NULL,
    "lng" decimal NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT NOW(),
    "updated_at" timestamp NOT NULL DEFAULT NOW()
);

comment on column address.person_id is 'Many-to-one relationship with person table.';
comment on column address.place_id is 'Place id originated from Google Geolocation API.';
comment on column address.address_line1 is 'JSON column that contains short_name and long_name fields.';
comment on column address.address_line2 is 'JSON column that contains short_name and long_name fields.';
comment on column address.city is 'JSON column that contains short_name and long_name fields.';
comment on column address.state is 'JSON column that contains short_name and long_name fields.';
comment on column address.country is 'JSON column that contains short_name and long_name fields.';
comment on column address.postal_code is 'JSON column that contains short_name and long_name fields.';
