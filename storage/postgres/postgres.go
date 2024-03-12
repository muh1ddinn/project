package storage

import (
	"cars_with_sql/config"
	"cars_with_sql/storage"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Store struct {
	DB *sql.DB
}

func Neww(cfg config.Config) (storage.IStorage, error) {
	url := fmt.Sprintf(`host=%s port=%v user=%s password=%s database=%s sslmode=disable`,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return Store{
		DB: db,
	}, nil
}
func (s Store) CloseDB() {
	s.DB.Close()
}

func (s Store) Car() storage.ICarstorage {
	newwCar := Newwcar(s.DB)

	return &newwCar
}

func (s Store) Customer() storage.ICustomerStorage {
	NewCustomer := (s.DB)

	return &NewCustomer
}
