package persistence

import (
	"fmt"
	"support-bot/internal/models"
)

func (r *Repository) SaveMessage(msg *models.Message) error {
	const query = `insert into messages (id, chat_id, type, text, created_at, updated_at) 
	values ($1,$2,$3,$4,$5,$6)`
	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(msg.ID, msg.Chat, msg.Type, msg.Text, msg.CreatedAt, msg.UpdatedAt)
	if err != nil {
		return fmt.Errorf("saving message %w", err)
	}
	return nil
}
