CREATE TABLE IF NOT EXISTS order_status2 (
    order_id bigserial NOT NULL,
    item_id bigserial NOT NULL,
    seller_id bigserial NOT NULL,
    client_id bigserial NOT NULL,
    update_date timestamp NOT NULL,
    is_cancelled boolean NOT NULL
);