package tables

import (
	"github.com/jackc/pgx"
	"log"
	"time"
)

type ActivityType int16

const (
	ActivityType_Unknown ActivityType = iota
	ActivityType_CheckIn
	ActivityType_CheckOut
)

type ActivitiesTableRecord struct {
	ID       uint64       `json:"id"`
	UserID   int16        `json:"user_id"`
	Office   int16        `json:"office"`
	Datetime time.Time    `json:"datetime"`
	Type     ActivityType `json:"type"`
}

type ActivitiesTable interface {
	GetAll() ([]ActivitiesTableRecord, error)
	GetByUserID(userID int16) ([]ActivitiesTableRecord, error)
	GetByOfficeID(officeID int16) ([]ActivitiesTableRecord, error)
	Insert(record ActivitiesTableRecord) (uint64, error)
	Delete(id uint64) error
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
		"userid SMALLINT, " +
		"office SMALLINT, " +
		"datetime TIMESTAMP, " +
		"type SMALLINT, " +
		"FOREIGN KEY (office) REFERENCES offices(id), " +
		"FOREIGN KEY (userid) REFERENCES users(id) " +
		")")
	if err != nil {
		log.Fatalf("failed to initial activities table with error: %v", err)
	}
}

func (t *activitiesTableImpl) GetAll() ([]ActivitiesTableRecord, error) {
	rows, err := t.conn.Query("SELECT * FROM activities")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []ActivitiesTableRecord
	for rows.Next() {
		var record ActivitiesTableRecord
		err = rows.Scan(&record.ID, &record.UserID, &record.Office, &record.Datetime, &record.Type)
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

func (t *activitiesTableImpl) GetByUserID(userID int16) ([]ActivitiesTableRecord, error) {
	rows, err := t.conn.Query("SELECT * FROM activities WHERE userid = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []ActivitiesTableRecord
	for rows.Next() {
		var record ActivitiesTableRecord
		err = rows.Scan(&record.ID, &record.UserID, &record.Office, &record.Datetime, &record.Type)
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

func (t *activitiesTableImpl) GetByOfficeID(officeID int16) ([]ActivitiesTableRecord, error) {
	rows, err := t.conn.Query("SELECT * FROM activities WHERE office = $1", officeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []ActivitiesTableRecord
	for rows.Next() {
		var record ActivitiesTableRecord
		err = rows.Scan(&record.ID, &record.UserID, &record.Office, &record.Datetime, &record.Type)
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

func (t *activitiesTableImpl) Insert(record ActivitiesTableRecord) (uint64, error) {
	var id uint64
	err := t.conn.QueryRow("INSERT INTO activities (userid, office, datetime, type) VALUES ($1, $2, $3, $4) RETURNING id",
		record.UserID, record.Office, record.Datetime, record.Type).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (t *activitiesTableImpl) Delete(id uint64) error {
	_, err := t.conn.Exec("DELETE FROM activities WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
