CREATE TABLE IF NOT EXISTS cancel_orders (
    id bigserial PRIMARY KEY,
    order_id bigserial NOT NULL UNIQUE
        REFERENCES orders (id),
    cancellation_date timestamp NOT NULL
);