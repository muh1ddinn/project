package storage

import (
	"database/sql"
	"fmt"
	"lms_backed_pr/configg"
	"lms_backed_pr/storage"

	_ "github.com/lib/pq"
)

type Store struct {
	DB *sql.DB
}

func New(cfg configg.Config) (storage.IStorage, error) {

	url := fmt.Sprintf(`host=%s port=%v user=%s password=%s database=%s sslmode=disable`,

		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

	db, err := sql.Open("postgres", url)
	fmt.Println("err opening :", err)
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

func (s Store) Student() storage.IStorage {
	newstudent := Newstudent(s.DB)
	return &newstudent

}
