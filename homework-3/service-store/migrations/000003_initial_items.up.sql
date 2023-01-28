INSERT INTO
    items (seller_id, price, item_name, creation_date)
    VALUES (1, 1000, 'IPHONE 13', '2021-10-19 10:23:54')
ON CONFLICT DO NOTHING;


INSERT INTO
    items (seller_id, price, item_name, creation_date)
VALUES (2, 950, 'SAMSUNG S15', '2021-11-14 15:01:04')
ON CONFLICT DO NOTHING;


INSERT INTO
    items (seller_id, price, item_name, creation_date)
VALUES (3, 450, 'GOOGLE PIXEL 6', '2022-01-15 16:56:35')
ON CONFLICT DO NOTHING;