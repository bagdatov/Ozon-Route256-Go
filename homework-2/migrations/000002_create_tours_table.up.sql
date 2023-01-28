CREATE TABLE IF NOT EXISTS tours (
    id serial PRIMARY KEY,
    text_id varchar(32) UNIQUE NOT NULL,
    tournament_id serial NOT NULL REFERENCES tournaments (id),
    title varchar(64) NOT NULL,
    editors varchar(256) NOT NULL,
    questions_num int2 NOT NULL
);