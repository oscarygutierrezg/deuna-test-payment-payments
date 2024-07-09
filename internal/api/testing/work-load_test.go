package testing

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"payment-payments-api/internal/models"
	"payment-payments-api/pkg/mongodb"
	"testing"
)

func TestCreateWorkLoadSuccess(t *testing.T) {
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

	assert.Nil(createMunicipalityErr)
	assert.Nil(createFeeErr)
	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessDestinationErr)
	assert.Nil(createOrderErr)
	assert.Nil(createDriverErr)
	assert.Nil(createWorkLoadErr)
	assert.Nil(updateOrderErr)
}

func TestGetWorkLoadSuccess(t *testing.T) {
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

	var getRequest = Request{
		Method: http.MethodGet,
		Path:   "/work-loads/",
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
	getRequest.Path += workLoad.Id
	w, resp, err := simulateJSONRequest(router, getRequest, true)

	assert.Nil(createMunicipalityErr)
	assert.Nil(createFeeErr)
	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessDestinationErr)
	assert.Nil(createOrderErr)
	assert.Nil(createDriverErr)
	assert.Nil(createWorkLoadErr)
	assert.Nil(updateOrderErr)
	assert.Equal(http.StatusOK, w.Code, resp.Data)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("WorkLoad found successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
}

func TestListWorkLoadSuccess(t *testing.T) {
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

	var getWorkLoadRequest = Request{
		Method: http.MethodGet,
		Path:   "/work-loads/",
	}

	var listRequest = Request{
		Method: http.MethodGet,
		Path:   "/work-loads/list",
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
	getWorkLoadRequest.Path += workLoad.Id
	_, resp, getWorkLoadErr := simulateJSONRequest(router, getWorkLoadRequest, true)
	w, resp, err := simulateJSONRequest(router, listRequest, true)

	assert.Nil(createMunicipalityErr)
	assert.Nil(createFeeErr)
	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessDestinationErr)
	assert.Nil(createOrderErr)
	assert.Nil(createDriverErr)
	assert.Nil(createWorkLoadErr)
	assert.Nil(updateOrderErr)
	assert.Nil(getWorkLoadErr)
	assert.Equal(http.StatusOK, w.Code, resp.Data)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("WorkLoad list successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType([]interface{}{}, resp.Data)
	assert.Len(resp.Data, 1)
}

func TestSetDriverWorkLoadSuccess(t *testing.T) {
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

	var getRequest = Request{
		Method: http.MethodGet,
		Path:   "/work-loads/",
	}

	var setDriverRequest = Request{
		Method: http.MethodPatch,
		Path:   "/work-loads/",
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
	createWorkLoadErr := db.SaveDocStruct("work-load", workLoad)
	order.State = models.OrderAssigned
	order.WorkLoad.OriginIndex = 0
	order.WorkLoad.DestinationIndex = 1
	order.WorkLoad.WorkLoadId = workLoad.Id
	oId, _ := primitive.ObjectIDFromHex(order.Id)
	updateOrderErr := db.UpdateOneStruct("order", mongodb.Query{Selector: bson.M{"_id": oId}}, &order)
	setDriverRequest.Path += workLoad.Id + "/set-driver/" + driverId
	_, setDriverResp, setDriverRequestErr := simulateJSONRequest(router, setDriverRequest, true)
	getRequest.Path += workLoad.Id
	w, getRequestResp, err := simulateJSONRequest(router, getRequest, true)

	assert.Nil(createMunicipalityErr)
	assert.Nil(createFeeErr)
	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessDestinationErr)
	assert.Nil(createOrderErr)
	assert.Nil(createDriverErr)
	assert.Nil(createWorkLoadErr)
	assert.Nil(updateOrderErr)
	assert.Nil(setDriverRequestErr)
	assert.Nil(err)
	assert.Equal(http.StatusOK, w.Code, getRequestResp.Data)
	assert.Equal("success", setDriverResp.Status)
	assert.Equal("Driver assigned to WorkLoad successfully.", setDriverResp.Message)
	assert.NotNil(setDriverResp.Data)
	assert.IsType(map[string]interface{}{}, setDriverResp.Data)
	assert.Equal("success", getRequestResp.Status)
	assert.Equal("WorkLoad found successfully.", getRequestResp.Message)
	assert.NotNil(getRequestResp.Data)
	assert.IsType(map[string]interface{}{}, getRequestResp.Data)
	assert.Equal(driverId, getRequestResp.Data.(map[string]interface{})["driverId"])
}
