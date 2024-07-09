package umdw

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"reflect"
)

const BodyKey = "body"

func BodyContext(c *gin.Context) {

	var body map[string]interface{}
	_ = c.ShouldBind(&body)

	if c.Keys == nil {
		c.Keys = map[string]interface{}{}
	}
	c.Keys[BodyKey] = body

	c.Next()
}

func BodyParse(o interface{}, c *gin.Context) error {
	if reflect.TypeOf(o).Kind() != reflect.Ptr {
		return errors.New("o parameter must be a pointer &Struct")
	}

	m, ok := c.Keys[BodyKey]
	if !ok {
		return errors.New("BodyKey doesn't exists")
	}

	jsonStr, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonStr, &o)
	if err != nil {
		return err
	}

	return nil
}

func BodyGetMultipartFormDataFile(c *gin.Context) ([]byte, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return nil, err
	}

	content, err := file.Open()
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, file.Size)
	_, err = content.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
