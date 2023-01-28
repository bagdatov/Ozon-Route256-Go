CREATE TABLE IF NOT EXISTS items (
    id bigserial PRIMARY KEY,
    seller_id bigserial NOT NULL,
    price bigserial NOT NULL,
    item_name text NOT NULL,
    creation_date timestamp NOT NULL
);