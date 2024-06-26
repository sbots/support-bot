package persistence

import (
	"database/sql"
	"fmt"
	"support-bot/internal/infrastructure/errors"
	"support-bot/internal/models"
)

func (r *Repository) UpsertTenant(tenant *models.Tenant) error {
	const query = `insert into tenants (id, name, created_at, updated_at) values ($1,$2,$3,$4)`
	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(tenant.ID, tenant.Name, tenant.CreatedAt, tenant.UpdatedAt)
	if err != nil {
		return fmt.Errorf("executing upsert tenant statement %w", err)
	}
	return nil
}

func (r *Repository) GetTenantByID(id string) (*models.Tenant, error) {
	const query = `select * from tenants where id = $1 and deleted_at is null limit 1`
	row := r.db.QueryRow(query, id)
	if row.Err() == sql.ErrNoRows {
		return nil, errors.NotFound
	}
	var tenant models.Tenant
	if err := row.Scan(&tenant.ID, &tenant.Name, &tenant.CreatedAt, &tenant.UpdatedAt, &tenant.DeletedAt); err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (r *Repository) GetTenantByName(name string) (*models.Tenant, error) {
	const query = `select * from tenants where name = $1 and deleted_at is null limit 1`
	row := r.db.QueryRow(query, name)
	if row.Err() == sql.ErrNoRows {
		return nil, errors.NotFound
	}
	var tenant models.Tenant
	if err := row.Scan(&tenant.ID, &tenant.Name, &tenant.CreatedAt, &tenant.UpdatedAt, &tenant.DeletedAt); err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (r *Repository) TenantHasSuperuser(tenantID string) (bool, error) {
	const query = `select id from users where company_id = $1 and deleted_at is null limit 1`
	row := r.db.QueryRow(query, tenantID)
	if row.Err() == sql.ErrNoRows {
		return false, nil
	}
	return false, row.Err()
}
