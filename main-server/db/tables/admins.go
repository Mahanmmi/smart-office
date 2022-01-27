package tables

import (
	"github.com/jackc/pgx"
	"log"
)

type AdminsTableRecord struct {
	Username string
	Password string
	Office   int16
}

type AdminsTable interface {
	GetAll() ([]AdminsTableRecord, error)
	GetByUsername(username string) (AdminsTableRecord, error)
	Insert(record AdminsTableRecord) (string, error)
	Delete(username string) error
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

func (t *adminsTableImpl) Insert(record AdminsTableRecord) (string, error) {
	var id string
	err := t.conn.QueryRow("INSERT INTO admins (username, password, office) VALUES ($1, $2, $3) RETURNING username",
		record.Username, record.Password, record.Office).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (t *adminsTableImpl) GetAll() ([]AdminsTableRecord, error) {
	rows, err := t.conn.Query("SELECT username, password, office FROM admins")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []AdminsTableRecord
	for rows.Next() {
		var record AdminsTableRecord
		err = rows.Scan(&record.Username, &record.Password, &record.Office)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

func (t *adminsTableImpl) GetByUsername(username string) (AdminsTableRecord, error) {
	var record AdminsTableRecord
	err := t.conn.QueryRow("SELECT username, password, office FROM admins WHERE username = $1", username).Scan(&record.Username, &record.Password, &record.Office)
	if err != nil {
		return record, err
	}
	return record, nil
}

func (t *adminsTableImpl) Delete(username string) error {
	_, err := t.conn.Exec("DELETE FROM admins WHERE username = $1", username)
	if err != nil {
		return err
	}
	return nil
}
