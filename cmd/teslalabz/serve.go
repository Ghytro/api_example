package main

import (
	"log"

	"github.com/gignhit/teslalabz/fixtures"
	"github.com/gignhit/teslalabz/internal/api/liquid"
	"github.com/gignhit/teslalabz/internal/api/orders"
	"github.com/gignhit/teslalabz/internal/config"
	"github.com/gignhit/teslalabz/internal/server"
	"github.com/go-pg/pg/v10"
)

func checkPGConn(db *pg.DB) error {
	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}

func serve() {
	dbOpts, err := pg.ParseURL(config.DBUrl)
	if err != nil {
		log.Fatal(err)
	}
	db := pg.Connect(dbOpts)
	if err := checkPGConn(db); err != nil {
		log.Fatal(err)
	}
	fixtures.DefineDBModel(db)
	liquidsApi := liquid.NewLiquidsApi(db)
	usersApi := orders.NewOrdersApi(db)

	server.NewGracefulShutdownServer(liquidsApi, usersApi).Listen()
}
