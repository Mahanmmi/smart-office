package tables

import (
	"github.com/jackc/pgx"
	"log"
)

type AdminsTable interface {
}

type adminsTableImpl struct {
	conn *pgx.Conn
}

func NewAdminsTable(conn *pgx.Conn) AdminsTable {
	t := &adminsTableImpl{conn: conn}
	t.init()
	return t
}

func (t *adminsTableImpl) init() {
	_, err := t.conn.Exec("CREATE TABLE IF NOT EXISTS admins (" +
		"username TEXT PRIMARY KEY, " +
		"password TEXT, " +
		"office SMALLINT, " +
		"FOREIGN KEY (office) REFERENCES offices(id) " +
		")")
	if err != nil {
		log.Fatalf("failed to initial admins table with error: %v", err)
	}
}
