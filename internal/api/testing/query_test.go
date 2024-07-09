package testing

import (
	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
	"net/http"
	"payment-payments-api/internal/models"
	"payment-payments-api/pkg/optimizer"
	"testing"
)

func TestQuerySuccess(t *testing.T) {
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

	queryBody := models.Query{
		Orders: []optimizer.Order{
			{
				Id: "1",
				Pickup: optimizer.Location{
					Lat: -33.54546,
					Lng: -70.62984,
				},
				DropOff: optimizer.Location{
					Lat: -33.429047,
					Lng: -70.793274,
				},
				Weight: 9,
			},
			{
				Id: "2",
				Pickup: optimizer.Location{
					Lat: -33.466366,
					Lng: -70.627884,
				},
				DropOff: optimizer.Location{
					Lat: -33.43819,
					Lng: -70.79535,
				},
				Weight: 5,
			},
			{
				Id: "3",
				Pickup: optimizer.Location{
					Lat: -33.578835,
					Lng: -70.70188,
					TimeWindow: optimizer.TimeWindow{
						Starts: 2,
						Ends:   3,
					},
				},
				DropOff: optimizer.Location{
					Lat: -33.429047,
					Lng: -70.793274,
				},
				Weight: 3,
			},
			{
				Id: "4",
				Pickup: optimizer.Location{
					Lat: -33.35267,
					Lng: -70.51852,
				},
				DropOff: optimizer.Location{
					Lat: -33.409733,
					Lng: -70.5446,
					TimeWindow: optimizer.TimeWindow{
						Starts: 4,
						Ends:   8,
					},
				},
				Weight: 1,
			},
		},
		Warehouses: []optimizer.Warehouse{
			{
				Id:  "CD - Renca",
				Lat: -33.3977345,
				Lng: -70.770331,
			},
		},
		Vehicles: []optimizer.Vehicle{
			{
				Id:  "BMWD-90",
				Lat: -33.3977345,
				Lng: -70.770331,
			},
			{
				Id:  "IR-632",
				Lat: -33.3977345,
				Lng: -70.770331,
			},
			{
				Id:  "16.941.768-k",
				Lat: -33.3977345,
				Lng: -70.770331,
			},
		},
		Settings: optimizer.Settings{
			ProcessingTime:  5,
			MaxWorkingHours: 8,
			AverageStandby:  5,
			AverageVelocity: 20,
		},
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createBusinessApiKeyRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-api-keys/",
		Body:   &businessApiKey,
	}

	var queryRequest = Request{
		Method: http.MethodPost,
		Path:   "/query/",
		Body:   queryBody,
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessApiKey.BusinessId = businessId
	_, resp, createBusinessApiKeyErr := simulateJSONRequest(router, createBusinessApiKeyRequest, true)
	apiKey := resp.Data.(map[string]interface{})["apiKey"].(string)
	queryRequest.Path += "?apikey=" + apiKey
	_, resp, queryErr := simulateJSONRequest(router, queryRequest, false)

	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessApiKeyErr)
	assert.Nil(queryErr)
	assert.NotEmpty(apiKey)

	_, _ = pretty.Println(resp)
}
