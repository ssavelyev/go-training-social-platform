//up
create table if not exists user_invitations (
  token bytea primary key,
  user_id bigint not null
);

//down
drop table if exists user_invitations;