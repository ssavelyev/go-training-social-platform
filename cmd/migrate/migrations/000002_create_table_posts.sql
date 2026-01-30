//up
create table if not exists posts (
  id bigserial primary key,        
  title text not null,
  content text not null,
  user_id bigint not null,           
  tags text[] default '{}',          
  created_at timestamp(0) with time zone not null default now(),
  updated_at timestamp(0) with time zone not null default now()
);

//down
drop table if exists posts;