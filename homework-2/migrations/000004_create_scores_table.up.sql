CREATE TABLE IF NOT EXISTS scores (
    id serial PRIMARY KEY,
    chat_id int8 NOT NULL,
    username varchar(32) NOT NULL,
    question_id serial NOT NULL REFERENCES questions (id),
    is_correct boolean NOT NULL,
    created timestamp NOT NULL,
    UNIQUE (question_id, username)
);