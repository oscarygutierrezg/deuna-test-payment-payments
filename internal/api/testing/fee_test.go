package testing

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"payment-payments-api/internal/api/util"
	"payment-payments-api/internal/models"
	"payment-payments-api/internal/services"
	"strconv"
	"testing"
)

func TestCreateFeeSuccess(t *testing.T) {
	assert := assert.New(t)

	fee := models.Fee{
		WeightRange: models.WeightRange0610,
		Price:       2000,
	}

	municipality := models.Municipality{
		Country: models.MunicipalityCountry,
		Region:  "Region Metropolitana",
		Name:    "Santiago",
		Kind:    models.MunicipalityBase,
	}

	var createMunicipalityRequest = Request{
		Method: http.MethodPost,
		Path:   "/municipalities/",
		Body:   &municipality,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/fees/",
		Body:   &fee,
	}

	TestSetApiRoutes(t)
	_, resp, createMunicipalityErr := simulateJSONRequest(router, createMunicipalityRequest, true)
	municipalityId := resp.Data.(map[string]interface{})["_id"].(string)
	fee.OriginMunicipalityId = municipalityId
	fee.DestinationMunicipalityId = municipalityId
	w, resp, err := simulateJSONRequest(router, createRequest, true)

	assert.Nil(createMunicipalityErr)
	assert.Equal(http.StatusOK, w.Code, resp.Data)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Fee created successfully.", resp.Message, resp.Data)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
}

