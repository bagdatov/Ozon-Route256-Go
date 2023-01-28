CREATE TABLE IF NOT EXISTS reservations (
    id bigserial PRIMARY KEY,
    order_id bigserial NOT NULL UNIQUE,
    update_date timestamp NOT NULL
);