package testing

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
	"payment-payments-api/internal/api/util"
	"payment-payments-api/internal/models"
	"payment-payments-api/pkg/mongodb"
	"testing"
)

func TestCreateOrderSuccess(t *testing.T) {
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
				Region:          "Santiago",
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
		Name:            "NewBusinessDestination",
		Address:         "La Moneda 1080",
		AddressOptional: "Dpto. 1003",
		Municipality:    "Santiago",
		Region:          "Santiago",
		Lat:             -30.50,
		Lng:             -31.73,
		AttendantName:   "John Doe",
		AttendantPhone:  "+56987654321",
		AttendantEmail:  "businessowner@gmail.com",
	}

	municipality := models.Municipality{
		Country: models.MunicipalityCountry,
		Region:  "Santiago",
		Name:    "Santiago",
		Kind:    models.MunicipalityBase,
	}

	fee := models.Fee{
		WeightRange: models.WeightRange0610,
		Price:       2000,
	}

	order := models.Order{
		BusinessOrderId: "BusinessOrderId",
		GoodsPrice:      100,
		PackageQty:      3,
		WeightRange:     models.WeightRange0610,
		Origin:          models.BusinessOrigin{},
		Destination:     models.BusinessDestination{},
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createBusinessDestinationRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-destinations/",
		Body:   &businessDestination,
	}

	var createMunicipalityRequest = Request{
		Method: http.MethodPost,
		Path:   "/municipalities/",
		Body:   &municipality,
	}

	var createFeeRequest = Request{
		Method: http.MethodPost,
		Path:   "/fees/",
		Body:   &fee,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/orders/",
		Body:   &order,
	}

	TestSetApiRoutes(t)
	_, resp, createMunicipalityErr := simulateJSONRequest(router, createMunicipalityRequest, true)
	municipalityId := resp.Data.(map[string]interface{})["_id"].(string)
	fee.OriginMunicipalityId = municipalityId
	fee.DestinationMunicipalityId = municipalityId
	_, resp, createFeeErr := simulateJSONRequest(router, createFeeRequest, true)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	originId := resp.Data.(map[string]interface{})["origins"].([]interface{})[0].(map[string]interface{})["_id"].(string)
	order.BusinessId = businessId
	order.Origin.Id = originId
	businessDestination.BusinessId = businessId
	_, resp, createBusinessDestinationErr := simulateJSONRequest(router, createBusinessDestinationRequest, true)
	destinationId := resp.Data.(map[string]interface{})["_id"].(string)
	order.Destination.Id = destinationId
	w, resp, err := simulateJSONRequest(router, createRequest, true)

	assert.Nil(createMunicipalityErr)
	assert.Nil(createFeeErr)
	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessDestinationErr)
	assert.Equal(http.StatusOK, w.Code, resp.Data)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Order created successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
	assert.Equal(models.OrderAvailable, resp.Data.(map[string]interface{})["state"])
}

