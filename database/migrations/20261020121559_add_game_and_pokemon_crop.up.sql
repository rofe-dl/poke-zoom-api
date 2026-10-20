CREATE TABLE games (
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    score integer NOT NULL DEFAULT 0,
    start_time timestamp NOT NULL,
    end_time timestamp,
    max_duration integer NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);

CREATE TABLE pokemon_crops (
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    pokemon_id integer NOT NULL REFERENCES pokemon (id) ON UPDATE CASCADE ON DELETE CASCADE,
    game_id uuid NOT NULL REFERENCES games (id) ON UPDATE CASCADE ON DELETE CASCADE,
    x1 integer NOT NULL,
    y1 integer NOT NULL,
    x2 integer NOT NULL,
    y2 integer NOT NULL,
    x3 integer NOT NULL,
    y3 integer NOT NULL,
    is_answered boolean NOT NULL DEFAULT false,
    is_correct boolean,
    time_taken integer,
    tries_taken integer NOT NULL DEFAULT 0,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);

CREATE INDEX idx_pokemon_crops_game_id ON pokemon_crops (game_id);
