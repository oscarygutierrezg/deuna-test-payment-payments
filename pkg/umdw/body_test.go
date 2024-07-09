package umdw

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"testing"
)

func TestBodyContext(t *testing.T) {
	assert := assert.New(t)

	m := map[string]interface{}{
		"firstName": "John",
		"lastName":  "Doe",
		"age":       float64(38),
	}

	body, _ := json.Marshal(m)

	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/testing", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")

	c := gin.Context{
		Request: req,
	}

	BodyContext(&c)
	bodyCtx, ok := c.Keys[BodyKey]

	assert.True(ok)
	assert.NotNil(bodyCtx)
	assert.IsType(map[string]interface{}{}, bodyCtx)
	assert.Equal(m["firstName"], bodyCtx.(map[string]interface{})["firstName"])
	assert.Equal(m["lastName"], bodyCtx.(map[string]interface{})["lastName"])
	assert.Equal(m["age"], bodyCtx.(map[string]interface{})["age"])
}

func TestBodyParse(t *testing.T) {
	assert := assert.New(t)

	m := map[string]interface{}{
		"firstName": "John",
		"lastName":  "Doe",
		"age":       float64(38),
	}

	type person struct {
		FirstName string  `json:"firstName"`
		LastName  string  `json:"lastName"`
		Age       float64 `json:"age"`
	}

	var p person

	body, _ := json.Marshal(m)

	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/testing", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")

	c := gin.Context{
		Request: req,
	}

	BodyContext(&c)
	err := BodyParse(&p, &c)

	assert.Nil(err)
	assert.Equal(p.FirstName, m["firstName"])
	assert.Equal(p.LastName, m["lastName"])
	assert.Equal(p.Age, m["age"])
}

func TestBodyGetMultipartFormDataFile(t *testing.T) {
	assert := assert.New(t)

	const fileContent = "Hello World!"

	var body bytes.Buffer
	multipartWriter := multipart.NewWriter(&body)
	fw, multipartWriterErr := multipartWriter.CreateFormFile("file", "test.txt")
	r := strings.NewReader(fileContent)
	_, copyErr := io.Copy(fw, r)
	_ = multipartWriter.Close()

	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/testing", &body)
	req.Header.Add("Content-Type", multipartWriter.FormDataContentType())

	c := gin.Context{
		Request: req,
	}

	BodyContext(&c)
	content, err := BodyGetMultipartFormDataFile(&c)

	assert.Nil(multipartWriterErr)
	assert.Nil(copyErr)
	assert.Nil(err)
	assert.Equal(string(content), fileContent)
}
