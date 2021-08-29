package persistence

import "database/sql"

func migrate(db *sql.DB) error {
	const queryBots = `
	create table if not exists bots(
		"id" UUID not null primary key,
		"token" text not null,
		"type" text not null                   
	)`
	statementBots, err := db.Prepare(queryBots)
	if err != nil {
		return err
	}

	if _, err = statementBots.Exec(); err != nil {
		return err
	}

	const queryTenants = `
		create table if not exists tenants (
			id UUID not null,
			name varchar(225) not null,
			created_at timestamp not null,
			updated_at timestamp not null,
			primary key (id)
		);
	`

	statementTenant, err := db.Prepare(queryTenants)
	if err != nil {
		return err
	}

	if _, err = statementTenant.Exec(); err != nil {
		return err
	}

	const queryUser = `
		create table if not exists users (
			id UUID not null,
			email varchar(225) not null unique,
			name varchar(225) not null,
			surname varchar(225) not null,
			FOREIGN KEY company_id REFERENCES tenants (id) not null 
			password varchar(225) not null,
			token_hash varchar(15) not null,
			created_at timestamp not null,
			updated_at timestamp not null,
			primary key (id)
		);`

	statementUser, err := db.Prepare(queryUser)
	if err != nil {
		return err
	}

	if _, err = statementUser.Exec(); err != nil {
		return err
	}
	return nil
}
