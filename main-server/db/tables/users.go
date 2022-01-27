package tables

import (
	"github.com/jackc/pgx"
	"log"
)

type UsersTable interface {
}

type usersTableImpl struct {
	conn *pgx.Conn
}

func NewUsersTable(conn *pgx.Conn) UsersTable {
	t := &usersTableImpl{conn: conn}
	t.init()
	return t
}

func (t *usersTableImpl) init() {
	_, err := t.conn.Exec("CREATE TABLE IF NOT EXISTS users (" +
		"id SERIAL PRIMARY KEY, " +
		"password TEXT, " +
		"light SMALLINT, " +
		"office SMALLINT, " +
		"room SMALLINT, " +
		"FOREIGN KEY (office) REFERENCES offices(id) " +
		")")
	if err != nil {
		log.Fatalf("failed to initial users table with error: %v", err)
	}
}
