package datastore

import (
	"fmt"

	"github.com/arionalmond/go-api-boilerplate/config"
	// in context to connect to mysql datastore
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// MySQLDS is a struct with a MySql DB connection.
type MySQLDS struct {
	db *sqlx.DB
}

// GetMySQLDS returns a MysqlDS instance or an error if there is a problem connecting to the Mysql database
func GetMySQLDS(c config.Config) (msd *MySQLDS, err error) {
	db, err := sqlx.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			c.DBUsername,
			c.DBPassword,
			c.DBHost,
			c.DBPort,
			c.DBName),
	)

	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}

	// db.SetMaxIdleConns(c.DBMaxIdleConns)
	// db.SetConnMaxLifetime(time.Duration(c.DBMaxConnLifetimeInMinutes) * time.Minute)

	// TODO: Figure out best max idle, max open, & idle timeout configs for db
	msd = &MySQLDS{
		db: db,
	}

	return
}
