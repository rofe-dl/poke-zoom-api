-- +migrate Up
CREATE TABLE IF NOT EXISTS games (
    id                     UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    score                  INT NOT NULL DEFAULT 0,
    start_time             TIMESTAMPTZ NOT NULL,
    end_time               TIMESTAMPTZ,
    max_duration           INT NOT NULL,
    current_crop_id        UUID,
    current_crop_difficulty TEXT NOT NULL DEFAULT 'hard',
    created_at             TIMESTAMPTZ NOT NULL,
    updated_at             TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS pokemon_crops (
    id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    pokemon_id        INT NOT NULL,
    game_id           UUID NOT NULL,
    easy_crop  TEXT NOT NULL,
    medium_crop TEXT NOT NULL,
    hard_crop  TEXT NOT NULL,
    is_answered       BOOLEAN NOT NULL DEFAULT FALSE,
    is_correct        BOOLEAN,
    time_taken        INT,
    tries_taken       INT NOT NULL DEFAULT 0,
    created_at        TIMESTAMPTZ NOT NULL,
    updated_at        TIMESTAMPTZ NOT NULL,
    CONSTRAINT fk_pokemon_crops_pokemon FOREIGN KEY (pokemon_id) REFERENCES pokemon(id),
    CONSTRAINT fk_pokemon_crops_game FOREIGN KEY (game_id) REFERENCES games(id)
);

CREATE INDEX IF NOT EXISTS idx_pokemon_crops_game_id ON pokemon_crops (game_id);
