create schema if not exists donation;

create table if not exists donation.campaign
(
    id          serial primary key,
    version     int          not null default 1,
    title       varchar(255) not null,
    description text,
    created_at  timestamptz default current_timestamp,
    updated_at  timestamptz default current_timestamp
    )
;