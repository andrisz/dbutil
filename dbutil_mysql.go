//go:build mysql

package dbutil

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func formatPlaceholder(index int) string {
	return "?"
}

func connect(host string, user string, password string, dbname string) (*sql.DB, error) {
	connstring := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbname)
	return sql.Open("mysql", connstring)
}