func TestCreateOrdersAndVerifySystemIDSuccess(t *testing.T) {
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
				Region:          "Santiago",
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
		Name:            "NewBusinessDestination",
		Address:         "La Moneda 1080",
		AddressOptional: "Dpto. 1003",
		Municipality:    "Santiago",
		Region:          "Santiago",
		Lat:             -30.50,
		Lng:             -31.73,
		AttendantName:   "John Doe",
		AttendantPhone:  "+56987654321",
		AttendantEmail:  "businessowner@gmail.com",
	}

	municipality := models.Municipality{
		Country: models.MunicipalityCountry,
		Region:  "Santiago",
		Name:    "Santiago",
		Kind:    models.MunicipalityBase,
	}

	fee := models.Fee{
		WeightRange: models.WeightRange0610,
		Price:       2000,
	}

	order1 := models.Order{
		BusinessOrderId: "BusinessOrderId",
		GoodsPrice:      100,
		PackageQty:      3,
		WeightRange:     models.WeightRange0610,
		Origin:          models.BusinessOrigin{},
		Destination:     models.BusinessDestination{},
	}

	order2 := models.Order{
		BusinessOrderId: "BusinessOrderId",
		GoodsPrice:      100,
		PackageQty:      3,
		WeightRange:     models.WeightRange0610,
		Origin:          models.BusinessOrigin{},
		Destination:     models.BusinessDestination{},
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createBusinessDestinationRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-destinations/",
		Body:   &businessDestination,
	}

	var createMunicipalityRequest = Request{
		Method: http.MethodPost,
		Path:   "/municipalities/",
		Body:   &municipality,
	}

	var createFeeRequest = Request{
		Method: http.MethodPost,
		Path:   "/fees/",
		Body:   &fee,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/orders/",
	}

	TestSetApiRoutes(t)
	_, resp, createMunicipalityErr := simulateJSONRequest(router, createMunicipalityRequest, true)
	municipalityId := resp.Data.(map[string]interface{})["_id"].(string)
	fee.OriginMunicipalityId = municipalityId
	fee.DestinationMunicipalityId = municipalityId
	_, resp, createFeeErr := simulateJSONRequest(router, createFeeRequest, true)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	originId := resp.Data.(map[string]interface{})["origins"].([]interface{})[0].(map[string]interface{})["_id"].(string)
	businessDestination.BusinessId = businessId
	_, resp, createBusinessDestinationErr := simulateJSONRequest(router, createBusinessDestinationRequest, true)
	destinationId := resp.Data.(map[string]interface{})["_id"].(string)
	order1.BusinessId = businessId
	order1.Origin.Id = originId
	order1.Destination.Id = destinationId
	createRequest.Body = order1
	_, _, _ = simulateJSONRequest(router, createRequest, true)
	order2.BusinessId = businessId
	order2.Origin.Id = originId
	order2.Destination.Id = destinationId
	createRequest.Body = order2
	w, resp, err := simulateJSONRequest(router, createRequest, true)

	assert.Nil(createMunicipalityErr)
	assert.Nil(createFeeErr)
	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessDestinationErr)
	assert.Equal(http.StatusOK, w.Code, resp.Data)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Order created successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
	assert.Equal("P2", resp.Data.(map[string]interface{})["systemId"])
}

func TestGetOrderSuccess(t *testing.T) {
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
				Region:          "Santiago",
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
		Name:            "NewBusinessDestination",
		Address:         "La Moneda 1080",
		AddressOptional: "Dpto. 1003",
		Municipality:    "Santiago",
		Region:          "Santiago",
		Lat:             -30.50,
		Lng:             -31.73,
		AttendantName:   "John Doe",
		AttendantPhone:  "+56987654321",
		AttendantEmail:  "businessowner@gmail.com",
	}

	municipality := models.Municipality{
		Country: models.MunicipalityCountry,
		Region:  "Santiago",
		Name:    "Santiago",
		Kind:    models.MunicipalityBase,
	}

	fee := models.Fee{
		WeightRange: models.WeightRange0610,
		Price:       2000,
	}

	order := models.Order{
		BusinessOrderId: "BusinessOrderId",
		GoodsPrice:      100,
		PackageQty:      3,
		WeightRange:     models.WeightRange0610,
		Origin:          models.BusinessOrigin{},
		Destination:     models.BusinessDestination{},
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createBusinessDestinationRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-destinations/",
		Body:   &businessDestination,
	}

	var createOrderRequest = Request{
		Method: http.MethodPost,
		Path:   "/orders/",
		Body:   &order,
	}

	var createMunicipalityRequest = Request{
		Method: http.MethodPost,
		Path:   "/municipalities/",
		Body:   &municipality,
	}

	var createFeeRequest = Request{
		Method: http.MethodPost,
		Path:   "/fees/",
		Body:   &fee,
	}

	var getRequest = Request{
		Method: http.MethodGet,
		Path:   "/orders/",
	}

	TestSetApiRoutes(t)
	_, resp, createMunicipalityErr := simulateJSONRequest(router, createMunicipalityRequest, true)
	municipalityId := resp.Data.(map[string]interface{})["_id"].(string)
	fee.OriginMunicipalityId = municipalityId
	fee.DestinationMunicipalityId = municipalityId
	_, resp, createFeeErr := simulateJSONRequest(router, createFeeRequest, true)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	originId := resp.Data.(map[string]interface{})["origins"].([]interface{})[0].(map[string]interface{})["_id"].(string)
	order.BusinessId = businessId
	order.Origin.Id = originId
	businessDestination.BusinessId = businessId
	_, resp, createBusinessDestinationErr := simulateJSONRequest(router, createBusinessDestinationRequest, true)
	destinationId := resp.Data.(map[string]interface{})["_id"].(string)
	order.Destination.Id = destinationId
	_, resp, createOrderErr := simulateJSONRequest(router, createOrderRequest, true)
	orderId := resp.Data.(map[string]interface{})["_id"].(string)
	getRequest.Path += orderId
	w, resp, err := simulateJSONRequest(router, getRequest, true)

	assert.Nil(createMunicipalityErr)
	assert.Nil(createFeeErr)
	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessDestinationErr)
	assert.Nil(createOrderErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Order found successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
}

