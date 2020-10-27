package dbconfig

import "fmt"

// Article ...
type Article struct {
	ID    int
	Title string
	Body  []byte
}

// PostgresDriver ...
const PostgresDriver = "postgres"

// User ...
const User = "postgres"

// Host ...
const Host = "localhost"

// Port ...
const Port = "5432"

// Password ...
const Password = "vapordev123"

// DbName ...
const DbName = "MobyDick"

// TableName ...
const TableName = "article"

// DataSourceName ...
var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
