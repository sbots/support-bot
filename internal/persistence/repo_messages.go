package persistence

import (
	"fmt"
	"support-bot/internal/models"
)

func (r *Repository) SaveMessage(msg *models.Message) error {
	const query = `insert into messages (id, company_id, user_id, customer_id, text, created_at, updated_at) 
	values ($1,$2,$3,$4,$5,$6,$7)`
	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(msg.ID, msg.Company, msg.User, msg.Customer, msg.Text, msg.CreatedAt, msg.UpdatedAt)
	if err != nil {
		return fmt.Errorf("saving message %w", err)
	}
	return nil
}