func TestUpdateOrderSuccess(t *testing.T) {
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
				Region:          "Santiago",
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
		Name:            "NewBusinessDestination",
		Address:         "La Moneda 1080",
		AddressOptional: "Dpto. 1003",
		Municipality:    "Santiago",
		Region:          "Santiago",
		Lat:             -30.50,
		Lng:             -31.73,
		AttendantName:   "John Doe",
		AttendantPhone:  "+56987654321",
		AttendantEmail:  "businessowner@gmail.com",
	}

	municipality := models.Municipality{
		Country: models.MunicipalityCountry,
		Region:  "Santiago",
		Name:    "Santiago",
		Kind:    models.MunicipalityBase,
	}

	fee := models.Fee{
		WeightRange: models.WeightRange0610,
		Price:       2000,
	}

	order := models.Order{
		BusinessOrderId: "BusinessOrderId",
		GoodsPrice:      100,
		PackageQty:      3,
		WeightRange:     models.WeightRange0610,
		Origin:          models.BusinessOrigin{},
		Destination:     models.BusinessDestination{},
	}

	updateOrder := models.Order{
		SystemId:        "SystemId",
		BusinessOrderId: "BusinessOrderId",
		State:           models.OrderAssigned,
		GoodsPrice:      50,
		FeePrice:        150,
		PackageQty:      5,
		WeightRange:     models.WeightRange0610,
		Origin:          models.BusinessOrigin{},
		Destination:     models.BusinessDestination{},
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createBusinessDestinationRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-destinations/",
		Body:   &businessDestination,
	}

	var createOrderRequest = Request{
		Method: http.MethodPost,
		Path:   "/orders/",
		Body:   &order,
	}

	var createMunicipalityRequest = Request{
		Method: http.MethodPost,
		Path:   "/municipalities/",
		Body:   &municipality,
	}

	var createFeeRequest = Request{
		Method: http.MethodPost,
		Path:   "/fees/",
		Body:   &fee,
	}

	var updateRequest = Request{
		Method: http.MethodPut,
		Path:   "/orders/",
		Body:   &updateOrder,
	}

	TestSetApiRoutes(t)
	_, resp, createMunicipalityErr := simulateJSONRequest(router, createMunicipalityRequest, true)
	municipalityId := resp.Data.(map[string]interface{})["_id"].(string)
	fee.OriginMunicipalityId = municipalityId
	fee.DestinationMunicipalityId = municipalityId
	_, resp, createFeeErr := simulateJSONRequest(router, createFeeRequest, true)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	originId := resp.Data.(map[string]interface{})["origins"].([]interface{})[0].(map[string]interface{})["_id"].(string)
	order.BusinessId = businessId
	order.Origin.Id = originId
	businessDestination.BusinessId = businessId
	_, resp, createBusinessDestinationErr := simulateJSONRequest(router, createBusinessDestinationRequest, true)
	destinationId := resp.Data.(map[string]interface{})["_id"].(string)
	order.Destination.Id = destinationId
	_, resp, createOrderErr := simulateJSONRequest(router, createOrderRequest, true)
	orderId := resp.Data.(map[string]interface{})["_id"].(string)
	updateOrder.BusinessId = businessId
	updateOrder.Origin.Id = originId
	updateOrder.Destination.Id = destinationId
	updateRequest.Path += orderId
	w, resp, err := simulateJSONRequest(router, updateRequest, true)

	assert.Nil(createMunicipalityErr)
	assert.Nil(createFeeErr)
	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessDestinationErr)
	assert.Nil(createOrderErr)
	assert.Equal(http.StatusOK, w.Code, resp.Data)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Order updated successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
}

