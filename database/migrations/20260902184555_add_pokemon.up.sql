CREATE TABLE pokemon (
    name text NOT NULL,
    official_artwork text NOT NULL,
    id integer PRIMARY KEY,
    primary_type text NOT NULL,
    secondary_type text,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp
);