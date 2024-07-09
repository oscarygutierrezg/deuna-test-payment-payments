package testing

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"payment-payments-api/internal/models"
	"testing"
)

func TestCreateMunicipalitySuccess(t *testing.T) {
	assert := assert.New(t)

	municipality := models.Municipality{
		Country: models.MunicipalityCountry,
		Region:  "Region Metropolitana",
		Name:    "Santiago",
		Kind:    models.MunicipalityBase,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/municipalities/",
		Body:   &municipality,
	}

	TestSetApiRoutes(t)
	w, resp, err := simulateJSONRequest(router, createRequest, true)

	assert.Equal(http.StatusOK, w.Code, resp.Data)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Municipality created successfully.", resp.Message, resp.Data)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
}

func TestGetMunicipalitySuccess(t *testing.T) {
	assert := assert.New(t)

	municipality := models.Municipality{
		Country: models.MunicipalityCountry,
		Region:  "Region Metropolitana",
		Name:    "Santiago",
		Kind:    models.MunicipalityBase,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/municipalities/",
		Body:   &municipality,
	}

	var getRequest = Request{
		Method: http.MethodGet,
		Path:   "/municipalities/",
	}

	TestSetApiRoutes(t)
	_, resp, createMunicipalityErr := simulateJSONRequest(router, createRequest, true)
	municipalityId := resp.Data.(map[string]interface{})["_id"].(string)
	getRequest.Path += municipalityId
	w, resp, err := simulateJSONRequest(router, getRequest, true)

	assert.Nil(createMunicipalityErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Municipality found successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
}

func TestUpdateMunicipalitySuccess(t *testing.T) {
	assert := assert.New(t)

	municipality := models.Municipality{
		Country: models.MunicipalityCountry,
		Region:  "Region Metropolitana",
		Name:    "Santiago",
		Kind:    models.MunicipalityBase,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/municipalities/",
		Body:   &municipality,
	}

	var updateRequest = Request{
		Method: http.MethodPut,
		Path:   "/municipalities/",
		Body: models.Municipality{
			Country: models.MunicipalityCountry,
			Region:  "Region Metropolitana",
			Name:    "Santiago",
			Kind:    models.MunicipalityExtreme,
		},
	}

	TestSetApiRoutes(t)
	_, resp, createMunicipalityErr := simulateJSONRequest(router, createRequest, true)
	municipalityId := resp.Data.(map[string]interface{})["_id"].(string)
	updateRequest.Path += municipalityId
	w, resp, err := simulateJSONRequest(router, updateRequest, true)

	assert.Nil(createMunicipalityErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Municipality updated successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
}

func TestDeleteMunicipalitySuccess(t *testing.T) {
	assert := assert.New(t)

	municipality := models.Municipality{
		Country: models.MunicipalityCountry,
		Region:  "Region Metropolitana",
		Name:    "Santiago",
		Kind:    models.MunicipalityBase,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/municipalities/",
		Body:   &municipality,
	}

	var deleteRequest = Request{
		Method: http.MethodDelete,
		Path:   "/municipalities/",
	}

	TestSetApiRoutes(t)
	_, resp, createMunicipalityErr := simulateJSONRequest(router, createRequest, true)
	municipalityId := resp.Data.(map[string]interface{})["_id"].(string)
	deleteRequest.Path += municipalityId
	w, resp, err := simulateJSONRequest(router, deleteRequest, true)

	assert.Nil(createMunicipalityErr)
	assert.Equal(http.StatusOK, w.Code, resp.Data)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Municipality deleted successfully.", resp.Message)
	assert.Nil(resp.Data)
}

func TestListMunicipalitySuccess(t *testing.T) {
	assert := assert.New(t)

	municipality := models.Municipality{
		Country: models.MunicipalityCountry,
		Region:  "Region Metropolitana",
		Name:    "Santiago",
		Kind:    models.MunicipalityBase,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/municipalities/",
		Body:   &municipality,
	}

	var listRequest = Request{
		Method: http.MethodGet,
		Path:   "/municipalities/list",
	}

	TestSetApiRoutes(t)
	_, _, createErr := simulateJSONRequest(router, createRequest, true)
	w, resp, err := simulateJSONRequest(router, listRequest, true)

	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code, resp.Data)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Municipality list successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType([]interface{}{}, resp.Data)
	assert.Len(resp.Data, 2)
}
