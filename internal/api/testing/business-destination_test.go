package testing

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"payment-payments-api/internal/models"
	"testing"
)

func TestCreateBusinessDestinationSuccess(t *testing.T) {
	assert := assert.New(t)

	business := models.Business{
		Name:         "New Business",
		Alias:        "Business Alias",
		Rut:          "11.111.111-1",
		Category:     "Food",
		BusinessLine: "Food",
		Address:      "La Moneda 970",
		Municipality: "Santiago",
		Region:       "Región Metropolitana",
		Phone:        "+56956123456",
		Email:        "newbusiness@gmail.com",
		Owner: models.BusinessOwner{
			FirstName: "New",
			LastName:  "Owner",
		},
		Origins: []models.BusinessOrigin{
			{
				Name:            "default",
				Address:         "La Moneda 970",
				AddressOptional: "Dpto. 710",
				Municipality:    "Santiago",
				Region:          "Región Metropolitana",
				Lat:             -30.10,
				Lng:             -31.22,
				AttendantName:   "John Doe",
				AttendantPhone:  "+56987654321",
				AttendantEmail:  "businessowner@gmail.com",
				OpeningTime:     "10:00",
				ClosingTime:     "16:30",
			},
		},
	}

	businessDestination := models.BusinessDestination{
		Name:            "CreateBusinessDestination",
		Address:         "La Moneda 1080",
		AddressOptional: "Dpto. 1003",
		Municipality:    "Santiago",
		Region:          "Región Metropolitana",
		Lat:             -30.50,
		Lng:             -31.73,
		AttendantName:   "John Doe",
		AttendantPhone:  "+56987654321",
		AttendantEmail:  "businessowner@gmail.com",
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-destinations/",
		Body:   &businessDestination,
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessDestination.BusinessId = businessId
	w, resp, err := simulateJSONRequest(router, createRequest, true)

	assert.Nil(createBusinessErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business destination created successfully.", resp.Message, resp.Data)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
}

func TestCreateBusinessDestinationFailedNameExists(t *testing.T) {
	assert := assert.New(t)

	business := models.Business{
		Name:         "New Business",
		Alias:        "Business Alias",
		Rut:          "11.111.111-1",
		Category:     "Food",
		BusinessLine: "Food",
		Address:      "La Moneda 970",
		Municipality: "Santiago",
		Region:       "Región Metropolitana",
		Phone:        "+56956123456",
		Email:        "newbusiness@gmail.com",
		Owner: models.BusinessOwner{
			FirstName: "New",
			LastName:  "Owner",
		},
		Origins: []models.BusinessOrigin{
			{
				Name:            "default",
				Address:         "La Moneda 970",
				AddressOptional: "Dpto. 710",
				Municipality:    "Santiago",
				Region:          "Región Metropolitana",
				Lat:             -30.10,
				Lng:             -31.22,
				AttendantName:   "John Doe",
				AttendantPhone:  "+56987654321",
				AttendantEmail:  "businessowner@gmail.com",
				OpeningTime:     "10:00",
				ClosingTime:     "16:30",
			},
		},
	}

	businessDestination := models.BusinessDestination{
		Name:            "CreateBusinessDestination",
		Address:         "La Moneda 1080",
		AddressOptional: "Dpto. 1003",
		Municipality:    "Santiago",
		Region:          "Región Metropolitana",
		Lat:             -30.50,
		Lng:             -31.73,
		AttendantName:   "John Doe",
		AttendantPhone:  "+56987654321",
		AttendantEmail:  "businessowner@gmail.com",
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-destinations/",
		Body:   &businessDestination,
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessDestination.BusinessId = businessId
	_, _, _ = simulateJSONRequest(router, createRequest, true)
	w, resp, err := simulateJSONRequest(router, createRequest, true)

	assert.Nil(createBusinessErr)
	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Nil(err)
	assert.Equal("error", resp.Status)
	assert.Equal("Business destination cannot be created. Name exists.", resp.Message, resp.Data)
	assert.Nil(resp.Data)
}

func TestGetBusinessDestinationSuccess(t *testing.T) {
	assert := assert.New(t)

	business := models.Business{
		Name:         "New Business",
		Alias:        "Business Alias",
		Rut:          "11.111.111-1",
		Category:     "Food",
		BusinessLine: "Food",
		Address:      "La Moneda 970",
		Municipality: "Santiago",
		Region:       "Región Metropolitana",
		Phone:        "+56956123456",
		Email:        "newbusiness@gmail.com",
		Owner: models.BusinessOwner{
			FirstName: "New",
			LastName:  "Owner",
		},
		Origins: []models.BusinessOrigin{
			{
				Name:            "default",
				Address:         "La Moneda 970",
				AddressOptional: "Dpto. 710",
				Municipality:    "Santiago",
				Region:          "Región Metropolitana",
				Lat:             -30.10,
				Lng:             -31.22,
				AttendantName:   "John Doe",
				AttendantPhone:  "+56987654321",
				AttendantEmail:  "businessowner@gmail.com",
				OpeningTime:     "10:00",
				ClosingTime:     "16:30",
			},
		},
	}

	businessDestination := models.BusinessDestination{
		Name:            "GetBusinessDestination",
		Address:         "La Moneda 1080",
		AddressOptional: "Dpto. 1003",
		Municipality:    "Santiago",
		Region:          "Región Metropolitana",
		Lat:             -30.50,
		Lng:             -31.73,
		AttendantName:   "John Doe",
		AttendantPhone:  "+56987654321",
		AttendantEmail:  "businessowner@gmail.com",
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-destinations/",
		Body:   &businessDestination,
	}

	var getRequest = Request{
		Method: http.MethodGet,
		Path:   "/business-destinations/",
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessDestination.BusinessId = businessId
	_, resp, createBusinessDestinationErr := simulateJSONRequest(router, createRequest, true)
	destinationId := resp.Data.(map[string]interface{})["_id"].(string)
	getRequest.Path += destinationId
	w, resp, err := simulateJSONRequest(router, getRequest, true)

	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessDestinationErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business destination found successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
}

func TestUpdateBusinessDestinationSuccess(t *testing.T) {
	assert := assert.New(t)

	business := models.Business{
		Name:         "New Business",
		Alias:        "Business Alias",
		Rut:          "11.111.111-1",
		Category:     "Food",
		BusinessLine: "Food",
		Address:      "La Moneda 970",
		Municipality: "Santiago",
		Region:       "Región Metropolitana",
		Phone:        "+56956123456",
		Email:        "newbusiness@gmail.com",
		Owner: models.BusinessOwner{
			FirstName: "New",
			LastName:  "Owner",
		},
		Origins: []models.BusinessOrigin{
			{
				Name:            "default",
				Address:         "La Moneda 970",
				AddressOptional: "Dpto. 710",
				Municipality:    "Santiago",
				Region:          "Región Metropolitana",
				Lat:             -30.10,
				Lng:             -31.22,
				AttendantName:   "John Doe",
				AttendantPhone:  "+56987654321",
				AttendantEmail:  "businessowner@gmail.com",
				OpeningTime:     "10:00",
				ClosingTime:     "16:30",
			},
		},
	}

	businessDestination := models.BusinessDestination{
		Name:            "CreateBusinessDestination",
		Address:         "La Moneda 1080",
		AddressOptional: "Dpto. 1003",
		Municipality:    "Santiago",
		Region:          "Región Metropolitana",
		Lat:             -30.50,
		Lng:             -31.73,
		AttendantName:   "John Doe",
		AttendantPhone:  "+56987654321",
		AttendantEmail:  "businessowner@gmail.com",
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-destinations/",
		Body:   &businessDestination,
	}

	var updateRequest = Request{
		Method: http.MethodPut,
		Path:   "/business-destinations/",
		Body: models.BusinessDestination{
			Name:            "UpdateBusinessDestination",
			Address:         "La Moneda 2005",
			AddressOptional: "Dpto. 1527",
			Municipality:    "Santiago",
			Region:          "Región Metropolitana",
			Lat:             -30.50,
			Lng:             -31.73,
			AttendantName:   "John Doe",
			AttendantPhone:  "+56987654321",
			AttendantEmail:  "businessowner@gmail.com",
		},
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessDestination.BusinessId = businessId
	_, resp, createBusinessDestinationErr := simulateJSONRequest(router, createRequest, true)
	destinationId := resp.Data.(map[string]interface{})["_id"].(string)
	updateRequest.Path += destinationId
	w, resp, err := simulateJSONRequest(router, updateRequest, true)

	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessDestinationErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business destination updated successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
}

func TestUpdateBusinessDestinationFailedNameExists(t *testing.T) {
	assert := assert.New(t)

	business := models.Business{
		Name:         "New Business",
		Alias:        "Business Alias",
		Rut:          "11.111.111-1",
		Category:     "Food",
		BusinessLine: "Food",
		Address:      "La Moneda 970",
		Municipality: "Santiago",
		Region:       "Región Metropolitana",
		Phone:        "+56956123456",
		Email:        "newbusiness@gmail.com",
		Owner: models.BusinessOwner{
			FirstName: "New",
			LastName:  "Owner",
		},
		Origins: []models.BusinessOrigin{
			{
				Name:            "default",
				Address:         "La Moneda 970",
				AddressOptional: "Dpto. 710",
				Municipality:    "Santiago",
				Region:          "Región Metropolitana",
				Lat:             -30.10,
				Lng:             -31.22,
				AttendantName:   "John Doe",
				AttendantPhone:  "+56987654321",
				AttendantEmail:  "businessowner@gmail.com",
				OpeningTime:     "10:00",
				ClosingTime:     "16:30",
			},
		},
	}

	businessDestination := models.BusinessDestination{
		Name:            "CreateBusinessDestination",
		Address:         "La Moneda 1080",
		AddressOptional: "Dpto. 1003",
		Municipality:    "Santiago",
		Region:          "Región Metropolitana",
		Lat:             -30.50,
		Lng:             -31.73,
		AttendantName:   "John Doe",
		AttendantPhone:  "+56987654321",
		AttendantEmail:  "businessowner@gmail.com",
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-destinations/",
		Body:   &businessDestination,
	}

	var updateRequest = Request{
		Method: http.MethodPut,
		Path:   "/business-destinations/",
		Body: models.BusinessDestination{
			Name:            "CreateBusinessDestination",
			Address:         "La Moneda 2005",
			AddressOptional: "Dpto. 1527",
			Municipality:    "Santiago",
			Region:          "Región Metropolitana",
			Lat:             -30.50,
			Lng:             -31.73,
			AttendantName:   "John Doe",
			AttendantPhone:  "+56987654321",
			AttendantEmail:  "businessowner@gmail.com",
		},
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessDestination.BusinessId = businessId
	_, resp, createBusinessDestinationErr := simulateJSONRequest(router, createRequest, true)
	businessDestination.Name = "CreateBusinessDestination2"
	_, resp, createBusinessDestination2Err := simulateJSONRequest(router, createRequest, true)
	destinationId := resp.Data.(map[string]interface{})["_id"].(string)
	updateRequest.Path += destinationId
	w, resp, err := simulateJSONRequest(router, updateRequest, true)

	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessDestinationErr)
	assert.Nil(createBusinessDestination2Err)
	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Nil(err)
	assert.Equal("error", resp.Status)
	assert.Equal("Business destination cannot be created. Name exists.", resp.Message)
	assert.Nil(resp.Data)
}

func TestDeleteBusinessDestinationSuccess(t *testing.T) {
	assert := assert.New(t)

	business := models.Business{
		Name:         "New Business",
		Alias:        "Business Alias",
		Rut:          "11.111.111-1",
		Category:     "Food",
		BusinessLine: "Food",
		Address:      "La Moneda 970",
		Municipality: "Santiago",
		Region:       "Región Metropolitana",
		Phone:        "+56956123456",
		Email:        "newbusiness@gmail.com",
		Owner: models.BusinessOwner{
			FirstName: "New",
			LastName:  "Owner",
		},
		Origins: []models.BusinessOrigin{
			{
				Name:            "default",
				Address:         "La Moneda 970",
				AddressOptional: "Dpto. 710",
				Municipality:    "Santiago",
				Region:          "Región Metropolitana",
				Lat:             -30.10,
				Lng:             -31.22,
				AttendantName:   "John Doe",
				AttendantPhone:  "+56987654321",
				AttendantEmail:  "businessowner@gmail.com",
				OpeningTime:     "10:00",
				ClosingTime:     "16:30",
			},
		},
	}

	businessDestination := models.BusinessDestination{
		Name:            "DeleteBusinessDestination",
		Address:         "La Moneda 1080",
		AddressOptional: "Dpto. 1003",
		Municipality:    "Santiago",
		Region:          "Región Metropolitana",
		Lat:             -30.50,
		Lng:             -31.73,
		AttendantName:   "John Doe",
		AttendantPhone:  "+56987654321",
		AttendantEmail:  "businessowner@gmail.com",
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-destinations/",
		Body:   &businessDestination,
	}

	var deleteRequest = Request{
		Method: http.MethodDelete,
		Path:   "/business-destinations/",
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessDestination.BusinessId = businessId
	_, resp, createBusinessOriginErr := simulateJSONRequest(router, createRequest, true)
	destinationId := resp.Data.(map[string]interface{})["_id"].(string)
	deleteRequest.Path += destinationId
	w, resp, err := simulateJSONRequest(router, deleteRequest, true)

	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessOriginErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business destination deleted successfully.", resp.Message)
	assert.Nil(resp.Data)
}

func TestListBusinessDestinationSuccess(t *testing.T) {
	assert := assert.New(t)

	business := models.Business{
		Name:         "New Business",
		Alias:        "Business Alias",
		Rut:          "11.111.111-1",
		Category:     "Food",
		BusinessLine: "Food",
		Address:      "La Moneda 970",
		Municipality: "Santiago",
		Region:       "Región Metropolitana",
		Phone:        "+56956123456",
		Email:        "listbusiness@gmail.com",
		Owner: models.BusinessOwner{
			FirstName: "New",
			LastName:  "Owner",
		},
		Origins: []models.BusinessOrigin{
			{
				Name:            "ListBusinessDestination",
				Address:         "La Moneda 970",
				AddressOptional: "Dpto. 710",
				Municipality:    "Santiago",
				Region:          "Región Metropolitana",
				Lat:             -30.10,
				Lng:             -31.22,
				AttendantName:   "John Doe",
				AttendantPhone:  "+56987654321",
				AttendantEmail:  "businessowner@gmail.com",
				OpeningTime:     "10:00",
				ClosingTime:     "16:30",
			},
		},
	}

	businessDestination := models.BusinessDestination{
		Name:            "ListBusinessDestination",
		Address:         "La Moneda 1080",
		AddressOptional: "Dpto. 1003",
		Municipality:    "Santiago",
		Region:          "Región Metropolitana",
		Lat:             -30.50,
		Lng:             -31.73,
		AttendantName:   "John Doe",
		AttendantPhone:  "+56987654321",
		AttendantEmail:  "businessowner@gmail.com",
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-destinations/",
		Body:   &businessDestination,
	}

	var listRequest = Request{
		Method: http.MethodGet,
		Path:   "/business-destinations/list/",
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessDestination.BusinessId = businessId
	_, resp, createBusinessDestinationErr := simulateJSONRequest(router, createRequest, true)
	listRequest.Path += businessId
	w, resp, err := simulateJSONRequest(router, listRequest, true)

	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessDestinationErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business destination list successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType([]interface{}{}, resp.Data)
	assert.Len(resp.Data, 1)
}
