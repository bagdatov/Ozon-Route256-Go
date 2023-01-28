package migrations

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" //nolint:golint
	_ "github.com/golang-migrate/migrate/v4/source/file"       //nolint:golint
	"gitlab.ozon.dev/bagdatov/homework-3/service-create-order/config"
	"log"
)

type service struct {
	mig *migrate.Migrate
}

func New(conf *config.Config) (*service, error) {

	source := "file://service-create-order/migrations"
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=public",
		conf.PG.Username, conf.PG.Password, conf.PG.Host, conf.PG.Port, conf.PG.Dbname)

	m, err := migrate.New(source, url)
	if err != nil {
		return nil, err
	}

	return &service{
		mig: m,
	}, nil
}

func (s *service) RunBlocking() error {

	log.Println("Migration: start")

	err := s.mig.Up()
	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
		log.Println("Migration: ErrNoChange happened")
	}

	log.Println("Migration: done")

	ver, dirty, err := s.mig.Version()
	if err != nil {
		if errors.Is(err, migrate.ErrNilVersion) {
			log.Println("Migration: no service has been applied")
		} else {
			log.Println("Migration: version err: ", err)
		}
	}

	log.Printf("Migration: version: %d, dirty: %v\n", ver, dirty)
	return nil
}

func (s *service) Close() error {
	errSource, errDatabase := s.mig.Close()
	if errSource != nil {
		return errSource
	}

	if errDatabase != nil {
		return errSource
	}

	return nil
}
