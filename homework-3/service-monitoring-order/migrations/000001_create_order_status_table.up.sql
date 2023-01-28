CREATE TABLE IF NOT EXISTS order_status (
    order_id bigserial NOT NULL,
    item_id bigserial NOT NULL,
    seller_id bigserial NOT NULL,
    client_id bigserial NOT NULL,
    update_date timestamp NOT NULL,
    is_cancelled boolean NOT NULL
) PARTITION BY HASH (order_id);

CREATE EXTENSION postgres_fdw;

CREATE SERVER shard1 FOREIGN DATA WRAPPER postgres_fdw OPTIONS (host '172.20.0.6', port '5432', dbname 'postgres');
CREATE SERVER shard2 FOREIGN DATA WRAPPER postgres_fdw OPTIONS (host '172.20.0.7', port '5432', dbname 'postgres');

CREATE FOREIGN TABLE order_status1 PARTITION OF order_status FOR VALUES WITH (MODULUS 2, REMAINDER 0) SERVER shard1;
CREATE FOREIGN TABLE order_status2 PARTITION OF order_status FOR VALUES WITH (MODULUS 2, REMAINDER 1) SERVER shard2;

CREATE USER MAPPING FOR postgres server shard1 options (user 'postgres', password 'postgres');
CREATE USER MAPPING FOR postgres server shard2 options (user 'postgres', password 'postgres');

