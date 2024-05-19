package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password int    `json:"password"`
}

func TestGetBody(t *testing.T) {
	// Create an instance of Echo
	e := echo.New()
	// Create a new request with the JSON body
	loginJSON := `{"username":"John Doe","password":12345}`
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(loginJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// Create a new recorder
	rec := httptest.NewRecorder()

	// Create a new echo context
	c := e.NewContext(req, rec)

	// Call the GetBody function
	var login LoginRequest
	resultLogin, err := GetBody[LoginRequest](c, login)
	if err != nil {
		t.Errorf("GetBody returned an unexpected error: %v", err)
	}

	// Assuming LoginRequest is a struct that you have defined
	expectedLogin := LoginRequest{Username: "John Doe", Password: 12345}
	if *resultLogin != expectedLogin {
		t.Errorf("GetBody() = %v, want %v", *resultLogin, expectedLogin)
	}
}
