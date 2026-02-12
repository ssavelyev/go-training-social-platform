//up
alter table users
add column is_active boolean not null default false;

//down
alter table users
drop column is_active;