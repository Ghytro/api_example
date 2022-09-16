package repository

import "github.com/go-pg/pg/v10"

type baseRepo struct {
	db *pg.DB
}
