//up
create table if not exists roles (
	id bigserial primary key,
	name varchar(255) unique not null,
	level int not null default 0
);

insert into roles (name, level)
values 
('user', 1),
('moderator', 2),
('admin', 3);

//down
drop table if exists roles;