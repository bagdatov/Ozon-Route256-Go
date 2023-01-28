CREATE TABLE IF NOT EXISTS questions (
    id serial PRIMARY KEY,
    num int2 NOT NULL,
    tour_id serial NOT NULL REFERENCES tours (id),
    question varchar(1024) NOT NULL,
    answer varchar(1024) NOT NULL,
    authors varchar(512) NOT NULL,
    "comments" varchar(1024) NOT NULL
);