func TestGetFeeSuccess(t *testing.T) {
	assert := assert.New(t)

	fee := models.Fee{
		WeightRange: models.WeightRange0610,
		Price:       2000,
	}

	municipality := models.Municipality{
		Country: models.MunicipalityCountry,
		Region:  "Region Metropolitana",
		Name:    "Santiago",
		Kind:    models.MunicipalityBase,
	}

	var createMunicipalityRequest = Request{
		Method: http.MethodPost,
		Path:   "/municipalities/",
		Body:   &municipality,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/fees/",
		Body:   &fee,
	}

	var getRequest = Request{
		Method: http.MethodGet,
		Path:   "/fees/",
	}

	TestSetApiRoutes(t)
	_, resp, createMunicipalityErr := simulateJSONRequest(router, createMunicipalityRequest, true)
	municipalityId := resp.Data.(map[string]interface{})["_id"].(string)
	fee.OriginMunicipalityId = municipalityId
	fee.DestinationMunicipalityId = municipalityId
	_, resp, createFeeErr := simulateJSONRequest(router, createRequest, true)
	feeId := resp.Data.(map[string]interface{})["_id"].(string)
	getRequest.Path += feeId
	w, resp, err := simulateJSONRequest(router, getRequest, true)

	assert.Nil(createMunicipalityErr)
	assert.Nil(createFeeErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Fee found successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
}

func TestUpdateFeeSuccess(t *testing.T) {
	assert := assert.New(t)

	fee := models.Fee{
		WeightRange: models.WeightRange0610,
		Price:       2000,
	}

	municipality := models.Municipality{
		Country: models.MunicipalityCountry,
		Region:  "Region Metropolitana",
		Name:    "Santiago",
		Kind:    models.MunicipalityBase,
	}

	feeUpdate := models.Fee{
		WeightRange: models.WeightRange0306,
		Price:       5000,
	}

	var createMunicipalityRequest = Request{
		Method: http.MethodPost,
		Path:   "/municipalities/",
		Body:   &municipality,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/fees/",
		Body:   &fee,
	}

	var updateRequest = Request{
		Method: http.MethodPut,
		Path:   "/fees/",
		Body:   &feeUpdate,
	}

	TestSetApiRoutes(t)
	_, resp, createMunicipalityErr := simulateJSONRequest(router, createMunicipalityRequest, true)
	municipalityId := resp.Data.(map[string]interface{})["_id"].(string)
	fee.OriginMunicipalityId = municipalityId
	fee.DestinationMunicipalityId = municipalityId
	_, resp, createFeeErr := simulateJSONRequest(router, createRequest, true)
	feeId := resp.Data.(map[string]interface{})["_id"].(string)
	feeUpdate.OriginMunicipalityId = municipalityId
	feeUpdate.DestinationMunicipalityId = municipalityId
	updateRequest.Path += feeId
	w, resp, err := simulateJSONRequest(router, updateRequest, true)

	assert.Nil(createMunicipalityErr)
	assert.Nil(createFeeErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Fee updated successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
}

func TestDeleteFeeSuccess(t *testing.T) {
	assert := assert.New(t)

	fee := models.Fee{
		WeightRange: models.WeightRange0610,
		Price:       2000,
	}

	municipality := models.Municipality{
		Country: models.MunicipalityCountry,
		Region:  "Region Metropolitana",
		Name:    "Santiago",
		Kind:    models.MunicipalityBase,
	}

	var createMunicipalityRequest = Request{
		Method: http.MethodPost,
		Path:   "/municipalities/",
		Body:   &municipality,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/fees/",
		Body:   &fee,
	}

	var deleteRequest = Request{
		Method: http.MethodDelete,
		Path:   "/fees/",
	}

	TestSetApiRoutes(t)
	_, resp, createMunicipalityErr := simulateJSONRequest(router, createMunicipalityRequest, true)
	municipalityId := resp.Data.(map[string]interface{})["_id"].(string)
	fee.OriginMunicipalityId = municipalityId
	fee.DestinationMunicipalityId = municipalityId
	_, resp, createFeeErr := simulateJSONRequest(router, createRequest, true)
	feeId := resp.Data.(map[string]interface{})["_id"].(string)
	deleteRequest.Path += feeId
	w, resp, err := simulateJSONRequest(router, deleteRequest, true)

	assert.Nil(createMunicipalityErr)
	assert.Nil(createFeeErr)
	assert.Equal(http.StatusOK, w.Code, resp.Data)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Fee deleted successfully.", resp.Message)
	assert.Nil(resp.Data)
}

func TestListFeeSuccess(t *testing.T) {
	assert := assert.New(t)

	fee := models.Fee{
		WeightRange: models.WeightRange0610,
		Price:       2000,
	}

	municipality := models.Municipality{
		Country: models.MunicipalityCountry,
		Region:  "Region Metropolitana",
		Name:    "Santiago",
		Kind:    models.MunicipalityBase,
	}

	var createMunicipalityRequest = Request{
		Method: http.MethodPost,
		Path:   "/municipalities/",
		Body:   &municipality,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/fees/",
		Body:   &fee,
	}

	var listRequest = Request{
		Method: http.MethodGet,
		Path:   "/fees/list",
	}

	TestSetApiRoutes(t)
	_, resp, createMunicipalityErr := simulateJSONRequest(router, createMunicipalityRequest, true)
	municipalityId := resp.Data.(map[string]interface{})["_id"].(string)
	fee.OriginMunicipalityId = municipalityId
	fee.DestinationMunicipalityId = municipalityId
	_, resp, createFeeErr := simulateJSONRequest(router, createRequest, true)
	w, resp, err := simulateJSONRequest(router, listRequest, true)

	assert.Nil(createMunicipalityErr)
	assert.Nil(createFeeErr)
	assert.Equal(http.StatusOK, w.Code, resp.Data)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Fee list successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType([]interface{}{}, resp.Data)
	assert.Len(resp.Data, 1)
}

func TestMassiveUploadFeeAndVerifyPricesSuccess(t *testing.T) {
	assert := assert.New(t)

	xlsFeeFile := "./assets/fees.xlsx"

	var uploadRequest = Request{
		Method: http.MethodPost,
		Path:   "/fees/massive-upload",
	}

	TestSetApiRoutes(t)
	w, resp, err := simulateFormDataFileRequest(router, uploadRequest, xlsFeeFile, true)

	assert.Equal(http.StatusOK, w.Code, resp.Data)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Fees created successfully.", resp.Message)
	assert.Nil(resp.Data)

	municipalityRepository := mongodb_repository.NewMunicipalityMongoDB(&db)
	municipalityService := services.NewMunicipalityService(municipalityRepository)
	feeRepository := mongodb_repository.NewFeeMongoDB(&db)
	feeService := services.NewFeeService(feeRepository)

	const feeSheet = "Tarifario"
	const municipalityCountry = "Chile"

	xlsContent, _ := os.ReadFile(xlsFeeFile)
	xlsRows, _ := api_util.GetExcelRowsFromContent(xlsContent, feeSheet)

	limit := 1000
	count := 0

	for i, row := range xlsRows {
		if count >= limit {
			break
		}

		if len(row) <= 6 || i == 0 {
			continue
		}

		originMunicipality := row[0]
		originRegion := row[1]
		om, originMunicipalityErr := municipalityService.GetMunicipalityByCountryRegionName(municipalityCountry, originRegion, originMunicipality)

		assert.Nil(originMunicipalityErr, row)
		assert.NotNil(om)

		destinationMunicipality := row[3]
		destinationRegion := row[4]
		dm, destinationMunicipalityErr := municipalityService.GetMunicipalityByCountryRegionName(municipalityCountry, destinationRegion, destinationMunicipality)

		assert.Nil(destinationMunicipalityErr)
		assert.NotNil(dm)

		for i := 6; i < len(row); i++ {
			var weightRange string
			switch i {
			case 6:
				weightRange = models.WeightRange0003
			case 7:
				weightRange = models.WeightRange0306
			case 8:
				weightRange = models.WeightRange0610
			case 9:
				weightRange = models.WeightRange1016
			}
			feePrice, _ := strconv.Atoi(row[i])
			fee, feeErr := feeService.GetFeeByOriginDestinationWeightRange(om.Id, dm.Id, weightRange)

			assert.Nil(feeErr)
			assert.Equal(float64(feePrice), fee.Price)
		}

		count++
	}
}
