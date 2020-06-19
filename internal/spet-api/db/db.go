package db

import (
	"SPET/internal/spet-api/settings"
	"database/sql"
)

// Хранилище
type Store struct {
	config *settings.Config
	db     *sql.DB
}

// Создать экземпляр хранилища
func New(config *settings.Config) *Store {
	return &Store{
		config: config,
	}
}

// Подключение к базе данных
func (s *Store) OpenPostgres() error {
	// Открываем соединение
	db, err := sql.Open(s.config.Database.Name, s.config.Database.URL)
	if err != nil {
		return err
	}

	// Проверяем на доступность
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db

	return nil
}

// Закрываем соединение с базой данной
func (s *Store) Close() error {
	err := s.db.Close()
	if err != nil {
		return err
	}
	return nil
}

//func (s *Store) Student() *StudentRep {
//	if s.studentRep != nil {
//		return s.studentRep
//	}
//
//	s.studentRep = &StudentRep{
//		store: s,
//	}
//
//	return s.studentRep
//}
