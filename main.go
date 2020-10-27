package main

import (
	"database/sql"
	"fmt"

	dbConfig "./dbconfig"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func main() {

	fmt.Printf("Accessing %s ... ", dbConfig.DbName)
	db, err = sql.Open(dbConfig.PostgresDriver, dbConfig.DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}

	defer db.Close()

	sqlSelect()
	//sqlSelectID()
	//sqlInsert()
	//sqlUpdate()
	//sqlDelete()
}

func sqlSelect() {

	sqlStatement, err := db.Query("SELECT id, title, body FROM " + dbConfig.TableName)
	checkErr(err)

	for sqlStatement.Next() {

		var article dbConfig.Article

		err = sqlStatement.Scan(&article.ID, &article.Title, &article.Body)
		checkErr(err)

		fmt.Printf("%d\t%s\t%s \n", article.ID, article.Title, article.Body)
	}
}

func sqlSelectID() {

	var article dbConfig.Article

	sqlStatement := fmt.Sprintf("SELECT id, title, body FROM %s where id = $1", dbConfig.TableName)

	err = db.QueryRow(sqlStatement, 1).Scan(&article.ID, &article.Title, &article.Body)
	checkErr(err)

	fmt.Printf("%d\t%s\t%s \n", article.ID, article.Title, article.Body)
}

func sqlInsert() {

	sqlStatement := fmt.Sprintf("INSERT INTO %s VALUES ($1,$2, $3)", dbConfig.TableName)

	insert, err := db.Prepare(sqlStatement)
	checkErr(err)

	result, err := insert.Exec(5, "Maps in Golang", "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium")
	checkErr(err)

	affect, err := result.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func sqlUpdate() {

	sqlStatement := fmt.Sprintf("update %s set body=$1 where id=$2", dbConfig.TableName)

	update, err := db.Prepare(sqlStatement)
	checkErr(err)

	result, err := update.Exec("But I must explain to you how all this mistaken idea", 5)
	checkErr(err)

	affect, err := result.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func sqlDelete() {

	sqlStatement := fmt.Sprintf("delete from %s where id=$1", dbConfig.TableName)

	delete, err := db.Prepare(sqlStatement)
	checkErr(err)

	result, err := delete.Exec(5)
	checkErr(err)

	affect, err := result.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
