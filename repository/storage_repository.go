package repository

import (
	"github.com/developer-kikikaikai/api_server_simulator/gateways"
)

//define all interface
var dataStorage = gateways.DataStorage{}

func GetDataStorageAccesser() DataStorageAccesser {
	return &dataStorage
}
