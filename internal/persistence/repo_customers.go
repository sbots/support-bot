package persistence

import (
	"database/sql"
	"fmt"
	"support-bot/internal/infrastructure/errors"
	"support-bot/internal/models"
)

func (r *Repository) UpsertCustomer(customer *models.Customer) error {
	const query = `insert into customers (id, name, surname, platform, platform_user_ud, company_id, phone, created_at, updated_at) 
	values ($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(customer.ID, customer.Name, customer.Surname, customer.Platform, customer.PlatformUserID, customer.Company, customer.Phone, customer.CreatedAt, customer.UpdatedAt)
	if err != nil {
		return fmt.Errorf("executing upsert customer statement %w", err)
	}
	return nil
}

func (r *Repository) GetCustomerByID(id string) (*models.Customer, error) {
	const query = `
		select id, name, surname, platform, platform_user_id, company_id, phone, created_at, updated_at, deleted_at
		from customers 
		where id = $1 and deleted_at is null 
		limit 1`
	row := r.db.QueryRow(query, id)
	if row.Err() == sql.ErrNoRows {
		return nil, errors.NotFound
	}
	var customer models.Customer
	if err := row.Scan(&customer.ID, &customer.Name, &customer.Surname, &customer.Platform, &customer.PlatformUserID, &customer.Company, &customer.Phone, &customer.CreatedAt, &customer.UpdatedAt, &customer.DeletedAt); err != nil {
		return nil, err
	}
	return &customer, nil
}