func TestDeleteOrderSuccess(t *testing.T) {
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
				Region:          "Santiago",
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
		Name:            "NewBusinessDestination",
		Address:         "La Moneda 1080",
		AddressOptional: "Dpto. 1003",
		Municipality:    "Santiago",
		Region:          "Santiago",
		Lat:             -30.50,
		Lng:             -31.73,
		AttendantName:   "John Doe",
		AttendantPhone:  "+56987654321",
		AttendantEmail:  "businessowner@gmail.com",
	}

	municipality := models.Municipality{
		Country: models.MunicipalityCountry,
		Region:  "Santiago",
		Name:    "Santiago",
		Kind:    models.MunicipalityBase,
	}

	fee := models.Fee{
		WeightRange: models.WeightRange0610,
		Price:       2000,
	}

	order := models.Order{
		BusinessOrderId: "BusinessOrderId",
		GoodsPrice:      100,
		PackageQty:      3,
		WeightRange:     models.WeightRange0610,
		Origin:          models.BusinessOrigin{},
		Destination:     models.BusinessDestination{},
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createBusinessDestinationRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-destinations/",
		Body:   &businessDestination,
	}

	var createOrderRequest = Request{
		Method: http.MethodPost,
		Path:   "/orders/",
		Body:   &order,
	}

	var createMunicipalityRequest = Request{
		Method: http.MethodPost,
		Path:   "/municipalities/",
		Body:   &municipality,
	}

	var createFeeRequest = Request{
		Method: http.MethodPost,
		Path:   "/fees/",
		Body:   &fee,
	}

	var deleteRequest = Request{
		Method: http.MethodDelete,
		Path:   "/orders/",
	}

	TestSetApiRoutes(t)
	_, resp, createMunicipalityErr := simulateJSONRequest(router, createMunicipalityRequest, true)
	municipalityId := resp.Data.(map[string]interface{})["_id"].(string)
	fee.OriginMunicipalityId = municipalityId
	fee.DestinationMunicipalityId = municipalityId
	_, resp, createFeeErr := simulateJSONRequest(router, createFeeRequest, true)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	originId := resp.Data.(map[string]interface{})["origins"].([]interface{})[0].(map[string]interface{})["_id"].(string)
	order.BusinessId = businessId
	order.Origin.Id = originId
	businessDestination.BusinessId = businessId
	_, resp, createBusinessDestinationErr := simulateJSONRequest(router, createBusinessDestinationRequest, true)
	destinationId := resp.Data.(map[string]interface{})["_id"].(string)
	order.Destination.Id = destinationId
	_, resp, createOrderErr := simulateJSONRequest(router, createOrderRequest, true)
	orderId := resp.Data.(map[string]interface{})["_id"].(string)
	deleteRequest.Path += orderId
	w, resp, err := simulateJSONRequest(router, deleteRequest, true)

	assert.Nil(createMunicipalityErr)
	assert.Nil(createFeeErr)
	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessDestinationErr)
	assert.Nil(createOrderErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Order cancelled successfully.", resp.Message)
	assert.Nil(resp.Data)
}

