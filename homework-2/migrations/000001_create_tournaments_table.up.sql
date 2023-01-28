CREATE TABLE IF NOT EXISTS tournaments (
    id serial PRIMARY KEY,
    text_id varchar(32) UNIQUE NOT NULL,
    title varchar(64) NOT NULL,
    tours_num int2 NOT NULL,
    questions_num int2 NOT NULL,
    created date NOT NULL
);