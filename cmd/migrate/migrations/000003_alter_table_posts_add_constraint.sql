//up
alter table posts
add constraint fk_posts_user
foreign key (user_id) 
references users (id) 
on delete cascade;

//down
posts drop constraint fk_posts_user;