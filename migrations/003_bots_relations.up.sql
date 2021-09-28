alter table bots
    add column company_id uuid;

alter table bots
    ADD FOREIGN KEY (company_id) REFERENCES tenants(id);
