CREATE TABLE IF NOT EXISTS sessions (
    id serial PRIMARY KEY,
    chat_id int8 NOT NULL,
    tournament_id serial NOT NULL,
    is_active boolean NOT NULL,
    created timestamp NOT NULL,
    UNIQUE (chat_id, tournament_id)
);