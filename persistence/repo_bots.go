package persistence

import (
	"fmt"
	"support-bot/models"
)

func (r *Repository) UpsertBot(bot *models.Bot) (*models.Bot, error) {
	const query = `insert into bots (id, token) values (?,?)`
	statement, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	_, err = statement.Exec(bot.ID, bot.Token)
	if err != nil {
		return nil, fmt.Errorf("executing upsert bot statement %w", err)
	}
	return bot, nil
}

func (r *Repository) GetBot(id string) (*models.Bot, error) {
	const query = `select * from bots where id = $1 limit 1`
	row := r.db.QueryRow(query, id)
	var bot models.Bot
	if err := row.Scan(&bot.ID, &bot.Token); err != nil {
		return nil, err
	}
	return &bot, nil
}
