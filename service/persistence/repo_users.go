package persistence

import (
	"database/sql"
	"fmt"
	models2 "support-bot/service/models"
)

func (r *Repository) UpsertUser(user *models2.User) (*models2.User, error) {
	const query = `insert into users (id, name, password, company_id, surname, email, phone, created_at, updated_at) 
	values (?,?,?,?,?,?,?,?,?) `
	statement, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	_, err = statement.Exec(user.ID, user.Name, user.Password, user.Company, user.Surname, user.Email, user.Phone, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("executing upsert user statement %w", err)
	}
	return user, nil
}

func (r *Repository) GetUserByEmail(email, tenant string) (*models2.User, error) {
	const query = `select * from users where email = $1 and company = $2 limit 1`
	row := r.db.QueryRow(query, email, tenant)
	var user models2.User
	if err := row.Scan(&user.ID, &user.Name, &user.Surname, &user.Password, &user.Company, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) GetUserByID(id string) (*models2.User, error) {
	const query = `select * from users where id = $1 limit 1`
	row := r.db.QueryRow(query, id)
	var user models2.User
	if err := row.Scan(&user.ID, &user.Name, &user.Surname, &user.Password, &user.Company, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) UpsertTenant(tenant *models2.Tenant) (*models2.Tenant, error) {
	const query = `insert into tenants (id, name, created_at, updated_at) values (?,?,?,?) `
	statement, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	_, err = statement.Exec(tenant.ID, tenant.Name, tenant.CreatedAt, tenant.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("executing upsert tenant statement %w", err)
	}
	return tenant, nil
}

func (r *Repository) GetTenantByID(id string) (*models2.Tenant, error) {
	const query = `select * from tenants where id = $1 limit 1`
	row := r.db.QueryRow(query, id)
	var tenant models2.Tenant
	if err := row.Scan(&tenant.ID, &tenant.Name, &tenant.CreatedAt, &tenant.UpdatedAt); err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (r *Repository) TenantHasSuperuser(tenantID string) (bool, error) {
	const query = `select id from users where company_id = $1 limit 1`
	row := r.db.QueryRow(query, tenantID)
	if row.Err() == sql.ErrNoRows {
		return true, nil
	}
	return false, row.Err()
}
