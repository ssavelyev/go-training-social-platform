//up
alter table posts
add column version int default 0;

//down
alter table posts
drop column version;