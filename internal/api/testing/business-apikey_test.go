package testing

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"payment-payments-api/internal/models"
	"testing"
)

func TestCreateBusinessApiKeySuccess(t *testing.T) {
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

	businessApiKey := models.BusinessApiKey{
		BusinessId: "",
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-api-keys/",
		Body:   &businessApiKey,
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessApiKey.BusinessId = businessId
	w, resp, err := simulateJSONRequest(router, createRequest, true)

	assert.Nil(createBusinessErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business apikey created successfully.", resp.Message, resp.Data)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
}

func TestUpdateBusinessApiKeySuccess(t *testing.T) {
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

	businessApiKey := models.BusinessApiKey{
		Enabled: false,
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-api-keys/",
		Body:   &businessApiKey,
	}

	var updateRequest = Request{
		Method: http.MethodPut,
		Path:   "/business-api-keys/",
		Body:   &businessApiKey,
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessApiKey.BusinessId = businessId
	_, resp, createRequestErr := simulateJSONRequest(router, createRequest, true)
	businessApiKeyId := resp.Data.(map[string]interface{})["_id"].(string)
	updateRequest.Path += businessApiKeyId
	w, resp, err := simulateJSONRequest(router, updateRequest, true)

	assert.Nil(createBusinessErr)
	assert.Nil(createRequestErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business apikey updated successfully.", resp.Message, resp.Data)
	assert.Nil(resp.Data)
}

func TestDeleteBusinessApiKeySuccess(t *testing.T) {
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

	businessApiKey := models.BusinessApiKey{
		Enabled: false,
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-api-keys/",
		Body:   &businessApiKey,
	}

	var deleteRequest = Request{
		Method: http.MethodDelete,
		Path:   "/business-api-keys/",
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessApiKey.BusinessId = businessId
	_, resp, createRequestErr := simulateJSONRequest(router, createRequest, true)
	businessApiKeyId := resp.Data.(map[string]interface{})["_id"].(string)
	deleteRequest.Path += businessApiKeyId
	w, resp, err := simulateJSONRequest(router, deleteRequest, true)

	assert.Nil(createBusinessErr)
	assert.Nil(createRequestErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business apikey deleted successfully.", resp.Message, resp.Data)
	assert.Nil(resp.Data)
}

func TestListBusinessApiKeySuccess(t *testing.T) {
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

	businessApiKey := models.BusinessApiKey{
		Enabled: false,
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-api-keys/",
		Body:   &businessApiKey,
	}

	var listRequest = Request{
		Method: http.MethodGet,
		Path:   "/business-api-keys/list/",
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessApiKey.BusinessId = businessId
	_, resp, createRequestErr := simulateJSONRequest(router, createRequest, true)
	listRequest.Path += businessId
	w, resp, err := simulateJSONRequest(router, listRequest, true)

	assert.Nil(createBusinessErr)
	assert.Nil(createRequestErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business apikey list successfully.", resp.Message, resp.Data)
	assert.NotNil(resp.Data)
	assert.IsType([]interface{}{}, resp.Data)
	assert.Len(resp.Data, 1)
}

func TestListBusinessApiKeySuccessListWeightRanges(t *testing.T) {
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

	businessApiKey := models.BusinessApiKey{
		Enabled: false,
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-api-keys/",
		Body:   &businessApiKey,
	}

	var listRequest = Request{
		Method: http.MethodGet,
		Path:   "/business-api-keys/list/",
	}

	var listWeightRangesRequest = Request{
		Method: http.MethodGet,
		Path:   "/orders/list/weight-ranges",
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessApiKey.BusinessId = businessId
	_, resp, createRequestErr := simulateJSONRequest(router, createRequest, true)
	listRequest.Path += businessId
	_, resp, listRequestErr := simulateJSONRequest(router, listRequest, true)
	apiKey := resp.Data.([]interface{})[0].(map[string]interface{})["apiKey"].(string)
	listWeightRangesRequest.Path += "?apikey=" + apiKey
	w, resp, err := simulateJSONRequest(router, listWeightRangesRequest, true)

	assert.Nil(createBusinessErr)
	assert.Nil(createRequestErr)
	assert.Nil(listRequestErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Order Weight Ranges list successfully.", resp.Message, resp.Data)
	assert.NotNil(resp.Data)
	assert.IsType([]interface{}{}, resp.Data)
	assert.Len(resp.Data, len(models.WeightRanges))
}

func TestListBusinessApiKeyFailListWeightRanges(t *testing.T) {
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

	businessApiKey := models.BusinessApiKey{
		Enabled: false,
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-api-keys/",
		Body:   &businessApiKey,
	}

	var listRequest = Request{
		Method: http.MethodGet,
		Path:   "/business-api-keys/list/",
	}

	var listWeightRangesRequest = Request{
		Method: http.MethodGet,
		Path:   "/orders/list/weight-ranges",
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessApiKey.BusinessId = businessId
	_, resp, createRequestErr := simulateJSONRequest(router, createRequest, true)
	listRequest.Path += businessId
	_, resp, listRequestErr := simulateJSONRequest(router, listRequest, true)
	apiKey := resp.Data.([]interface{})[0].(map[string]interface{})["apiKey"].(string)
	listWeightRangesRequest.Path += "?apikey=" + apiKey + "wrongStr"
	w, resp, err := simulateJSONRequest(router, listWeightRangesRequest, false)

	assert.Nil(createBusinessErr)
	assert.Nil(createRequestErr)
	assert.Nil(listRequestErr)
	assert.Equal(http.StatusUnauthorized, w.Code)
	assert.Nil(err)
	assert.Equal("error", resp.Status)
	assert.Equal("Unauthorized", resp.Message, resp.Data)
	assert.Nil(resp.Data)
}
