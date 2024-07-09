package testing

import (
	"github.com/stretchr/testify/assert"
	"payment-payments-api/pkg/mongodb"
	"testing"
)

var db mongodb.Database

const dbName = "test"

func TestInitMongoDBTest(t *testing.T) {
	assert := assert.New(t)

	errConn := db.ConnectStr("mongodb://127.0.0.1:27017", dbName)

	assert.Nil(errConn)

	_ = db.DeleteDB()
}
