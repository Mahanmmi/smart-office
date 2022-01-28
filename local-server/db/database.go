package db

import (
	"github.com/Mahanmmi/fuzzy-lamp/local-server/db/tables"
	"github.com/jackc/pgx"
)

type LocalServerDatabase struct {
	Users      tables.UsersTable
}

func NewLocalServerDatabase(conn *pgx.Conn) *LocalServerDatabase {
	return &LocalServerDatabase{
		Users:      tables.NewUsersTable(conn),
	}
}
