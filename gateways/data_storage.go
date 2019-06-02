package gateways

import (
	"errors"
	"fmt"

	"github.com/developer-kikikaikai/api_server_simulator/datatypes"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

/*
DB name: Get by GetDBName
Collection:
- DefinedEndpoints
- ServerConfigurations
*/
const collectionNameEndpoints = "DefinedEndpoints"

/*
- DefinedEndpoints
Format:
{
  "endpoint": "string",
  "method": "string",
  "response": {
    "body": "string",
    "header": [
      {
        "type": "string",
        "value": "string"
      }
    ],
    "status": 0
  }
}
*/

const collectionNameConfig = "ServerConfigurations"

/*
- ServerConfigurations
{
	"EnableSwagger": True/False
	"Endpoint":"string"
}
*/
type ServerConfigurations struct {
	EnableSwagger bool   `json:enableSwagger`
	Endpoint      string `json:endpoint`
}

type DataStorage struct {
	connection *mgo.Session
	db         *mgo.Database
}

func (this *DataStorage) connectToDB() {
	dbname := GetDBName()
	this.connection, _ = mgo.Dial("mongodb://localhost/" + dbname)
	this.db = this.connection.DB(dbname)
}

func (this *DataStorage) disconnectFromDB() {
	this.connection.Close()
}

func (this *DataStorage) findDefinedEndpoints() []datatypes.Endpoint {
	collection := this.db.C(collectionNameEndpoints)
	iter := collection.Find(bson.M{})
	var result []datatypes.Endpoint

	err := iter.All(&result)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return nil
	}

	return result
}

func (this *DataStorage) findOneDefinedEndpoints(endpoint string, method string) (datatypes.Endpoint, error) {
	collection := this.db.C(collectionNameEndpoints)
	iter := collection.Find(bson.M{"endpoint": endpoint, "method": method})

	var result datatypes.Endpoint
	err := iter.One(&result)
	return result, err
}

func (this *DataStorage) insertDefinedEndpoints(endpoint *datatypes.Endpoint) error {
	collection := this.db.C(collectionNameEndpoints)
	if _, err := this.findOneDefinedEndpoints(endpoint.Endpoint, endpoint.Method); err == nil {
		return errors.New("There is a same key's data")
	}

	return collection.Insert(endpoint)
}

func (this *DataStorage) updateDefinedEndpoints(endpoint *datatypes.Endpoint) error {
	collection := this.db.C(collectionNameEndpoints)
	//filtering by Endpoint and Method
	filter := datatypes.ExistingEndpoint{endpoint.Endpoint, endpoint.Method}
	return collection.Update(filter, endpoint)
}

func (this *DataStorage) deleteDefinedEndpoints(uri string, method string) error {
	collection := this.db.C(collectionNameEndpoints)
	//filtering by Endpoint and Method
	filter := datatypes.ExistingEndpoint{uri, method}
	return collection.Remove(filter)
}

func (this *DataStorage) GetAll() []datatypes.Endpoint {
	this.connectToDB()
	defer this.disconnectFromDB()

	return this.findDefinedEndpoints()
}

func (this *DataStorage) Add(endpoint *datatypes.Endpoint) error {
	this.connectToDB()
	defer this.disconnectFromDB()

	return this.insertDefinedEndpoints(endpoint)
}

func (this *DataStorage) Update(endpoint *datatypes.Endpoint) error {
	this.connectToDB()
	defer this.disconnectFromDB()

	return this.updateDefinedEndpoints(endpoint)
}

func (this *DataStorage) Delete(uri string, method string) error {
	this.connectToDB()
	defer this.disconnectFromDB()

	return this.deleteDefinedEndpoints(uri, method)
}

func (this *DataStorage) IsVisuableAPIDefinition() bool {
	return true
}

func (this *DataStorage) GetAPIEndpoint() string {
	return "/endpoints"
}

func (this *DataStorage) GetHTTPSConfiguration() datatypes.HTTPSConfiguration {
	return datatypes.HTTPSConfiguration{61443, "", ""}
}
