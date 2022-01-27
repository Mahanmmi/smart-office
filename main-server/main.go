package main

import (
	"github.com/Mahanmmi/fuzzy-lamp/main-server/config"
	"github.com/Mahanmmi/fuzzy-lamp/main-server/db"
	"github.com/Mahanmmi/fuzzy-lamp/main-server/src"
	"github.com/jackc/pgx"
	"log"
)

func main() {
	conf := config.NewMainServerConfig()
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     conf.PostgresDB.Host,
		Port:     conf.PostgresDB.Port,
		Database: conf.PostgresDB.Database,
		User:     conf.PostgresDB.User,
		Password: conf.PostgresDB.Password,
	})
	if err != nil {
		log.Fatalf("failed to connect to postgres with error: %v", err)
	}

	database := db.NewMainServerDatabase(conn)
	src.NewMainServer(conf, database).Start()
}
