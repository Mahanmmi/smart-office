package test

import (
	"github.com/Mahanmmi/fuzzy-lamp/main-server/config"
	"github.com/Mahanmmi/fuzzy-lamp/main-server/db"
	"github.com/Mahanmmi/fuzzy-lamp/main-server/src"
	"github.com/jackc/pgx"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
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
	go src.NewMainServer(conf, database).Start()

	m.Run()
}

func Test(t *testing.T) {

}
