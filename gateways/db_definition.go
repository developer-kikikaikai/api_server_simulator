// +build !testdb

package gateways

const dbname = "APIServerDB"

func GetDBName() string {
	return dbname
}
