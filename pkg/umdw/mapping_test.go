package umdw

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetMapFromStruct(t *testing.T) {
	assert := assert.New(t)
	type userInfo struct {
		FirstName    string `json:"firstName"`
		LastName     string `json:"lastName"`
		Rut          string `json:"rut"`
		Phone        string `json:"phone"`
		Password     string `json:"password,omitempty"`
		PasswordSalt string `json:"passwordSalt,omitempty"`
		IsEnabled    bool   `json:"isEnabled"`
	}
	type admin struct {
		Id          string    `json:"_id,omitempty"`
		Rev         string    `json:"_rev,omitempty"`
		Type        string    `json:"type"`
		CreatedTime time.Time `json:"createdTime"`
		UpdatedTime time.Time `json:"updatedTime"`
		Info        userInfo  `json:"info"`
	}
	var a admin

	m, err := GetMapFromStruct(a)
	i, ok := m["info"].(map[string]interface{})

	assert.Nil(err)
	assert.IsType(new(map[string]interface{}), &m)
	assert.True(ok, "ConvertObjToMap doesn't transform sub levels")
	assert.IsType(new(map[string]interface{}), &i)
}

func TestGetPathFromMap(t *testing.T) {
	assert := assert.New(t)
	m := map[string]interface{}{
		"type": "admin",
		"info": map[string]interface{}{
			"firstName": "John",
			"lastName":  "Doe",
		},
	}

	o, err := GetPathFromMap(m, "info.lastName")
	v, ok := o.(string)

	assert.Nil(err)
	assert.True(ok)
	assert.Equal("Doe", v)
}
