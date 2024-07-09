package testing

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"payment-payments-api/internal/models"
	"testing"
)

func TestCreateBusinessOriginSuccess(t *testing.T) {
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

	businessOrigin := models.BusinessOrigin{
		Name:            "CreateBusinessOrigin",
		Address:         "La Moneda 1080",
		AddressOptional: "Dpto. 1003",
		Municipality:    "Santiago",
		Region:          "Región Metropolitana",
		Lat:             -30.50,
		Lng:             -31.73,
		AttendantName:   "John Doe",
		AttendantPhone:  "+56987654321",
		AttendantEmail:  "businessowner@gmail.com",
		OpeningTime:     "10:00",
		ClosingTime:     "16:30",
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-origins/",
		Body:   &businessOrigin,
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessOrigin.BusinessId = businessId
	w, resp, err := simulateJSONRequest(router, createRequest, true)

	assert.Nil(createBusinessErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business origin created successfully.", resp.Message, resp.Data)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
}

func TestCreateBusinessOriginFailedNameExists(t *testing.T) {
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

	businessOrigin := models.BusinessOrigin{
		Name:            "CreateBusinessOrigin",
		Address:         "La Moneda 1080",
		AddressOptional: "Dpto. 1003",
		Municipality:    "Santiago",
		Region:          "Región Metropolitana",
		Lat:             -30.50,
		Lng:             -31.73,
		AttendantName:   "John Doe",
		AttendantPhone:  "+56987654321",
		AttendantEmail:  "businessowner@gmail.com",
		OpeningTime:     "10:00",
		ClosingTime:     "16:30",
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-origins/",
		Body:   &businessOrigin,
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessOrigin.BusinessId = businessId
	_, _, _ = simulateJSONRequest(router, createRequest, true)
	w, resp, err := simulateJSONRequest(router, createRequest, true)

	assert.Nil(createBusinessErr)
	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Nil(err)
	assert.Equal("error", resp.Status)
	assert.Equal("Business origin cannot be created. Name exists.", resp.Message, resp.Data)
	assert.Nil(resp.Data)
}

func TestGetBusinessOriginSuccess(t *testing.T) {
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

	businessOrigin := models.BusinessOrigin{
		Name:            "GetBusinessOrigin",
		Address:         "La Moneda 1080",
		AddressOptional: "Dpto. 1003",
		Municipality:    "Santiago",
		Region:          "Región Metropolitana",
		Lat:             -30.50,
		Lng:             -31.73,
		AttendantName:   "John Doe",
		AttendantPhone:  "+56987654321",
		AttendantEmail:  "businessowner@gmail.com",
		OpeningTime:     "10:00",
		ClosingTime:     "16:30",
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-origins/",
		Body:   &businessOrigin,
	}

	var getRequest = Request{
		Method: http.MethodGet,
		Path:   "/business-origins/",
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessOrigin.BusinessId = businessId
	_, resp, createBusinessOriginErr := simulateJSONRequest(router, createRequest, true)
	originId := resp.Data.(map[string]interface{})["_id"].(string)
	getRequest.Path += originId
	w, resp, err := simulateJSONRequest(router, getRequest, true)

	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessOriginErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business origin found successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
}

func TestUpdateBusinessOriginSuccess(t *testing.T) {
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

	businessOrigin := models.BusinessOrigin{
		Name:            "CreateBusinessOrigin",
		Address:         "La Moneda 1080",
		AddressOptional: "Dpto. 1003",
		Municipality:    "Santiago",
		Region:          "Región Metropolitana",
		Lat:             -30.50,
		Lng:             -31.73,
		AttendantName:   "John Doe",
		AttendantPhone:  "+56987654321",
		AttendantEmail:  "businessowner@gmail.com",
		OpeningTime:     "10:00",
		ClosingTime:     "16:30",
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-origins/",
		Body:   &businessOrigin,
	}

	var updateRequest = Request{
		Method: http.MethodPut,
		Path:   "/business-origins/",
		Body: models.BusinessOrigin{
			Name:            "UpdateBusinessOrigin",
			Address:         "La Moneda 2005",
			AddressOptional: "Dpto. 1527",
			Municipality:    "Santiago",
			Region:          "Región Metropolitana",
			Lat:             -30.50,
			Lng:             -31.73,
			AttendantName:   "John Doe",
			AttendantPhone:  "+56987654321",
			AttendantEmail:  "businessowner@gmail.com",
			OpeningTime:     "08:00",
			ClosingTime:     "21:30",
		},
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessOrigin.BusinessId = businessId
	_, resp, createBusinessOriginErr := simulateJSONRequest(router, createRequest, true)
	originId := resp.Data.(map[string]interface{})["_id"].(string)
	updateRequest.Path += originId
	w, resp, err := simulateJSONRequest(router, updateRequest, true)

	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessOriginErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business origin updated successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
}

func TestUpdateBusinessOriginFailedNameExists(t *testing.T) {
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

	businessOrigin := models.BusinessOrigin{
		Name:            "CreateBusinessOrigin",
		Address:         "La Moneda 1080",
		AddressOptional: "Dpto. 1003",
		Municipality:    "Santiago",
		Region:          "Región Metropolitana",
		Lat:             -30.50,
		Lng:             -31.73,
		AttendantName:   "John Doe",
		AttendantPhone:  "+56987654321",
		AttendantEmail:  "businessowner@gmail.com",
		OpeningTime:     "10:00",
		ClosingTime:     "16:30",
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-origins/",
		Body:   &businessOrigin,
	}

	var updateRequest = Request{
		Method: http.MethodPut,
		Path:   "/business-origins/",
		Body: models.BusinessOrigin{
			Name:            "CreateBusinessOrigin",
			Address:         "La Moneda 2005",
			AddressOptional: "Dpto. 1527",
			Municipality:    "Santiago",
			Region:          "Región Metropolitana",
			Lat:             -30.50,
			Lng:             -31.73,
			AttendantName:   "John Doe",
			AttendantPhone:  "+56987654321",
			AttendantEmail:  "businessowner@gmail.com",
			OpeningTime:     "08:00",
			ClosingTime:     "21:30",
		},
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessOrigin.BusinessId = businessId
	_, resp, createBusinessOriginErr := simulateJSONRequest(router, createRequest, true)
	businessOrigin.Name = "CreateBusinessOrigin2"
	_, resp, createBusinessOrigin2Err := simulateJSONRequest(router, createRequest, true)
	originId := resp.Data.(map[string]interface{})["_id"].(string)
	updateRequest.Path += originId
	w, resp, err := simulateJSONRequest(router, updateRequest, true)

	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessOriginErr)
	assert.Nil(createBusinessOrigin2Err)
	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Nil(err)
	assert.Equal("error", resp.Status)
	assert.Equal("Business origin cannot be created. Name exists.", resp.Message)
	assert.Nil(resp.Data)
}

func TestDeleteBusinessOriginSuccess(t *testing.T) {
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

	businessOrigin := models.BusinessOrigin{
		Name:            "DeleteBusinessOrigin",
		Address:         "La Moneda 1080",
		AddressOptional: "Dpto. 1003",
		Municipality:    "Santiago",
		Region:          "Región Metropolitana",
		Lat:             -30.50,
		Lng:             -31.73,
		AttendantName:   "John Doe",
		AttendantPhone:  "+56987654321",
		AttendantEmail:  "businessowner@gmail.com",
		OpeningTime:     "10:00",
		ClosingTime:     "16:30",
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-origins/",
		Body:   &businessOrigin,
	}

	var deleteRequest = Request{
		Method: http.MethodDelete,
		Path:   "/business-origins/",
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessOrigin.BusinessId = businessId
	_, resp, createBusinessOriginErr := simulateJSONRequest(router, createRequest, true)
	originId := resp.Data.(map[string]interface{})["_id"].(string)
	deleteRequest.Path += originId
	w, resp, err := simulateJSONRequest(router, deleteRequest, true)

	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessOriginErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business origin deleted successfully.", resp.Message)
	assert.Nil(resp.Data)
}

func TestDeleteBusinessOriginFailed(t *testing.T) {
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

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var deleteRequest = Request{
		Method: http.MethodDelete,
		Path:   "/business-origins/",
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	originId := resp.Data.(map[string]interface{})["origins"].([]interface{})[0].(map[string]interface{})["_id"].(string)
	deleteRequest.Path += originId
	w, resp, err := simulateJSONRequest(router, deleteRequest, true)

	assert.Nil(createBusinessErr)
	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Nil(err)
	assert.Equal("error", resp.Status)
	assert.Equal("Business origin cannot be deleted.", resp.Message)
	assert.Equal("Business needs at least one business origin.", resp.Data)
}

func TestListBusinessOriginSuccess(t *testing.T) {
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
				Name:            "ListBusinessOrigin",
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

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var listRequest = Request{
		Method: http.MethodGet,
		Path:   "/business-origins/list/",
	}

	TestSetApiRoutes(t)
	_, resp, createErr := simulateJSONRequest(router, createRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	listRequest.Path += businessId
	w, resp, err := simulateJSONRequest(router, listRequest, true)

	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business origin list successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType([]interface{}{}, resp.Data)
	assert.Len(resp.Data, 1)
}
