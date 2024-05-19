package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

type loginRequest struct {
	Username string `json:"username" validate:"required"`
	Password int    `json:"password" validate:"required"`
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
	var login loginRequest
	resultLogin, err := GetBody(c, login)
	if err != nil {
		t.Errorf("GetBody returned an unexpected error: %v", err)
	}

	// Assuming LoginRequest is a struct that you have defined
	expectedLogin := loginRequest{Username: "John Doe", Password: 12345}
	if *resultLogin != expectedLogin {
		t.Errorf("GetBody() = %v, want %v", *resultLogin, expectedLogin)
	}
}


func TestGetBodyFailed(t *testing.T) {
	// Create an instance of Echo
	e := echo.New()
	// Create a new request with the JSON body
	loginJSON := `{"username":"John Doe","password":"12345"}`
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(loginJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// Create a new recorder
	rec := httptest.NewRecorder()

	// Create a new echo context
	c := e.NewContext(req, rec)

	// Call the GetBody function
	var login loginRequest
	_, err := GetBody(c, login)

	if err == nil {
		t.Errorf("GetBody returned an unexpected error: %v", err)
	}
}


func TestGetBodyFailedMissingField(t *testing.T) {
	// Create an instance of Echo
	e := echo.New()
	// Create a new request with the JSON body
	loginJSON := `{"username":"John Doe"}`
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(loginJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// Create a new recorder
	rec := httptest.NewRecorder()

	// Create a new echo context
	c := e.NewContext(req, rec)

	// Call the GetBody function
	var login loginRequest

	_, err := GetBody(c, login)

	if err != nil {
		t.Errorf("GetBody returned an unexpected error: %v", err)
	}
}
