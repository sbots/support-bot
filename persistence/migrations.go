package persistence

import (
	"database/sql"
	"fmt"
)

func migrate(db *sql.DB) error {
	const queryBots = `
	create table if not exists bots(
		id UUID primary key not null unique,
		token text unique not null,
		type text not null,
		created_at timestamp not null,
		updated_at timestamp not null
	)`
	statementBots, err := db.Prepare(queryBots)
	if err != nil {
		return fmt.Errorf("preparing bots statement: %w", err)
	}

	if _, err = statementBots.Exec(); err != nil {
		return fmt.Errorf("executing bots statement: %w", err)
	}

	const queryTenants = `
		create table if not exists tenants(
			id UUID primary key not null unique,
			name varchar(225) unique not null,
			created_at timestamp not null,
			updated_at timestamp not null
		);
	`

	statementTenant, err := db.Prepare(queryTenants)
	if err != nil {
		return fmt.Errorf("preparing tenants statement: %w", err)
	}

	if _, err = statementTenant.Exec(); err != nil {
		return fmt.Errorf("executing tenants statement: %w", err)
	}

	const queryUser = `
		create table if not exists users (
			id UUID primary key not null unique,
			email varchar(225) unique not null,
			name varchar(225) not null,
			surname varchar(225) not null,
		    password varchar(225) not null,
		    phone varchar(225) not null,
			created_at timestamp not null,
			updated_at timestamp not null,
		    company_id UUID not null,
			foreign key (company_id) references tenants (id) ON DELETE CASCADE ON UPDATE NO ACTION
		);`

	statementUser, err := db.Prepare(queryUser)
	if err != nil {
		return fmt.Errorf("preparing users statement: %w", err)
	}

	if _, err = statementUser.Exec(); err != nil {
		return fmt.Errorf("executing users statement: %w", err)
	}
	return nil
}
