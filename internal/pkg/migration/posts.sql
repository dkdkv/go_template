create table posts (
  id serial primary key,
  title varchar(255) not null,
  content text,
  cover_url varchar(255)
);