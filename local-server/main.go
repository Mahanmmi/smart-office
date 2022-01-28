package main

import (
	"github.com/Mahanmmi/fuzzy-lamp/local-server/db"
	"github.com/Mahanmmi/fuzzy-lamp/local-server/src"
	"github.com/Mahanmmi/fuzzy-lamp/local-server/config"
	"github.com/jackc/pgx"
	"log"
)

func main() {
	conf := config.NewLocalServerConfig()
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

	database := db.NewLocalServerDatabase(conn)
	src.NewLocalServer(conf, database).Start()
}
