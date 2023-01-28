package postgres

import (
	"fmt"
	"reflect"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Unique constraint violation code in Postgres
const errUniqueViolation = "23505"

// db is a struct
// that implements usecase.Database interface
type db struct {
	pgx *pgxpool.Pool
}

// New is a constructor for database
func New(pool *pgxpool.Pool) (*db, error) {
	if isNil(pool) {
		return nil, fmt.Errorf("pgx pool is nil")
	}

	return &db{
		pgx: pool,
	}, nil
}

// isNil is using reflect
// package to determine rather input is nil
func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
