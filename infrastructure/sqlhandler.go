package infrastructure

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"clean_arch_golang/interfaces/database"
)

type SqlHandler struct {
	Conn *sql.DB
}

type SqlResult struct {
	Result sql.Result
}

type SqlRow struct {
	Row sql.Row
}

var _ database.SqlHandler = (*SqlHandler)(nil)
var _ database.Result = (*SqlResult)(nil)
var _ database.Row = (*SqlRow)(nil)

func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open("mysql", "root:@tcp(db:3306/sample")
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (hander *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	return nil, nil
}

func (hander *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	return nil, nil
}

func (result *SqlResult) LastInsertId() (int64, error) {
	return 0, nil
}

func (result *SqlResult) RowsAffected() (int64, error) {
	return 0, nil
}

func (row *SqlRow) Scan(...interface{}) error {
	return nil
}
func (row *SqlRow) Next() bool {
	return false
}

func (row *SqlRow) Close() bool {
	return false
}
