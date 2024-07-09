package testing

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"payment-payments-api/internal/models"
	"testing"
)

func TestCreateBusinessSuccess(t *testing.T) {
	assert := assert.New(t)

	business := models.Business{
		Name:         "New Business",
		Alias:        "Business Alias",
		Rut:          "11.111.111-1",
		Category:     "Food",
		BusinessLine: "Food",
		Address:      "La Moneda 970",
		Region:       "Región Metropolitana",
		Municipality: "Santiago",
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
				Region:          "Región Metropolitana",
				Municipality:    "Santiago",
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

	var request = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	TestSetApiRoutes(t)
	w, resp, err := simulateJSONRequest(router, request, true)

	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business created successfully.", resp.Message, resp.Data)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
	assert.IsType([]interface{}{}, resp.Data.(map[string]interface{})["origins"])
	assert.Len(resp.Data.(map[string]interface{})["origins"], 1)
	assert.Nil(resp.Data.(map[string]interface{})["owner"].(map[string]interface{})["password"])
	assert.Nil(resp.Data.(map[string]interface{})["owner"].(map[string]interface{})["passwordSalt"])
}

func TestCreateBusinessEmailExists(t *testing.T) {
	assert := assert.New(t)

	business := models.Business{
		Name:         "New Business",
		Alias:        "Business Alias",
		Rut:          "11.111.111-1",
		Category:     "Food",
		BusinessLine: "Food",
		Address:      "La Moneda 970",
		Municipality: "Santiago",
		Region:       "Santiago",
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

	var request = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	TestSetApiRoutes(t)
	_, _, _ = simulateJSONRequest(router, request, true)
	w, resp, err := simulateJSONRequest(router, request, true)

	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Nil(err)
	assert.Equal("error", resp.Status)
	assert.Equal("Business email already exists.", resp.Message, resp.Data)
	assert.Nil(resp.Data)
}

func TestGetBusinessSuccess(t *testing.T) {
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
		Email:        "getbusiness@gmail.com",
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

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var getRequest = Request{
		Method: http.MethodGet,
		Path:   "/businesses/",
	}

	TestSetApiRoutes(t)
	_, resp, createErr := simulateJSONRequest(router, createRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	getRequest.Path += businessId
	w, resp, err := simulateJSONRequest(router, getRequest, true)

	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business found successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data.(map[string]interface{})["owner"])
	assert.Nil(resp.Data.(map[string]interface{})["owner"].(map[string]interface{})["password"])
	assert.Nil(resp.Data.(map[string]interface{})["owner"].(map[string]interface{})["passwordSalt"])
}

func TestGetBusinessFail(t *testing.T) {
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
		Email:        "getbusinessfail@gmail.com",
		Owner: models.BusinessOwner{
			FirstName: "New",
			LastName:  "Owner",
		},
		Origins: []models.BusinessOrigin{
			{
				Address:         "La Moneda 970",
				AddressOptional: "Dpto. 710",
				Municipality:    "Santiago",
				Region:          "Región Metropolitana",
				Lat:             -30.10,
				Lng:             -31.22,
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

	var getRequest = Request{
		Method: http.MethodGet,
		Path:   "/businesses/",
	}

	TestSetApiRoutes(t)
	_, _, createErr := simulateJSONRequest(router, createRequest, true)
	getRequest.Path += uuid.New().String()
	w, resp, err := simulateJSONRequest(router, getRequest, true)

	assert.Nil(createErr)
	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Nil(err)
	assert.Equal("error", resp.Status)
	assert.Equal("Business not found.", resp.Message)
	assert.Nil(resp.Data)
}

func TestUpdateBusinessSuccess(t *testing.T) {
	assert := assert.New(t)

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body: models.Business{
			Name:         "New Business",
			Alias:        "Business Alias",
			Rut:          "11.111.111-1",
			Category:     "Food",
			BusinessLine: "Food",
			Address:      "La Moneda 970",
			Municipality: "Santiago",
			Region:       "Región Metropolitana",
			Phone:        "+56956123456",
			Email:        "updatebusiness@gmail.com",
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
		},
	}

	var updateRequest = Request{
		Method: http.MethodPut,
		Path:   "/businesses/",
		Body: models.Business{
			Name:         "Update Business",
			Alias:        "Update Alias",
			Rut:          "11.111.111-1",
			Category:     "Business Category",
			BusinessLine: "Business Line",
			Address:      "Other Address",
			Municipality: "Santiago",
			Region:       "Región Metropolitana",
			Phone:        "+56987654321",
			Email:        "updatebusiness@gmail.com",
			Owner: models.BusinessOwner{
				FirstName: "Other FirstName",
				LastName:  "Other LastName",
			},
			Enabled: false,
		},
	}

	TestSetApiRoutes(t)
	_, resp, createErr := simulateJSONRequest(router, createRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	updateRequest.Path += businessId
	w, resp, err := simulateJSONRequest(router, updateRequest, true)

	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business updated successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data.(map[string]interface{})["owner"])
	assert.Nil(resp.Data.(map[string]interface{})["owner"].(map[string]interface{})["password"])
	assert.Nil(resp.Data.(map[string]interface{})["owner"].(map[string]interface{})["passwordSalt"])
}

func TestDeleteBusinessSuccess(t *testing.T) {
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
		Email:        "deletebusiness@gmail.com",
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

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var deleteRequest = Request{
		Method: http.MethodDelete,
		Path:   "/businesses/",
	}

	TestSetApiRoutes(t)
	_, resp, createErr := simulateJSONRequest(router, createRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	deleteRequest.Path += businessId
	w, resp, err := simulateJSONRequest(router, deleteRequest, true)

	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business deleted successfully.", resp.Message)
	assert.Nil(resp.Data)
}

func TestListBusinessSuccess(t *testing.T) {
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

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var listRequest = Request{
		Method: http.MethodGet,
		Path:   "/businesses/list",
	}

	TestSetApiRoutes(t)
	_, _, createErr := simulateJSONRequest(router, createRequest, true)
	w, resp, err := simulateJSONRequest(router, listRequest, true)

	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business list successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType([]interface{}{}, resp.Data)
	assert.Len(resp.Data, 1)
}

func TestUpdateBusinessOwnerPasswordSuccess(t *testing.T) {
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
		Email:        "ownerbusiness@gmail.com",
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
	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}
	var updatePasswordRequest = Request{
		Method: http.MethodPatch,
		Path:   "/businesses/",
		Body: map[string]interface{}{
			"password": "password123",
		},
	}

	TestSetApiRoutes(t)
	_, resp, createErr := simulateJSONRequest(router, createRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	updatePasswordRequest.Path += businessId + "/password"
	w, resp, err := simulateJSONRequest(router, updatePasswordRequest, true)

	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business owner password updated successfully.", resp.Message)
	assert.Nil(resp.Data)
}

func TestUpdateBusinessOwnerPasswordAndLoginSuccess(t *testing.T) {
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
		Email:        "updateownerpassword@gmail.com",
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
	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}
	var updatePasswordRequest = Request{
		Method: http.MethodPatch,
		Path:   "/businesses/",
		Body: map[string]interface{}{
			"password": "password123",
		},
	}
	var loginRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/login",
		Body: map[string]interface{}{
			"email":    "updateownerpassword@gmail.com",
			"password": "password123",
		},
	}

	TestSetApiRoutes(t)
	_, resp, createErr := simulateJSONRequest(router, createRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	updatePasswordRequest.Path += businessId + "/password"
	_, resp, updatePasswordErr := simulateJSONRequest(router, updatePasswordRequest, true)
	w, resp, err := simulateJSONRequest(router, loginRequest, true)

	assert.Nil(createErr)
	assert.Nil(updatePasswordErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business owner logged successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
	assert.NotEmpty(resp.Data)
	assert.NotNil(resp.Data.(map[string]interface{})["user"])
	assert.Contains(resp.Data.(map[string]interface{})["token"], ".")
}

func TestUpdateBusinessLogoSuccess(t *testing.T) {
	assert := assert.New(t)

	logoFile := "./assets/logo.png"

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
		Email:        "ownerbusiness@gmail.com",
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

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var updateRequest = Request{
		Method: http.MethodPatch,
		Path:   "/businesses/",
	}

	TestSetApiRoutes(t)
	_, resp, createErr := simulateJSONRequest(router, createRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	updateRequest.Path += businessId + "/logo"

	w, resp, err := simulateFormDataFileRequest(router, updateRequest, logoFile, true)

	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business logo updated successfully.", resp.Message)
	assert.Nil(resp.Data)
}
