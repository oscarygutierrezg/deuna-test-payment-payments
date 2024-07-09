package testing

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"payment-payments-api/internal/models"
	"testing"
)

func TestUserLoginDefaultSuccess(t *testing.T) {
	assert := assert.New(t)

	auth := map[string]interface{}{
		"email":    "admin@example.com",
		"password": "admin123",
	}
	var request = Request{
		Method: http.MethodPost,
		Path:   "/users/login",
		Body:   auth,
	}

	TestSetApiRoutes(t)
	w, resp, err := simulateJSONRequest(router, request, true)

	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("User logged successfully", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
	assert.NotEmpty(resp.Data)
	assert.Nil(resp.Data.(map[string]interface{})["user"])
	assert.Contains(resp.Data.(map[string]interface{})["token"], ".")
}

func TestUserLoginDefaultFail(t *testing.T) {
	assert := assert.New(t)

	auth := map[string]interface{}{
		"email":    "admin@example.com",
		"password": "admin1234",
	}
	var request = Request{
		Method: http.MethodPost,
		Path:   "/users/login",
		Body:   auth,
	}

	TestSetApiRoutes(t)
	w, resp, err := simulateJSONRequest(router, request, true)

	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Nil(err)
	assert.Equal("error", resp.Status)
	assert.Equal("User cannot be logged.", resp.Message)
	assert.Nil(resp.Data)
}

func TestCreateUserSuccess(t *testing.T) {
	assert := assert.New(t)

	user := models.User{
		FirstName: "New",
		LastName:  "User",
		Email:     "newuser@gmail.com",
	}
	var request = Request{
		Method: http.MethodPost,
		Path:   "/users/",
		Body:   user,
	}

	TestSetApiRoutes(t)
	w, resp, err := simulateJSONRequest(router, request, true)

	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("User created successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
	assert.Nil(resp.Data.(map[string]interface{})["password"])
	assert.Nil(resp.Data.(map[string]interface{})["passwordSalt"])
}

func TestCreateUserEmailExists(t *testing.T) {
	assert := assert.New(t)

	user := models.User{
		FirstName: "New",
		LastName:  "User",
		Email:     "newuser@gmail.com",
	}
	var request = Request{
		Method: http.MethodPost,
		Path:   "/users/",
		Body:   user,
	}

	TestSetApiRoutes(t)
	_, _, _ = simulateJSONRequest(router, request, true)
	w, resp, err := simulateJSONRequest(router, request, true)

	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Nil(err)
	assert.Equal("error", resp.Status)
	assert.Equal("User email already exists.", resp.Message)
	assert.Nil(resp.Data)
}

func TestGetUserSuccess(t *testing.T) {
	assert := assert.New(t)

	user := models.User{
		FirstName: "New",
		LastName:  "User",
		Email:     "getuser@gmail.com",
	}
	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/users/",
		Body:   user,
	}
	var getRequest = Request{
		Method: http.MethodGet,
		Path:   "/users/",
	}

	TestSetApiRoutes(t)
	_, resp, createErr := simulateJSONRequest(router, createRequest, true)
	userId := resp.Data.(map[string]interface{})["_id"].(string)
	getRequest.Path += userId
	w, resp, err := simulateJSONRequest(router, getRequest, true)

	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("User found successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
	assert.Nil(resp.Data.(map[string]interface{})["password"])
	assert.Nil(resp.Data.(map[string]interface{})["passwordSalt"])
}

func TestGetUserFail(t *testing.T) {
	assert := assert.New(t)

	user := models.User{
		FirstName: "New",
		LastName:  "User",
		Email:     "getuserfail@gmail.com",
	}
	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/users/",
		Body:   user,
	}
	var getRequest = Request{
		Method: http.MethodGet,
		Path:   "/users/",
		Body:   user,
	}

	TestSetApiRoutes(t)
	_, _, createErr := simulateJSONRequest(router, createRequest, true)
	getRequest.Path += uuid.New().String()
	w, resp, err := simulateJSONRequest(router, getRequest, true)

	assert.Nil(createErr)
	assert.Equal(http.StatusBadRequest, w.Code)
	assert.Nil(err)
	assert.Equal("error", resp.Status)
	assert.Equal("User not found.", resp.Message)
	assert.Nil(resp.Data)
}

func TestUpdateUserSuccess(t *testing.T) {
	assert := assert.New(t)

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/users/",
		Body: models.User{
			FirstName: "New",
			LastName:  "User",
			Email:     "updateuser@gmail.com",
		},
	}

	var updateRequest = Request{
		Method: http.MethodPut,
		Path:   "/users/",
		Body: models.User{
			FirstName: "New2",
			LastName:  "User2",
			Email:     "newupdateuser@gmail.com",
			Enabled:   true,
		},
	}

	TestSetApiRoutes(t)
	_, resp, createErr := simulateJSONRequest(router, createRequest, true)
	userId := resp.Data.(map[string]interface{})["_id"].(string)
	updateRequest.Path += userId
	w, resp, err := simulateJSONRequest(router, updateRequest, true)

	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("User updated successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
	assert.Nil(resp.Data.(map[string]interface{})["password"])
	assert.Nil(resp.Data.(map[string]interface{})["passwordSalt"])
}

func TestDeleteUserSuccess(t *testing.T) {
	assert := assert.New(t)

	user := models.User{
		FirstName: "New",
		LastName:  "User",
		Email:     "deleteuser@gmail.com",
	}
	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/users/",
		Body:   user,
	}
	var deleteRequest = Request{
		Method: http.MethodDelete,
		Path:   "/users/",
	}

	TestSetApiRoutes(t)
	_, resp, createErr := simulateJSONRequest(router, createRequest, true)
	userId := resp.Data.(map[string]interface{})["_id"].(string)
	deleteRequest.Path += userId
	w, resp, err := simulateJSONRequest(router, deleteRequest, true)

	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("User deleted successfully.", resp.Message)
	assert.Nil(resp.Data)
}