func TestListOrderSuccess(t *testing.T) {
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
				Region:          "Santiago",
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
		Name:            "NewBusinessDestination",
		Address:         "La Moneda 1080",
		AddressOptional: "Dpto. 1003",
		Municipality:    "Santiago",
		Region:          "Santiago",
		Lat:             -30.50,
		Lng:             -31.73,
		AttendantName:   "John Doe",
		AttendantPhone:  "+56987654321",
		AttendantEmail:  "businessowner@gmail.com",
	}

	municipality := models.Municipality{
		Country: models.MunicipalityCountry,
		Region:  "Santiago",
		Name:    "Santiago",
		Kind:    models.MunicipalityBase,
	}

	fee := models.Fee{
		WeightRange: models.WeightRange0610,
		Price:       2000,
	}

	order := models.Order{
		BusinessOrderId: "BusinessOrderId",
		GoodsPrice:      100,
		PackageQty:      3,
		WeightRange:     models.WeightRange0610,
		Origin:          models.BusinessOrigin{},
		Destination:     models.BusinessDestination{},
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createBusinessDestinationRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-destinations/",
		Body:   &businessDestination,
	}

	var createOrderRequest = Request{
		Method: http.MethodPost,
		Path:   "/orders/",
		Body:   &order,
	}

	var createMunicipalityRequest = Request{
		Method: http.MethodPost,
		Path:   "/municipalities/",
		Body:   &municipality,
	}

	var createFeeRequest = Request{
		Method: http.MethodPost,
		Path:   "/fees/",
		Body:   &fee,
	}

	var listRequest = Request{
		Method: http.MethodGet,
		Path:   "/orders/list/",
	}

	TestSetApiRoutes(t)
	_, resp, createMunicipalityErr := simulateJSONRequest(router, createMunicipalityRequest, true)
	municipalityId := resp.Data.(map[string]interface{})["_id"].(string)
	fee.OriginMunicipalityId = municipalityId
	fee.DestinationMunicipalityId = municipalityId
	_, resp, createFeeErr := simulateJSONRequest(router, createFeeRequest, true)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	originId := resp.Data.(map[string]interface{})["origins"].([]interface{})[0].(map[string]interface{})["_id"].(string)
	order.BusinessId = businessId
	order.Origin.Id = originId
	businessDestination.BusinessId = businessId
	_, resp, createBusinessDestinationErr := simulateJSONRequest(router, createBusinessDestinationRequest, true)
	destinationId := resp.Data.(map[string]interface{})["_id"].(string)
	order.Destination.Id = destinationId
	_, resp, createOrderErr := simulateJSONRequest(router, createOrderRequest, true)
	listRequest.Path += businessId
	w, resp, err := simulateJSONRequest(router, listRequest, true)

	assert.Nil(createMunicipalityErr)
	assert.Nil(createFeeErr)
	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessDestinationErr)
	assert.Nil(createOrderErr)
	assert.Equal(http.StatusOK, w.Code, resp.Data)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Order list successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType([]interface{}{}, resp.Data)
	assert.Len(resp.Data, 1)
}

