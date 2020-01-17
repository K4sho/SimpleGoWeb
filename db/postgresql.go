package db

// Взаимодействие с БД
// Конструируются и исполняются операторы SQL

import (
	"simpleGoWeb/model"
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	ConnectString string
}

type pgDb struct {
	dbConn *sqlx.DB

	sqlSelectPeople *sqlx.Stmt
	sqlInsertPeople *sqlx.NamedStmt
	sqlSelectPerson *sql.Stmt
}

func (p *pgDb) createTablesIfNotExist() error {
	createSql := `

	CREATE TABLE IF NOT EXIST people (
		id SERIAL NOT NULL PRIMARY KEY,
		first TEXT NOT NULL,
		last TEXT NOT NULL
	);

	`

	if rows, err := p.dbConn.Query(createSql); err != nil {
		return err
	} else {
		rows.Close()
	}
	return nil
}

func (p *pgDb) prepareSqlStatements() (err error) {
	if p.sqlSelectPeople, err = p.dbConn.Preparex(
		"SELECT id, first, last FROM people",

		); err != nil {
			return err
	}
	if p.sqlInsertPeople, err = p.dbConn.PrepareNamed("INSERT INTO people (first, last) VALUES (:first, :last " +
		"RETURNING id, first, last",); err != nil {
		return err
	}
	if p.sqlSelectPerson, err = p.dbConn.Prepare("SELECT id, first, last FROM people WHERE id = $1",); err != nil {
		return err
	}

	return nil
}

func (p* pgDb) SelectPeople() ([]*model.Person, error) {
	people := make([]*model.Person, 0)

	if err := p.sqlSelectPeople.Select(&people); err != nil {
		return nil, err
	}

	return people, nil
}

// Инициализация базы данных. Используем PostgreSQL
func InitDB(cfg Config) (*pgDb, error) {
	if dbConn, err := sqlx.Connect("postgres", cfg.ConnectString); err != nil {
		return nil, err
	} else {
		p := &pgDb{dbConn: dbConn}
		if err := p.dbConn.Ping(); err != nil {
			return nil, err
		}
		if err := p.createTablesIfNotExist(); err != nil {
			return nil, err
		}
		if err := p.prepareSqlStatements(); err != nil {
			return nil, err
		}
		return p, nil
	}
}