package dbutil

import (
	"database/sql"
	"strings"
)

func SetPlaceholders(stmt string) string {
	index := 1
	offset := 0
	var sb strings.Builder

	for {
		pos := strings.Index(stmt[offset:], "?")
		if pos == -1 {
			break
		}
		pos += offset
		sb.WriteString(stmt[offset:pos])

		if pos != len(stmt)-1 && stmt[pos+1] == '?' {
			sb.WriteByte('?')
			pos++
		} else {
			sb.WriteString(formatPlaceholder(index))
			index++
		}
		offset = pos + 1
	}

	sb.WriteString(stmt[offset:])

	return sb.String()
}

func Connect(host string, user string, password string, dbname string) (*sql.DB, error) {
	return connect(host, user, password, dbname)
}
