package tables

import (
	"github.com/jackc/pgx"
	"log"
)

type UsersTableRecord struct {
	ID       int16
	Card     string
	Light    int16
	Room     int16
}

type UsersTable interface {
	GetAll() ([]UsersTableRecord, error)
	GetByID(id int16) (UsersTableRecord, error)
	Insert(record UsersTableRecord) (int16, error)
	UpdateLightByID(id int16, light int16) error
	Delete(id int16) error
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
		"card TEXT, " +
		"light SMALLINT, " +
		"office SMALLINT, " +
		"room SMALLINT, " +
		"FOREIGN KEY (office) REFERENCES offices(id) " +
		")")
	if err != nil {
		log.Fatalf("failed to initial users table with error: %v", err)
	}
}

func (t *usersTableImpl) GetAll() ([]UsersTableRecord, error) {
	var records []UsersTableRecord
	rows, err := t.conn.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var record UsersTableRecord
		err = rows.Scan(&record.ID, &record.Card, &record.Light, &record.Room)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

func (t *usersTableImpl) GetByID(id int16) (UsersTableRecord, error) {
	var record UsersTableRecord
	err := t.conn.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&record.ID, &record.Card, &record.Light, &record.Room)
	if err != nil {
		return record, err
	}
	return record, nil
}

func (t *usersTableImpl) Insert(record UsersTableRecord) (int16, error) {
	var id int16
	err := t.conn.QueryRow("INSERT INTO users (password, light, office, room) VALUES ($1, $2, $3, $4) RETURNING id", record.Card, record.Light, record.Room).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (t *usersTableImpl) UpdateLightByID(id int16, light int16) error {
	_, err := t.conn.Exec("UPDATE users SET light = $1 WHERE id = $2", light, id)
	if err != nil {
		return err
	}
	return nil
}

func (t *usersTableImpl) Delete(id int16) error {
	_, err := t.conn.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
