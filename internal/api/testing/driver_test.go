package testing

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"payment-payments-api/internal/models"
	"testing"
)

func TestCreateDriverSuccess(t *testing.T) {
	assert := assert.New(t)

	driver := models.Driver{
		FirstName: "John",
		LastName:  "Doe",
		Rut:       "11.111.111-1",
		Phone:     "+56987654321",
		Email:     "driver@example.com",
	}

	var request = Request{
		Method: http.MethodPost,
		Path:   "/drivers/",
		Body:   driver,
	}

	TestSetApiRoutes(t)
	w, resp, err := simulateJSONRequest(router, request, true)

	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Driver created successfully.", resp.Message, resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
	assert.Nil(resp.Data.(map[string]interface{})["password"])
	assert.Nil(resp.Data.(map[string]interface{})["passwordSalt"])
}

func TestCreateDriverEmailExistsFailed(t *testing.T) {
	assert := assert.New(t)

	driver := models.Driver{
		FirstName: "John",
		LastName:  "Doe",
		Rut:       "11.111.111-1",
		Phone:     "+56987654321",
		Email:     "driver@example.com",
	}

	var request = Request{
		Method: http.MethodPost,
		Path:   "/drivers/",
		Body:   driver,
	}

	TestSetApiRoutes(t)
	_, _, _ = simulateJSONRequest(router, request, true)
	w, resp, err := simulateJSONRequest(router, request, true)

	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Nil(err)
	assert.Equal("error", resp.Status)
	assert.Equal("Driver cannot be created.", resp.Message, resp.Data)
	assert.Equal("email exists.", resp.Data)
}

func TestCreateDriverRutExistsFailed(t *testing.T) {
	assert := assert.New(t)

	driver := models.Driver{
		FirstName: "John",
		LastName:  "Doe",
		Rut:       "11.111.111-1",
		Phone:     "+56987654321",
		Email:     "driver@example.com",
	}

	var request = Request{
		Method: http.MethodPost,
		Path:   "/drivers/",
		Body:   &driver,
	}

	TestSetApiRoutes(t)
	_, _, _ = simulateJSONRequest(router, request, true)
	driver.Email = "another_driver@example.com"
	w, resp, err := simulateJSONRequest(router, request, true)

	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Nil(err)
	assert.Equal("error", resp.Status)
	assert.Equal("Driver cannot be created.", resp.Message, resp.Data)
	assert.Equal("rut exists.", resp.Data)
}

func TestGetDriverSuccess(t *testing.T) {
	assert := assert.New(t)

	driver := models.Driver{
		FirstName: "John",
		LastName:  "Doe",
		Rut:       "11.111.111-1",
		Phone:     "+56987654321",
		Email:     "driver@example.com",
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/drivers/",
		Body:   driver,
	}

	var getRequest = Request{
		Method: http.MethodGet,
		Path:   "/drivers/",
	}

	TestSetApiRoutes(t)
	_, resp, createErr := simulateJSONRequest(router, createRequest, true)
	driverId := resp.Data.(map[string]interface{})["_id"].(string)
	getRequest.Path += driverId
	w, resp, err := simulateJSONRequest(router, getRequest, true)

	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Driver found successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
	assert.Nil(resp.Data.(map[string]interface{})["password"])
	assert.Nil(resp.Data.(map[string]interface{})["passwordSalt"])
}

func TestUpdateDriverSuccess(t *testing.T) {
	assert := assert.New(t)

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/drivers/",
		Body: models.Driver{
			FirstName: "John",
			LastName:  "Doe",
			Rut:       "11.111.111-1",
			Phone:     "+56987654321",
			Email:     "driver@example.com",
		},
	}

	var updateRequest = Request{
		Method: http.MethodPut,
		Path:   "/drivers/",
		Body: models.Driver{
			FirstName: "NewFirstName",
			LastName:  "NewLastName",
			Rut:       "11.111.111-1",
			Phone:     "+56987654321",
			Email:     "driver@example.com",
		},
	}

	TestSetApiRoutes(t)
	_, resp, createErr := simulateJSONRequest(router, createRequest, true)
	driverId := resp.Data.(map[string]interface{})["_id"].(string)
	updateRequest.Path += driverId
	w, resp, err := simulateJSONRequest(router, updateRequest, true)

	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Driver updated successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
	assert.Nil(resp.Data.(map[string]interface{})["password"])
	assert.Nil(resp.Data.(map[string]interface{})["passwordSalt"])
}

