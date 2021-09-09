package persistence

import (
	"fmt"
	"support-bot/models"
)

func (r *Repository) UpsertBot(bot *models.Bot) error {
	const query = `insert into bots (id, token, company_id, created_at, updated_at) values (?,?,?,?,?)`
	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(bot.ID, bot.Token, bot.Company, bot.CreatedAt, bot.UpdatedAt)
	if err != nil {
		return fmt.Errorf("executing upsert bot statement %w", err)
	}
	return nil
}

func (r *Repository) GetBotByID(id string) (*models.Bot, error) {
	const query = `select * from bots where id = $1 and deleted_at is null limit 1`
	row := r.db.QueryRow(query, id)
	var bot models.Bot
	if err := row.Scan(&bot.ID, &bot.Token); err != nil {
		return nil, err
	}
	return &bot, nil
}
