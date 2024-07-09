package testing

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"payment-payments-api/internal/models"
	"testing"
)

func TestCreateDriverVehicleSuccess(t *testing.T) {
	assert := assert.New(t)

	driver := models.Driver{
		FirstName: "John",
		LastName:  "Doe",
		Rut:       "11.111.111-1",
		Phone:     "+56987654321",
		Email:     "driver@example.com",
	}

	driverVehicle := models.DriverVehicle{
		Brand:          "Tesla",
		Plate:          "HTL-TRS",
		Capacity:       10,
		DrivenDistance: 100,
	}

	var createDriverRequest = Request{
		Method: http.MethodPost,
		Path:   "/drivers/",
		Body:   driver,
	}

	var request = Request{
		Method: http.MethodPost,
		Path:   "/driver-vehicles/",
		Body:   &driverVehicle,
	}

	TestSetApiRoutes(t)
	_, resp, createDriverErr := simulateJSONRequest(router, createDriverRequest, true)
	driverId := resp.Data.(map[string]interface{})["_id"].(string)
	driverVehicle.DriverId = driverId
	w, resp, err := simulateJSONRequest(router, request, true)

	assert.Nil(createDriverErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Driver vehicle created successfully.", resp.Message, resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
}

func TestCreateDriverVehiclePlateExistsFailed(t *testing.T) {
	assert := assert.New(t)

	driver := models.Driver{
		FirstName: "John",
		LastName:  "Doe",
		Rut:       "11.111.111-1",
		Phone:     "+56987654321",
		Email:     "driver@example.com",
	}

	driverVehicle1 := models.DriverVehicle{
		Brand:          "Tesla",
		Plate:          "HTL-TRS",
		Capacity:       10,
		DrivenDistance: 100,
	}

	driverVehicle2 := models.DriverVehicle{
		Brand:          "Tesla",
		Plate:          "HTL-TRS",
		Capacity:       10,
		DrivenDistance: 100,
	}

	var createDriverRequest = Request{
		Method: http.MethodPost,
		Path:   "/drivers/",
		Body:   driver,
	}

	var request = Request{
		Method: http.MethodPost,
		Path:   "/driver-vehicles/",
	}

	TestSetApiRoutes(t)
	_, resp, createDriverErr := simulateJSONRequest(router, createDriverRequest, true)
	driverId := resp.Data.(map[string]interface{})["_id"].(string)
	driverVehicle1.DriverId = driverId
	driverVehicle2.DriverId = driverId
	_, _, _ = simulateJSONRequest(router, request, true)
	request.Body = driverVehicle1
	_, _, _ = simulateJSONRequest(router, request, true)
	request.Body = driverVehicle2
	w, resp, err := simulateJSONRequest(router, request, true)

	assert.Nil(createDriverErr)
	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Nil(err)
	assert.Equal("error", resp.Status)
	assert.Equal("Driver vehicle cannot be created.", resp.Message, resp.Data)
	assert.Equal("plate exists.", resp.Data)
}

func TestGetDriverVehicleSuccess(t *testing.T) {
	assert := assert.New(t)

	driver := models.Driver{
		FirstName: "John",
		LastName:  "Doe",
		Rut:       "11.111.111-1",
		Phone:     "+56987654321",
		Email:     "driver@example.com",
	}

	driverVehicle := models.DriverVehicle{
		Brand:          "Tesla",
		Plate:          "HTL-TRS",
		Capacity:       10,
		DrivenDistance: 100,
	}

	var createDriverRequest = Request{
		Method: http.MethodPost,
		Path:   "/drivers/",
		Body:   driver,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/driver-vehicles/",
		Body:   &driverVehicle,
	}

	var getRequest = Request{
		Method: http.MethodGet,
		Path:   "/driver-vehicles/",
	}

	TestSetApiRoutes(t)
	_, resp, createDriverErr := simulateJSONRequest(router, createDriverRequest, true)
	driverId := resp.Data.(map[string]interface{})["_id"].(string)
	driverVehicle.DriverId = driverId
	_, resp, createErr := simulateJSONRequest(router, createRequest, true)
	driverVehicleId := resp.Data.(map[string]interface{})["_id"].(string)
	getRequest.Path += driverVehicleId
	w, resp, err := simulateJSONRequest(router, getRequest, true)

	assert.Nil(createDriverErr)
	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Driver vehicle found successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
}

func TestUpdateDriverVechileSuccess(t *testing.T) {
	assert := assert.New(t)

	driver := models.Driver{
		FirstName: "John",
		LastName:  "Doe",
		Rut:       "11.111.111-1",
		Phone:     "+56987654321",
		Email:     "driver@example.com",
	}

	driverVehicle := models.DriverVehicle{
		Brand:          "Tesla",
		Plate:          "HTL-TRS",
		Capacity:       10,
		DrivenDistance: 100,
	}

	var createDriverRequest = Request{
		Method: http.MethodPost,
		Path:   "/drivers/",
		Body:   driver,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/driver-vehicles/",
		Body:   &driverVehicle,
	}

	var updateRequest = Request{
		Method: http.MethodPut,
		Path:   "/driver-vehicles/",
		Body: models.DriverVehicle{
			Brand:          "Tesla",
			Plate:          "HTL-XXX",
			Capacity:       10,
			DrivenDistance: 100,
		},
	}

	TestSetApiRoutes(t)
	_, resp, createDriverErr := simulateJSONRequest(router, createDriverRequest, true)
	driverId := resp.Data.(map[string]interface{})["_id"].(string)
	driverVehicle.DriverId = driverId
	_, resp, createErr := simulateJSONRequest(router, createRequest, true)
	driverVehicleId := resp.Data.(map[string]interface{})["_id"].(string)
	updateRequest.Path += driverVehicleId
	w, resp, err := simulateJSONRequest(router, updateRequest, true)

	assert.Nil(createDriverErr)
	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Driver vehicle updated successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
}

func TestDeleteDriverVehicleSuccess(t *testing.T) {
	assert := assert.New(t)

	driver := models.Driver{
		FirstName: "John",
		LastName:  "Doe",
		Rut:       "11.111.111-1",
		Phone:     "+56987654321",
		Email:     "driver@example.com",
	}

	driverVehicle := models.DriverVehicle{
		Brand:          "Tesla",
		Plate:          "HTL-TRS",
		Capacity:       10,
		DrivenDistance: 100,
	}

	var createDriverRequest = Request{
		Method: http.MethodPost,
		Path:   "/drivers/",
		Body:   driver,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/driver-vehicles/",
		Body:   &driverVehicle,
	}

	var deleteRequest = Request{
		Method: http.MethodDelete,
		Path:   "/driver-vehicles/",
	}

	TestSetApiRoutes(t)
	_, resp, createDriverErr := simulateJSONRequest(router, createDriverRequest, true)
	driverId := resp.Data.(map[string]interface{})["_id"].(string)
	driverVehicle.DriverId = driverId
	_, resp, createErr := simulateJSONRequest(router, createRequest, true)
	driverParkingLotId := resp.Data.(map[string]interface{})["_id"].(string)
	deleteRequest.Path += driverParkingLotId
	w, resp, err := simulateJSONRequest(router, deleteRequest, true)

	assert.Nil(createDriverErr)
	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Driver vehicle deleted successfully.", resp.Message)
	assert.Nil(resp.Data)
}

func TestListDriverVehicleSuccess(t *testing.T) {
	assert := assert.New(t)

	driver := models.Driver{
		FirstName: "John",
		LastName:  "Doe",
		Rut:       "11.111.111-1",
		Phone:     "+56987654321",
		Email:     "driver@example.com",
	}

	driverVehicle := models.DriverVehicle{
		Brand:          "Tesla",
		Plate:          "HTL-TRS",
		Capacity:       10,
		DrivenDistance: 100,
	}

	var createDriverRequest = Request{
		Method: http.MethodPost,
		Path:   "/drivers/",
		Body:   driver,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/driver-vehicles/",
		Body:   &driverVehicle,
	}

	var listRequest = Request{
		Method: http.MethodGet,
		Path:   "/driver-vehicles/list/",
	}

	TestSetApiRoutes(t)
	_, resp, createDriverErr := simulateJSONRequest(router, createDriverRequest, true)
	driverId := resp.Data.(map[string]interface{})["_id"].(string)
	driverVehicle.DriverId = driverId
	_, _, createErr := simulateJSONRequest(router, createRequest, true)
	listRequest.Path += driverId
	w, resp, err := simulateJSONRequest(router, listRequest, true)

	assert.Nil(createDriverErr)
	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Driver vehicle list successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType([]interface{}{}, resp.Data)
	assert.Len(resp.Data, 1)
}