func TestDeleteDriverSuccess(t *testing.T) {
	assert := assert.New(t)

	driver := models.Driver{
		FirstName: "John",
		LastName:  "Doe",
		Rut:       "11.111.111-1",
		Phone:     "+56987654321",
		Email:     "driver@example.com",
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/drivers/",
		Body:   driver,
	}

	var deleteRequest = Request{
		Method: http.MethodDelete,
		Path:   "/drivers/",
	}

	TestSetApiRoutes(t)
	_, resp, createErr := simulateJSONRequest(router, createRequest, true)
	driverId := resp.Data.(map[string]interface{})["_id"].(string)
	deleteRequest.Path += driverId
	w, resp, err := simulateJSONRequest(router, deleteRequest, true)

	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Driver deleted successfully.", resp.Message)
	assert.Nil(resp.Data)
}

func TestListDriverSuccess(t *testing.T) {
	assert := assert.New(t)

	driver := models.Driver{
		FirstName: "John",
		LastName:  "Doe",
		Rut:       "11.111.111-1",
		Phone:     "+56987654321",
		Email:     "driver@example.com",
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/drivers/",
		Body:   driver,
	}

	var listRequest = Request{
		Method: http.MethodGet,
		Path:   "/drivers/list",
	}

	TestSetApiRoutes(t)
	_, _, createErr := simulateJSONRequest(router, createRequest, true)
	w, resp, err := simulateJSONRequest(router, listRequest, true)

	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Driver list successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType([]interface{}{}, resp.Data)
	assert.Len(resp.Data, 1)
}

func TestUpdateDriverPasswordSuccess(t *testing.T) {
	assert := assert.New(t)

	driver := models.Driver{
		FirstName: "John",
		LastName:  "Doe",
		Rut:       "11.111.111-1",
		Phone:     "+56987654321",
		Email:     "driver@example.com",
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/drivers/",
		Body:   &driver,
	}

	var updatePasswordRequest = Request{
		Method: http.MethodPatch,
		Path:   "/drivers/",
		Body: map[string]interface{}{
			"password": "password123",
		},
	}

	TestSetApiRoutes(t)
	_, resp, createErr := simulateJSONRequest(router, createRequest, true)
	driverId := resp.Data.(map[string]interface{})["_id"].(string)
	updatePasswordRequest.Path += driverId + "/password"
	w, resp, err := simulateJSONRequest(router, updatePasswordRequest, true)

	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Driver password updated successfully.", resp.Message)
	assert.Nil(resp.Data)
}

func TestUpdateDriverPasswordAndLoginSuccess(t *testing.T) {
	assert := assert.New(t)

	driver := models.Driver{
		FirstName: "John",
		LastName:  "Doe",
		Rut:       "11.111.111-1",
		Phone:     "+56987654321",
		Email:     "driver@example.com",
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/drivers/",
		Body:   &driver,
	}

	var updatePasswordRequest = Request{
		Method: http.MethodPatch,
		Path:   "/drivers/",
		Body: map[string]interface{}{
			"password": "password123",
		},
	}

	var loginRequest = Request{
		Method: http.MethodPost,
		Path:   "/drivers/login",
		Body: map[string]interface{}{
			"email":    "driver@example.com",
			"password": "password123",
		},
	}

	TestSetApiRoutes(t)
	_, resp, createErr := simulateJSONRequest(router, createRequest, true)
	driverId := resp.Data.(map[string]interface{})["_id"].(string)
	updatePasswordRequest.Path += driverId + "/password"
	_, _, _ = simulateJSONRequest(router, updatePasswordRequest, true)
	w, resp, err := simulateJSONRequest(router, loginRequest, true)

	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Driver logged successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
	assert.NotEmpty(resp.Data)
	assert.NotNil(resp.Data.(map[string]interface{})["user"])
	assert.Contains(resp.Data.(map[string]interface{})["token"], ".")
}
