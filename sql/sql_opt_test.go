package sql

import "testing"

func TestSqlOpt(t *testing.T) {
	InitDbConfig("mysql", "root", "", "127.0.0.1:3306", "dbname")
	SqlOpt()
}