CREATE TABLE IF NOT EXISTS MOVIES(
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    year INTEGER NOT NULL,
    runtime INTEGER NOT NULL,
    genres text [] NOT NULL,
    version integer NOT NULL DEFAULT 1
)