func TestListAdminOrderSuccess(t *testing.T) {
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
				Region:          "Santiago",
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
		Name:            "NewBusinessDestination",
		Address:         "La Moneda 1080",
		AddressOptional: "Dpto. 1003",
		Municipality:    "Santiago",
		Region:          "Santiago",
		Lat:             -30.50,
		Lng:             -31.73,
		AttendantName:   "John Doe",
		AttendantPhone:  "+56987654321",
		AttendantEmail:  "businessowner@gmail.com",
	}

	municipality := models.Municipality{
		Country: models.MunicipalityCountry,
		Region:  "Santiago",
		Name:    "Santiago",
		Kind:    models.MunicipalityBase,
	}

	fee := models.Fee{
		WeightRange: models.WeightRange0610,
		Price:       2000,
	}

	order := models.Order{
		BusinessOrderId: "BusinessOrderId",
		GoodsPrice:      100,
		PackageQty:      3,
		WeightRange:     models.WeightRange0610,
		Origin:          models.BusinessOrigin{},
		Destination:     models.BusinessDestination{},
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createBusinessDestinationRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-destinations/",
		Body:   &businessDestination,
	}

	var createOrderRequest = Request{
		Method: http.MethodPost,
		Path:   "/orders/",
		Body:   &order,
	}

	var createMunicipalityRequest = Request{
		Method: http.MethodPost,
		Path:   "/municipalities/",
		Body:   &municipality,
	}

	var createFeeRequest = Request{
		Method: http.MethodPost,
		Path:   "/fees/",
		Body:   &fee,
	}

	var listRequest = Request{
		Method: http.MethodGet,
		Path:   "/orders/list",
	}

	TestSetApiRoutes(t)
	_, resp, createMunicipalityErr := simulateJSONRequest(router, createMunicipalityRequest, true)
	municipalityId := resp.Data.(map[string]interface{})["_id"].(string)
	fee.OriginMunicipalityId = municipalityId
	fee.DestinationMunicipalityId = municipalityId
	_, resp, createFeeErr := simulateJSONRequest(router, createFeeRequest, true)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	originId := resp.Data.(map[string]interface{})["origins"].([]interface{})[0].(map[string]interface{})["_id"].(string)
	order.BusinessId = businessId
	order.Origin.Id = originId
	businessDestination.BusinessId = businessId
	_, resp, createBusinessDestinationErr := simulateJSONRequest(router, createBusinessDestinationRequest, true)
	destinationId := resp.Data.(map[string]interface{})["_id"].(string)
	order.Destination.Id = destinationId
	_, resp, createOrderErr := simulateJSONRequest(router, createOrderRequest, true)
	w, resp, err := simulateJSONRequest(router, listRequest, true)

	assert.Nil(createMunicipalityErr)
	assert.Nil(createFeeErr)
	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessDestinationErr)
	assert.Nil(createOrderErr)
	assert.Equal(http.StatusOK, w.Code, resp.Data)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Order list successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType([]interface{}{}, resp.Data)
	assert.Len(resp.Data, 1)
}

func TestListOrderWeightRangesSuccess(t *testing.T) {
	assert := assert.New(t)

	var listRequest = Request{
		Method: http.MethodGet,
		Path:   "/orders/list/weight-ranges",
	}

	TestSetApiRoutes(t)
	w, resp, err := simulateJSONRequest(router, listRequest, true)

	assert.Equal(http.StatusOK, w.Code, resp.Data)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Order Weight Ranges list successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType([]interface{}{}, resp.Data)
	assert.ElementsMatch(resp.Data, models.WeightRanges)
}

func TestListOrderNextStatesSuccess(t *testing.T) {
	assert := assert.New(t)

	var listRequest = Request{
		Method: http.MethodGet,
		Path:   "/orders/list/next-states",
	}

	TestSetApiRoutes(t)
	w, resp, err := simulateJSONRequest(router, listRequest, true)

	assert.Equal(http.StatusOK, w.Code, resp.Data)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Order Next States list successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
}

func TestMassiveUploadOrderAndVerifyPricesSuccess(t *testing.T) {
	assert := assert.New(t)

	xlsOrderFile := "./assets/orders.xlsx"
	xlsFeeFile := "./assets/fees.xlsx"

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

	var origins []models.BusinessOrigin

	const orderSheet = "Pedidos"
	xlsContent, _ := os.ReadFile(xlsOrderFile)
	xlsRows, _ := api_util.GetExcelRowsFromContent(xlsContent, orderSheet)
	for i, row := range xlsRows {
		if i == 0 {
			continue
		}
		if len(row) >= 12 {
			origins = append(origins, models.BusinessOrigin{
				Name:            row[12],
				Address:         "La Moneda 970",
				AddressOptional: "Dpto. 710",
				Country:         "",
				Region:          "Región Metropolitana",
				Municipality:    "Santiago",
				Lat:             -30.10,
				Lng:             -31.22,
				AttendantName:   "John Doe",
				AttendantPhone:  "+56987654321",
				AttendantEmail:  "businessowner@gmail.com",
				OpeningTime:     "10:00",
				ClosingTime:     "16:30",
				Notes:           "",
				IsDefault:       false,
			})
		}
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createOriginRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-origins/",
	}

	var uploadFeeRequest = Request{
		Method: http.MethodPost,
		Path:   "/fees/massive-upload",
	}

	var uploadRequest = Request{
		Method: http.MethodPost,
		Path:   "/orders/massive-upload/",
	}

	TestSetApiRoutes(t)
	_, _, uploadFeeRequestErr := simulateFormDataFileRequest(router, uploadFeeRequest, xlsFeeFile, true)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	for _, origin := range origins {
		origin.BusinessId = businessId
		createOriginRequest.Body = origin
		_, _, originErr := simulateJSONRequest(router, createOriginRequest, true)
		assert.Nil(originErr)
	}
	uploadRequest.Path += businessId
	w, resp, err := simulateFormDataFileRequest(router, uploadRequest, xlsOrderFile, true)

	if arr, ok := resp.Data.([]interface{}); ok {
		for _, vArr := range arr {
			if vMap, ok := vArr.(map[string]interface{}); ok {
				for k, v := range vMap {
					if m, ok := v.(map[string]interface{}); ok {
						vStr, _ := m["value"].(string)
						assert.Nil(m["err"], fmt.Sprintf("%s : %s", k, vStr))
					}
				}
			}
		}
	}

	assert.Nil(uploadFeeRequestErr)
	assert.Nil(createBusinessErr)
	assert.Equal(http.StatusOK, w.Code, resp.Data)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Orders created successfully.", resp.Message)
	assert.Nil(resp.Data)
}

