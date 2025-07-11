create table if not exists users (
    id uuid primary key,
    email text not null unique,
    password text not null,
    personal_access_token text not null unique,
    created_at numeric not null default extract(
        epoch
        from now()
    ),
    updated_at numeric not null default extract(
        epoch
        from now()
    )
);

CREATE TABLE telegram_users (
    telegram_id BIGINT PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users (id)
);

create table if not exists cars (
    id uuid primary key,
    user_id uuid references users (id),
    latitude double precision not null,
    longitude double precision not null,
    battery integer not null,
    is_locked boolean not null,
    is_charging boolean not null,
    climate_on boolean not null,
    last_update numeric not null default extract(
        epoch
        from now()
    )
);

create table if not exists token_pairs (
    access_token text primary key,
    refresh_token text not null,
    expires_in bigint not null
);

create table if not exists log_events (
    id uuid primary key,
    user_id uuid references users (id),
    event text not null,
    timestamp double precision not null
);