func TestListUserSuccess(t *testing.T) {
	assert := assert.New(t)

	user := models.User{
		FirstName: "New",
		LastName:  "User",
		Email:     "listuser@gmail.com",
	}
	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/users/",
		Body:   user,
	}
	var listRequest = Request{
		Method: http.MethodGet,
		Path:   "/users/list",
	}

	TestSetApiRoutes(t)
	_, _, createErr := simulateJSONRequest(router, createRequest, true)
	w, resp, err := simulateJSONRequest(router, listRequest, true)

	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("User list successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType([]interface{}{}, resp.Data)
	assert.Len(resp.Data, 1)
}

func TestListOrderDescUserSuccess(t *testing.T) {
	assert := assert.New(t)

	user1 := models.User{
		FirstName: "New1",
		LastName:  "User1",
		Email:     "listuser1@gmail.com",
	}
	user2 := models.User{
		FirstName: "New2",
		LastName:  "User2",
		Email:     "listuser2@gmail.com",
	}
	user3 := models.User{
		FirstName: "New3",
		LastName:  "User3",
		Email:     "listuser3@gmail.com",
	}
	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/users/",
	}
	var listRequest = Request{
		Method: http.MethodGet,
		Path:   "/users/list?skip=1&limit=2&sort=desc",
	}

	TestSetApiRoutes(t)
	createRequest.Body = user1
	_, _, createErr1 := simulateJSONRequest(router, createRequest, true)
	createRequest.Body = user2
	_, _, createErr2 := simulateJSONRequest(router, createRequest, true)
	createRequest.Body = user3
	_, _, createErr3 := simulateJSONRequest(router, createRequest, true)
	w, resp, err := simulateJSONRequest(router, listRequest, true)

	assert.Nil(createErr1)
	assert.Nil(createErr2)
	assert.Nil(createErr3)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("User list successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType([]interface{}{}, resp.Data)
	assert.Len(resp.Data, 2)
	assert.Equal(user2.FirstName, resp.Data.([]interface{})[0].(map[string]interface{})["firstName"])
	assert.Equal(user1.FirstName, resp.Data.([]interface{})[1].(map[string]interface{})["firstName"])
}

