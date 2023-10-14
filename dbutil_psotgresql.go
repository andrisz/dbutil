//go:build postgres

package dbutil

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

func formatPlaceholder(index int) string {
	return fmt.Sprintf("$%d", index)
}

func connect(host string, user string, password string, dbname string) (*sql.DB, error) {
	connstring := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=", user, password, dbname)

	pos := strings.Index(host, ":")
	if pos != -1 {
		connstring += fmt.Sprintf("%s port=%s", host[:pos], host[pos+1:])
	} else {
		connstring += host
	}

	return sql.Open("postgres", connstring)
}
