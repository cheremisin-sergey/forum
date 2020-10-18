create table if not exists users
(
   id serial,
   username varchar (254) not null,
   about  text
   anonymous  boolean,
   name varchar (254),
   email citext UNIQUE,
   primary key (ID)
);
