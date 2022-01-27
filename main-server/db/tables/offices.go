package tables

import (
	"github.com/jackc/pgx"
	"log"
	"time"
)

type OfficesTableRecord struct {
	ID           int16
	LightOnTime  time.Time
	LightOffTime time.Time
}

type OfficesTable interface {
	GetAll() ([]OfficesTableRecord, error)
	GetByID(id int16) (OfficesTableRecord, error)
	Insert(record OfficesTableRecord) (int16, error)
	Update(record OfficesTableRecord) error
	Delete(id int16) error
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
		"light_on_time TIMESTAMP, " +
		"light_off_time TIMESTAMP " +
		")")
	if err != nil {
		log.Fatalf("failed to initial offices table with error: %v", err)
	}
}

func (t *officesTableImpl) GetAll() ([]OfficesTableRecord, error) {
	var records []OfficesTableRecord
	rows, err := t.conn.Query("SELECT id, light_on_time, light_off_time FROM offices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var record OfficesTableRecord
		err = rows.Scan(&record.ID, &record.LightOnTime, &record.LightOffTime)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

func (t *officesTableImpl) GetByID(id int16) (OfficesTableRecord, error) {
	var record OfficesTableRecord
	err := t.conn.QueryRow("SELECT id, light_on_time, light_off_time FROM offices WHERE id = $1", id).Scan(&record.ID, &record.LightOnTime, &record.LightOffTime)
	if err != nil {
		return record, err
	}
	return record, nil
}

func (t *officesTableImpl) Insert(record OfficesTableRecord) (int16, error) {
	var id int16
	err := t.conn.QueryRow("INSERT INTO offices (light_on_time, light_off_time) VALUES ($1, $2) RETURNING id", record.LightOnTime, record.LightOffTime).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (t *officesTableImpl) Update(record OfficesTableRecord) error {
	_, err := t.conn.Exec("UPDATE offices SET light_on_time = $1, light_off_time = $2 WHERE id = $3", record.LightOnTime, record.LightOffTime, record.ID)
	if err != nil {
		return err
	}
	return nil
}

func (t *officesTableImpl) Delete(id int16) error {
	_, err := t.conn.Exec("DELETE FROM offices WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
