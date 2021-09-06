alter table bots
    drop column if exists deleted_at;

alter table tenants
    drop column if exists deleted_at;

alter table users
    drop column if exists deleted_at;
