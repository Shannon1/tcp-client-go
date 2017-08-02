package sql

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var driverName_ string
var user_		string
var password_	string
var host_ 		string
var schema_		string

func InitDbConfig(driverName, user, password, host, schema string) {
	driverName_, user_, password_, host_, schema_ = driverName, user, password, host, schema

}

func generateSourceName() string {
	dbSourceName := ""
	dbSourceName += user_
	dbSourceName += ":"
	dbSourceName += password_
	dbSourceName += "@tcp("
	dbSourceName += host_
	dbSourceName += ")/"
	dbSourceName += schema_
	log.Println(dbSourceName)
	return dbSourceName
}



func SqlOpt() {
	db, err := sql.Open(driverName_, generateSourceName())
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tb_exts_ad_index;")
	if err != nil{
		log.Fatalln(err)
	}

	defer rows.Close()

	columns, _ := rows.Columns() 	// columns stores the column's name
	scanArgs := make([]interface{}, len(columns))	// as a rows.Scan argument
	values := make([]sql.RawBytes, len(columns))	// Make a slice for the values
	for i := range values {
		scanArgs[i] = &values[i]	// copy the references into such a slice
	}
	recordSet := make(map[string]string)	// save one record into a key-value map, key matches column's name
	for rows.Next() {

		err = rows.Scan(scanArgs...)
		if err !=nil{
			log.Fatalln(err)
		}

		for i, col := range values {
			if col == nil {
				recordSet[columns[i]] = "NULL"
			} else {
				recordSet[columns[i]] = string(col)
			}
		}

		log.Println(recordSet)
	}

	stmt, err := db.Prepare("INSERT INTO tb_exts_ad_index(product_id, ad_version, client_type, description) " +
		"VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec("TestApp", 999, 999, "中文测试")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(res.LastInsertId())
	log.Print(res.RowsAffected())
}