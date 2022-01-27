package tables

import (
	"github.com/jackc/pgx"
	"log"
)

type OfficesTable interface {
}

type officesTableImpl struct {
	conn *pgx.Conn
}

func NewOfficesTable(conn *pgx.Conn) OfficesTable {
	t := &officesTableImpl{conn: conn}
	t.init()
	return t
}

func (t *officesTableImpl) init() {
	_, err := t.conn.Exec("CREATE TABLE IF NOT EXISTS offices (" +
		"id SMALLINT PRIMARY KEY, " +
		"light_on_time time, " +
		"light_off_time time " +
		")")
	if err != nil {
		log.Fatalf("failed to initial offices table with error: %v", err)
	}
}