func TestCreateOrderNextStateHappyWaySuccess(t *testing.T) {
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
				Region:          "Santiago",
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
		Name:            "NewBusinessDestination",
		Address:         "La Moneda 1080",
		AddressOptional: "Dpto. 1003",
		Municipality:    "Santiago",
		Region:          "Santiago",
		Lat:             -30.50,
		Lng:             -31.73,
		AttendantName:   "John Doe",
		AttendantPhone:  "+56987654321",
		AttendantEmail:  "businessowner@gmail.com",
	}

	municipality := models.Municipality{
		Country: models.MunicipalityCountry,
		Region:  "Santiago",
		Name:    "Santiago",
		Kind:    models.MunicipalityBase,
	}

	fee := models.Fee{
		WeightRange: models.WeightRange0610,
		Price:       2000,
	}

	order := models.Order{
		BusinessOrderId: "BusinessOrderId",
		GoodsPrice:      100,
		PackageQty:      3,
		WeightRange:     models.WeightRange0610,
		Origin:          models.BusinessOrigin{},
		Destination:     models.BusinessDestination{},
	}

	driver := models.Driver{
		FirstName: "John",
		LastName:  "Doe",
		Rut:       "11.111.111-1",
		Phone:     "+56987654321",
		Email:     "driver@example.com",
	}

	workLoad, _ := models.NewWorkLoad(models.WorkLoad{})

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createBusinessDestinationRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-destinations/",
		Body:   &businessDestination,
	}

	var createMunicipalityRequest = Request{
		Method: http.MethodPost,
		Path:   "/municipalities/",
		Body:   &municipality,
	}

	var createFeeRequest = Request{
		Method: http.MethodPost,
		Path:   "/fees/",
		Body:   &fee,
	}

	var createOrderRequest = Request{
		Method: http.MethodPost,
		Path:   "/orders/",
		Body:   &order,
	}

	var createDriverRequest = Request{
		Method: http.MethodPost,
		Path:   "/drivers/",
		Body:   driver,
	}

	var nextStateRequest = Request{
		Method: http.MethodPatch,
		Path:   "/orders/next-state/",
	}

	TestSetApiRoutes(t)
	_, resp, createMunicipalityErr := simulateJSONRequest(router, createMunicipalityRequest, true)
	municipalityId := resp.Data.(map[string]interface{})["_id"].(string)
	fee.OriginMunicipalityId = municipalityId
	fee.DestinationMunicipalityId = municipalityId
	_, resp, createFeeErr := simulateJSONRequest(router, createFeeRequest, true)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	originId := resp.Data.(map[string]interface{})["origins"].([]interface{})[0].(map[string]interface{})["_id"].(string)
	order.BusinessId = businessId
	order.Origin.Id = originId
	businessDestination.BusinessId = businessId
	_, resp, createBusinessDestinationErr := simulateJSONRequest(router, createBusinessDestinationRequest, true)
	destinationId := resp.Data.(map[string]interface{})["_id"].(string)
	order.Destination.Id = destinationId
	_, resp, createOrderErr := simulateJSONRequest(router, createOrderRequest, true)
	order.Id = resp.Data.(map[string]interface{})["_id"].(string)
	workLoad.Orders = []models.WorkLoadOrder{
		{
			OrderIndex: 0,
			OrderId:    resp.Data.(map[string]interface{})["_id"].(string),
			OrderState: models.OrderAssigned,
			Lat:        resp.Data.(map[string]interface{})["origin"].(map[string]interface{})["lat"].(float64),
			Lng:        resp.Data.(map[string]interface{})["origin"].(map[string]interface{})["lng"].(float64),
		},
		{
			OrderIndex: 1,
			OrderId:    resp.Data.(map[string]interface{})["_id"].(string),
			OrderState: models.OrderAssigned,
			Lat:        resp.Data.(map[string]interface{})["destination"].(map[string]interface{})["lat"].(float64),
			Lng:        resp.Data.(map[string]interface{})["destination"].(map[string]interface{})["lng"].(float64),
		},
	}
	_, resp, createDriverErr := simulateJSONRequest(router, createDriverRequest, true)
	driverId := resp.Data.(map[string]interface{})["_id"].(string)
	workLoad.DriverId = driverId
	createWorkLoadErr := db.SaveDocStruct("work-load", workLoad)

	order.State = models.OrderAssigned
	order.WorkLoad.OriginIndex = 0
	order.WorkLoad.DestinationIndex = 1
	order.WorkLoad.WorkLoadId = workLoad.Id
	oId, _ := primitive.ObjectIDFromHex(order.Id)
	updateOrderErr := db.UpdateOneStruct("order", mongodb.Query{Selector: bson.M{"_id": oId}}, &order)
	nextStateRequest.Path += order.Id

	assert.Nil(createMunicipalityErr)
	assert.Nil(createFeeErr)
	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessDestinationErr)
	assert.Nil(createOrderErr)
	assert.Nil(createDriverErr)
	assert.Nil(createWorkLoadErr)
	assert.Nil(updateOrderErr)

	// OrderPickedUp
	nextStateRequest.Body = map[string]interface{}{
		"type":  models.OrderTypeOrigin,
		"state": models.OrderReqSuccess,
	}
	_, resp, nextStateErr := simulateJSONRequest(router, nextStateRequest, true)
	assert.Nil(nextStateErr)
	assert.Equal(models.OrderPickedUp, resp.Data.(map[string]interface{})["newState"].(string))

	oId, _ = primitive.ObjectIDFromHex(workLoad.Id)
	findWorkLoadErr := db.FindOneStruct("work-load", mongodb.Query{Selector: bson.M{"_id": oId}}, &workLoad)
	assert.Nil(findWorkLoadErr)
	assert.Equal(models.OrderPickedUp, workLoad.Orders[0].OrderState)
	assert.Equal(models.OrderAssigned, workLoad.Orders[1].OrderState)

	// OrderDelivered
	nextStateRequest.Body = map[string]interface{}{
		"type":  models.OrderTypeDestination,
		"state": models.OrderReqSuccess,
	}
	_, resp, nextStateErr = simulateJSONRequest(router, nextStateRequest, true)
	assert.Nil(nextStateErr)
	assert.Equal(models.OrderDelivered, resp.Data.(map[string]interface{})["newState"].(string))

	oId, _ = primitive.ObjectIDFromHex(workLoad.Id)
	findWorkLoadErr = db.FindOneStruct("work-load", mongodb.Query{Selector: bson.M{"_id": oId}}, &workLoad)
	assert.Nil(findWorkLoadErr)
	assert.Equal(models.OrderPickedUp, workLoad.Orders[0].OrderState)
	assert.Equal(models.OrderDelivered, workLoad.Orders[1].OrderState)
}
