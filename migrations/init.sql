create table users (
    id uuid primary key,
    email citext unique,
    password_hash text not null,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
);

create table torrents (
    id uuid primary key,
    path text not null,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
);

create table torrents_users {
    torrent_id uuid,
    user_id uuid,
    completed boolean,
    upload bigint,
    download bigint,
    active boolean,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
    constraint user_id_feed_id_unq UNIQUE (torrent_id, user_id)
}

create or replace function update_modified_column()
returns trigger as $$
begin
    new.updated_at = now();
    return new;
end;
$$ language 'plpgsql';

create trigger update_users_modtime before update on users for each row execute procedure update_modified_column();
create trigger update_torrents_modtime BEFORE update on torrents for each row execute procedure update_modified_column();
