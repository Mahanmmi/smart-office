package tables

import (
	"github.com/jackc/pgx"
	"log"
)

type ActivitiesTable interface {
}

type activitiesTableImpl struct {
	conn *pgx.Conn
}

func NewActivitiesTable(conn *pgx.Conn) ActivitiesTable {
	t := &activitiesTableImpl{conn: conn}
	t.init()
	return t
}

func (t *activitiesTableImpl) init() {
	_, err := t.conn.Exec("CREATE TABLE IF NOT EXISTS activities (" +
		"id SERIAL PRIMARY KEY, " +
		"user SMALLINT, " +
		"office SMALLINT, " +
		"datetime TIMESTAMP, " +
		"type SMALLINT, " +
		"FOREIGN KEY (office) REFERENCES offices(id), " +
		"FOREIGN KEY (user) REFERENCES users(id) " +
		")")
	if err != nil {
		log.Fatalf("failed to initial activities table with error: %v", err)
	}
}
