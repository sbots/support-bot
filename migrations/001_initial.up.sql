create table if not exists bots (
   id UUID primary key not null unique,
   token text unique not null,
   type text not null,
   created_at timestamp not null,
   updated_at timestamp not null
);

create table if not exists tenants
(
  id UUID primary key not null unique,
  name varchar(225) unique not null,
  created_at timestamp not null,
  updated_at timestamp not null
);

create table if not exists users
(
    id         UUID primary key    not null unique,
    email      varchar(225) unique not null,
    name       varchar(225)        not null,
    surname    varchar(225)        not null,
    password   varchar(225)        not null,
    phone      varchar(225)        not null,
    created_at timestamp           not null,
    updated_at timestamp           not null,
    company_id UUID                not null,
    foreign key (company_id) references tenants (id) ON DELETE CASCADE ON UPDATE NO ACTION
);