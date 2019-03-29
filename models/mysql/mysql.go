package mysql

import (
	"github.com/alimy/mir-music/models/core"
	"github.com/alimy/mir-music/models/model"
	"github.com/jmoiron/sqlx"
	"github.com/unisx/logus"

	_ "github.com/go-sql-driver/mysql" // mysql sql driver
)

type mysqlRepository struct {
	*core.Sqlx
}

// NewRepository build a new core.Repository that backend by mysql database
func NewRepository(config *model.Config) (core.Repository, error) {
	_, dsn := config.Dsn()
	logus.Debug("connect mysql", logus.String("dsn", dsn))
	mysqlDb, err := sqlx.Connect("mysql", dsn)
	return &mysqlRepository{Sqlx: &core.Sqlx{DB: mysqlDb}}, err
}
