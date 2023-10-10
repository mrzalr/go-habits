package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/mrzalr/go-habits/pkg/configuration"
)

func New(config configuration.Configuration) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Database.Mysql.User,
		config.Database.Mysql.Password,
		config.Database.Mysql.Host,
		config.Database.Mysql.Port,
		config.Database.Mysql.Dbname,
	)

	return sqlx.Connect(config.Database.Mysql.Driver, dsn)
}
