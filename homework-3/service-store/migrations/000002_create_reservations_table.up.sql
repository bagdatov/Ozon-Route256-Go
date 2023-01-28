CREATE TABLE IF NOT EXISTS reservations (
    id bigserial PRIMARY KEY,
    item_id bigserial NOT NULL UNIQUE REFERENCES items(id) ON DELETE CASCADE,
    order_id bigserial NOT NULL UNIQUE,
    update_date timestamp NOT NULL
);