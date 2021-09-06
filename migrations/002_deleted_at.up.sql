alter table bots
    add column deleted_at timestamp;

alter table tenants
    add column deleted_at timestamp;

alter table users
    add column deleted_at timestamp;