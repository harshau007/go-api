package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/harshau007/go-api/controller"
	"github.com/harshau007/go-api/models"
)


var respUser models.User

func TestCreateUser(t *testing.T) {
	requestBody := `{"first":"John","last":"Smith"}`

	req := httptest.NewRequest("POST", "/users", strings.NewReader(requestBody))
	rec := httptest.NewRecorder()

	controller.CreateUser(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK, got %v", res.StatusCode)
	}

	json.NewDecoder(res.Body).Decode(&respUser)

	if respUser.First != "John" {
		t.Error("expected user name to be 'John'")
	}
}

func TestUpdateUser(t *testing.T) {
	id := respUser.ID.String()
	updateBody := `{"first":"John","last":"Oliver"}`

	req := httptest.NewRequest("PUT", "/user/"+id, strings.NewReader(updateBody))
	rec := httptest.NewRecorder()

	controller.UpdateUser(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK, got %v", res.StatusCode)
	}

	json.NewDecoder(res.Body).Decode(&respUser)
	if respUser.Last != "Oliver" {
		t.Error("expected last name to be 'Oliver'")
	}
}

func TestGetAllUsers(t *testing.T) {
	req := httptest.NewRequest("GET", "/users", nil)
	rec := httptest.NewRecorder()

	controller.GetAllUser(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK, got %v", res.StatusCode)
	}

	var respUsers []models.User
	json.NewDecoder(res.Body).Decode(&respUsers)

	if len(respUsers) != 1 {
		t.Errorf("expected 1 users, got %d", len(respUsers))
	}
}

func TestDeleteUser(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/users/1", nil)
	rec := httptest.NewRecorder()

	controller.DeleteAllUser(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK, got %v", res.StatusCode)
	}

	var id int
	json.NewDecoder(res.Body).Decode(&id)

	if id != 1 {
		t.Errorf("expected deleted user id 1, got %d", id)
	}
}