func TestListOrderDescFirstNameUserSuccess(t *testing.T) {
	assert := assert.New(t)

	user1 := models.User{
		FirstName: "New1",
		LastName:  "User1",
		Email:     "listuser1@gmail.com",
	}
	user2 := models.User{
		FirstName: "New2",
		LastName:  "User2",
		Email:     "listuser2@gmail.com",
	}
	user3 := models.User{
		FirstName: "New3",
		LastName:  "User3",
		Email:     "listuser3@gmail.com",
	}
	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/users/",
	}
	var listRequest = Request{
		Method: http.MethodGet,
		Path:   "/users/list?limit=2&sort=desc&by=firstName",
	}

	TestSetApiRoutes(t)
	createRequest.Body = user1
	_, _, createErr1 := simulateJSONRequest(router, createRequest, true)
	createRequest.Body = user2
	_, _, createErr2 := simulateJSONRequest(router, createRequest, true)
	createRequest.Body = user3
	_, _, createErr3 := simulateJSONRequest(router, createRequest, true)
	w, resp, err := simulateJSONRequest(router, listRequest, true)

	assert.Nil(createErr1)
	assert.Nil(createErr2)
	assert.Nil(createErr3)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("User list successfully.", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType([]interface{}{}, resp.Data)
	assert.Len(resp.Data, 2)
	assert.Equal(user3.FirstName, resp.Data.([]interface{})[0].(map[string]interface{})["firstName"])
	assert.Equal(user2.FirstName, resp.Data.([]interface{})[1].(map[string]interface{})["firstName"])
}

func TestUpdateUserPasswordSuccess(t *testing.T) {
	assert := assert.New(t)

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/users/",
		Body: models.User{
			FirstName: "New",
			LastName:  "User",
			Email:     "updateuserpassword@gmail.com",
		},
	}

	var updatePasswordRequest = Request{
		Method: http.MethodPatch,
		Path:   "/users/",
		Body: map[string]interface{}{
			"password": "password123",
		},
	}

	TestSetApiRoutes(t)
	_, resp, createErr := simulateJSONRequest(router, createRequest, true)
	userId := resp.Data.(map[string]interface{})["_id"].(string)
	updatePasswordRequest.Path += userId + "/password"
	w, resp, err := simulateJSONRequest(router, updatePasswordRequest, true)

	assert.Nil(createErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("User password updated successfully.", resp.Message)
	assert.Nil(resp.Data)
}

func TestUpdateUserPasswordAndLoginSuccess(t *testing.T) {
	assert := assert.New(t)

	var createRequest = Request{
		Method: http.MethodPost,
		Path:   "/users/",
		Body: models.User{
			FirstName: "New",
			LastName:  "User",
			Email:     "updateuserpassword@gmail.com",
		},
	}

	var updatePasswordRequest = Request{
		Method: http.MethodPatch,
		Path:   "/users/",
		Body: map[string]interface{}{
			"password": "password123",
		},
	}

	var loginRequest = Request{
		Method: http.MethodPost,
		Path:   "/users/login",
		Body: map[string]interface{}{
			"email":    "updateuserpassword@gmail.com",
			"password": "password123",
		},
	}

	TestSetApiRoutes(t)
	_, resp, createErr := simulateJSONRequest(router, createRequest, true)
	userId := resp.Data.(map[string]interface{})["_id"].(string)
	updatePasswordRequest.Path += userId + "/password"
	_, resp, updatePasswordErr := simulateJSONRequest(router, updatePasswordRequest, true)
	w, resp, err := simulateJSONRequest(router, loginRequest, true)

	assert.Nil(createErr)
	assert.Nil(updatePasswordErr)
	assert.Equal(http.StatusOK, w.Code)
	assert.Nil(err)
	assert.Equal("success", resp.Status)
	assert.Equal("User logged successfully", resp.Message)
	assert.NotNil(resp.Data)
	assert.IsType(map[string]interface{}{}, resp.Data)
	assert.NotEmpty(resp.Data)
	assert.NotNil(resp.Data.(map[string]interface{})["user"])
	assert.Contains(resp.Data.(map[string]interface{})["token"], ".")
}
