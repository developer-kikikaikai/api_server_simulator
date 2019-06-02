// +build testdb

package gateways

const dbname = "TestAPIServerDB"

func GetDBName() string {
	return dbname
}
