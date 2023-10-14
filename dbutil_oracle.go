//go:build oracle

package dbutil

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/sijms/go-ora/v2"
)

func formatPlaceholder(index int) string {
	return fmt.Sprintf(":%d", index)
}

func connect(host string, user string, password string, dbname string) (*sql.DB, error) {

	connstring := fmt.Sprintf("oracle://%s:%s@", user, password)

	pos := strings.Index(host, ":")
	if pos != -1 {
		connstring += host
	} else {
		connstring += fmt.Sprintf("%s:1521", host)
	}

	connstring += "/" + dbname

	return sql.Open("oracle", connstring)
}
