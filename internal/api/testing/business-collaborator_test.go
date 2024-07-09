package testing

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"payment-payments-api/internal/models"
	"testing"
)

func TestCreateBusinessCollaboratorSuccess(t *testing.T) {
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

	businessCollaborator := models.BusinessCollaborator{
		FirstName: "New",
		LastName:  "Collaborator",
		Email:     "newcollaborator@gmail.com",
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-collaborators/",
		Body:   &businessCollaborator,
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessCollaborator.BusinessId = businessId
	w, resp, err := simulateJSONRequest(router, createRequest, true)

	assert.Nil(createBusinessErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business collaborator created successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
	assert.Nil(resp.Data.(map[string]interface{})["password"])
	assert.Nil(resp.Data.(map[string]interface{})["passwordSalt"])
}

func TestCreateBusinessCollaboratorEmailExists(t *testing.T) {
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

	businessCollaborator := models.BusinessCollaborator{
		FirstName: "New",
		LastName:  "Collaborator",
		Email:     "newcollaborator@gmail.com",
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-collaborators/",
		Body:   &businessCollaborator,
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessCollaborator.BusinessId = businessId
	_, _, _ = simulateJSONRequest(router, createRequest, true)
	w, resp, err := simulateJSONRequest(router, createRequest, true)

	assert.Nil(createBusinessErr)
	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Nil(err)
	assert.Equal("error", resp.Status)
	assert.Equal("Business collaborator email already exists.", resp.Message)
	assert.Nil(resp.Data)
}

func TestGetBusinessCollaboratorSuccess(t *testing.T) {
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

	businessCollaborator := models.BusinessCollaborator{
		FirstName: "New",
		LastName:  "Collaborator",
		Email:     "getcollaborator@gmail.com",
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-collaborators/",
		Body:   &businessCollaborator,
	}

	var getRequest = Request{
		Method: http.MethodGet,
		Path:   "/business-collaborators/",
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessCollaborator.BusinessId = businessId
	_, resp, createBusinessOriginErr := simulateJSONRequest(router, createRequest, true)
	originId := resp.Data.(map[string]interface{})["_id"].(string)
	getRequest.Path += originId
	w, resp, err := simulateJSONRequest(router, getRequest, true)

	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessOriginErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business collaborator found successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
	assert.Nil(resp.Data.(map[string]interface{})["password"])
	assert.Nil(resp.Data.(map[string]interface{})["passwordSalt"])
}

func TestUpdateBusinessCollaboratorSuccess(t *testing.T) {
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

	businessCollaborator := models.BusinessCollaborator{
		FirstName: "New",
		LastName:  "Collaborator",
		Email:     "updatecollaborator@gmail.com",
	}

	updateBusinessCollaborator := models.BusinessCollaborator{
		FirstName: "Update",
		LastName:  "Collaborator Updated",
		Email:     "newupdatedcollaborator@gmail.com",
		Enabled:   false,
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-collaborators/",
		Body:   &businessCollaborator,
	}

	var updateRequest = Request{
		Method: http.MethodPut,
		Path:   "/business-collaborators/",
		Body:   &updateBusinessCollaborator,
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessCollaborator.BusinessId = businessId
	_, resp, createBusinessCollaboratorErr := simulateJSONRequest(router, createRequest, true)
	businessCollaboratorId := resp.Data.(map[string]interface{})["_id"].(string)
	updateRequest.Path += businessCollaboratorId
	w, resp, err := simulateJSONRequest(router, updateRequest, true)

	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessCollaboratorErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business collaborator updated successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
	assert.Nil(resp.Data.(map[string]interface{})["password"])
	assert.Nil(resp.Data.(map[string]interface{})["passwordSalt"])
}

func TestDeleteBusinessCollaboratorSuccess(t *testing.T) {
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

	businessCollaborator := models.BusinessCollaborator{
		FirstName: "New",
		LastName:  "Collaborator",
		Email:     "deletecollaborator@gmail.com",
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-collaborators/",
		Body:   &businessCollaborator,
	}

	var deleteRequest = Request{
		Method: http.MethodDelete,
		Path:   "/business-collaborators/",
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessCollaborator.BusinessId = businessId
	_, resp, createBusinessCollaboratorErr := simulateJSONRequest(router, createRequest, true)
	collaboratorId := resp.Data.(map[string]interface{})["_id"].(string)
	deleteRequest.Path += collaboratorId
	w, resp, err := simulateJSONRequest(router, deleteRequest, true)

	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessCollaboratorErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business collaborator deleted successfully.", resp.Message)
	assert.Nil(resp.Data)
}

func TestListBusinessCollaboratorSuccess(t *testing.T) {
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

	businessCollaborator := models.BusinessCollaborator{
		FirstName: "New",
		LastName:  "Collaborator",
		Email:     "newcollaborator@gmail.com",
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-collaborators/",
		Body:   &businessCollaborator,
	}

	var listRequest = Request{
		Method: http.MethodGet,
		Path:   "/business-collaborators/list/",
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessCollaborator.BusinessId = businessId
	_, resp, createBusinessCollaboratorErr := simulateJSONRequest(router, createRequest, true)
	listRequest.Path += businessId
	w, resp, err := simulateJSONRequest(router, listRequest, true)

	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessCollaboratorErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business collaborator list successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType([]interface{}{}, resp.Data)
	assert.Len(resp.Data, 1)
}

func TestUpdateBusinessCollaboratorPasswordSuccess(t *testing.T) {
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

	businessCollaborator := models.BusinessCollaborator{
		FirstName: "New",
		LastName:  "Collaborator",
		Email:     "newcollaborator@gmail.com",
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-collaborators/",
		Body:   &businessCollaborator,
	}

	var updatePasswordRequest = Request{
		Method: http.MethodPatch,
		Path:   "/business-collaborators/",
		Body: map[string]interface{}{
			"password": "password123",
		},
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessCollaborator.BusinessId = businessId
	_, resp, createBusinessCollaboratorErr := simulateJSONRequest(router, createRequest, true)
	collaboratorId := resp.Data.(map[string]interface{})["_id"].(string)
	updatePasswordRequest.Path += collaboratorId + "/password"
	w, resp, err := simulateJSONRequest(router, updatePasswordRequest, true)

	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessCollaboratorErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business collaborator password updated successfully.", resp.Message)
	assert.Nil(resp.Data)
}

func TestUpdateBusinessCollaboratorPasswordAndLoginSuccess(t *testing.T) {
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

	businessCollaborator := models.BusinessCollaborator{
		FirstName: "New",
		LastName:  "Collaborator",
		Email:     "updatebusinesscollaboratorpassword@gmail.com",
	}

	var createBusinessRequest = Request{
		Method: http.MethodPost,
		Path:   "/businesses/",
		Body:   business,
	}

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-collaborators/",
		Body:   &businessCollaborator,
	}

	var updatePasswordRequest = Request{
		Method: http.MethodPatch,
		Path:   "/business-collaborators/",
		Body: map[string]interface{}{
			"password": "password123",
		},
	}

	var loginRequest = Request{
		Method: http.MethodPost,
		Path:   "/business-collaborators/login",
		Body: map[string]interface{}{
			"email":    "updatebusinesscollaboratorpassword@gmail.com",
			"password": "password123",
		},
	}

	TestSetApiRoutes(t)
	_, resp, createBusinessErr := simulateJSONRequest(router, createBusinessRequest, true)
	businessId := resp.Data.(map[string]interface{})["_id"].(string)
	businessCollaborator.BusinessId = businessId
	_, resp, createBusinessCollaboratorErr := simulateJSONRequest(router, createRequest, true)
	collaboratorId := resp.Data.(map[string]interface{})["_id"].(string)
	updatePasswordRequest.Path += collaboratorId + "/password"
	_, _, _ = simulateJSONRequest(router, updatePasswordRequest, true)
	w, resp, err := simulateJSONRequest(router, loginRequest, true)

	assert.Nil(createBusinessErr)
	assert.Nil(createBusinessCollaboratorErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("Business collaborator logged successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
	assert.NotEmpty(resp.Data)
	assert.NotNil(resp.Data.(map[string]interface{})["user"])
	assert.Contains(resp.Data.(map[string]interface{})["token"], ".")
}
