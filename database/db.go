package database

var connection string

func init() {
	connection = "mySql"
}

func Getdatabase() string {
	return connection
}
