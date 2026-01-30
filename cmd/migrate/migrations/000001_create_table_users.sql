//up
create table if not exists users (
	id bigserial primary key,
	username varchar(255) unique not null,
	email varchar(255) unique not null,
	password bytea not null,
	created_at timestamp(0) with time zone not null default now()
);

//down
drop table if exists users;