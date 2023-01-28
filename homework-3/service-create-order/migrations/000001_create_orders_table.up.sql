CREATE TABLE IF NOT EXISTS orders (
    id bigserial PRIMARY KEY,
    item_id bigserial NOT NULL,
    seller_id bigserial NOT NULL,
    client_id bigserial NOT NULL,
    creation_date timestamp NOT NULL
);