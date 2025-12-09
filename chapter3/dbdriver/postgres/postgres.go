package postgres

import (
	"database/sql"
	"database/sql/driver"
	"errors"
)

// PostgresDriver 为 sql 包
// 提供我们的实现。
type PostgresDriver struct{}

// Open 提供到数据库的连接。
func (dr PostgresDriver) Open(string) (driver.Conn, error) {
	return nil, errors.New("Unimplemented")
}

var d *PostgresDriver

// init 在 main 之前被调用。
func init() {
	d = new(PostgresDriver)
	sql.Register("postgres", d)
}
