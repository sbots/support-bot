package persistence

import (
	"database/sql"
	"fmt"
	"support-bot/internal/infrastructure/errors"
	"support-bot/internal/models"
)

func (r *Repository) UpsertUser(user *models.User) error {
	const query = `insert into users (id, name, password, company_id, surname, email, phone, created_at, updated_at) 
	values ($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(user.ID, user.Name, user.Password, user.Company, user.Surname, user.Email, user.Phone, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("executing upsert user statement %w", err)
	}
	return nil
}

func (r *Repository) GetUserByEmail(email string) (*models.User, error) {
	const query = `
		select id, name, surname, password, company_id, email, phone, created_at, updated_at, deleted_at 
		from users 
		where email = $1 and deleted_at is null 
		limit 1`
	row := r.db.QueryRow(query, email)
	if row.Err() == sql.ErrNoRows {
		return nil, errors.NotFound
	}
	var user models.User
	if err := row.Scan(&user.ID, &user.Name, &user.Surname, &user.Password, &user.Company, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) GetUserByID(id string) (*models.User, error) {
	const query = `
		select id, name, surname, password, company_id, email, phone, created_at, updated_at, deleted_at 
		from users 
		where id = $1 and deleted_at is null 
		limit 1`
	row := r.db.QueryRow(query, id)
	if row.Err() == sql.ErrNoRows {
		return nil, errors.NotFound
	}
	var user models.User
	if err := row.Scan(&user.ID, &user.Name, &user.Surname, &user.Password, &user.Company, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err != nil {
		return nil, err
	}
	return &user, nil
}
