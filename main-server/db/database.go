package db

import (
	"github.com/Mahanmmi/fuzzy-lamp/main-server/db/tables"
	"github.com/jackc/pgx"
)

type MainServerDatabase struct {
	Offices    tables.OfficesTable
	Admins     tables.AdminsTable
	Users      tables.UsersTable
	Activities tables.ActivitiesTable
}

func NewMainServerDatabase(conn *pgx.Conn) *MainServerDatabase {
	return &MainServerDatabase{
		Offices:    tables.NewOfficesTable(conn),
		Admins:     tables.NewAdminsTable(conn),
		Users:      tables.NewUsersTable(conn),
		Activities: tables.NewActivitiesTable(conn),
	}
}
