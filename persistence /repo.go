package persistence_

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"support-bot/models"
)

const table = "bots"

type Repository struct {
	db *sql.DB
	table string
}

func (r *Repository) CreateBot(bot *models.Bot) (*models.Bot, error){
	const query = `insert into bots (id, token) values (?,?)`
	statement, err := r.db.Prepare(query)
	if err != nil{
		return nil, err
	}
	_, err = statement.Exec(bot.ID, bot.Token)
	return &bot, err
}

func (r *Repository) GetBot(id string) (*models.Bot,error){
	const query = `select * from bots where id = $1 limit 1`
	row := r.db.QueryRow(query, id)
	var bot models.Bot
	if err := row.Scan(&bot.ID, &bot.Token); err != nil{
		return nil, err
	}
	return &bot, nil
}

func NewRepository() (*Repository, error){
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil{
		return nil, err
	}

	if err := prepare(db); err != nil{
		return nil, err
	}

	return &Repository{
		db:    db,
		table: table,
	}, nil
}

func prepare(db *sql.DB) error {
	const query = `
	create table bots(
		"id" text not null primary key,
		"token" text 
	)`
	statement, err := db.Prepare(query)
	if err != nil{
		return err
	}
	_, err = statement.Exec()
	return err
}

func (r *Repository) Close() {
	if err := r.db.Close(); err != nil{
		log.Fatalln(err)
	}
}