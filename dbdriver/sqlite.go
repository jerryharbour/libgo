package dbdriver

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Sqlite struct {
	conf DBConfig
	db   *gorm.DB
}

func NewSqlite(dbFilePath string) *Sqlite {
	s := Sqlite{
		conf: NewSqliteConfig(dbFilePath),
		db:   nil,
	}

	return &s
}

func (s *Sqlite) Open(logger logger.Interface) error {
	if s.db != nil {
		s.Close()
	}

	var err error
	s.db, err = gorm.Open(sqlite.Open(s.conf.GetDSN()), &gorm.Config{
		PrepareStmt:    true,
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger:         logger})
	if err != nil {
		return err
	}

	return nil
}

func (s *Sqlite) Close() {
	if s.db == nil {
		return
	}

	sqlDB, _ := s.db.DB()
	sqlDB.Close()
	s.db = nil
}

func (s *Sqlite) DB() *gorm.DB {
	return s.db
}
