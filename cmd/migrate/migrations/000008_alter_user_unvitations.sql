//up
alter table user_invitations
add column expiry timestamp(0) with time zone not null;

//down
alter table user_invitations
drop column if exists expiry;
