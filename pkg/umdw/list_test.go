package umdw

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestListSet(t *testing.T) {
	assert := assert.New(t)

	var list List

	err := list.Set("10", "10", "asc", "name")

	assert.Nil(err)
	assert.Equal(10, list.Skip)
	assert.Equal(10, list.Limit)
	assert.Equal("asc", list.Sort)
	assert.Equal("name", list.By)
}

func TestListContext(t *testing.T) {
	assert := assert.New(t)

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/testing/list?skip=10&limit=10&sort=desc&by=name", nil)

	c := gin.Context{
		Request: req,
	}

	list, err := ListContext(&c)

	assert.Nil(err)
	assert.Equal(10, list.Skip)
	assert.Equal(10, list.Limit)
	assert.Equal("desc", list.Sort)
	assert.Equal("name", list.By)
}
