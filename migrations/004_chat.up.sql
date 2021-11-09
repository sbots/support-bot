create table if not exists customers
(
    id         UUID primary key    not null unique,
    platform   varchar(50)  not null,
    platform_user_id varchar(50)  not null,
    name       varchar(225)        not null,
    surname    varchar(225)        not null,
    phone      varchar(225)        not null,
    created_at timestamp           not null,
    updated_at timestamp           not null,
    deleted_at timestamp           not null,
    company_id UUID                not null,
    foreign key (company_id) references tenants (id) ON DELETE CASCADE ON UPDATE NO ACTION
);

create table if not exists messages
(
    id UUID primary key    not null unique,
    user_id UUID not null,
    customer_id UUID not null,
    company_id UUID                not null,
    created_at timestamp           not null,
    deleted_at timestamp           not null,
    text varchar(225) not null,
    foreign key (user_id) references users (id) ON DELETE CASCADE ON UPDATE NO ACTION,
    foreign key (customer_id) references customers (id) ON DELETE CASCADE ON UPDATE NO ACTION,
    foreign key (company_id) references tenants (id) ON DELETE CASCADE ON UPDATE NO ACTION
)