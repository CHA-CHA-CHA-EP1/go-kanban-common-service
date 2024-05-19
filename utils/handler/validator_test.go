package handler

import (
	"testing"
)

func TestValidatePass(t *testing.T) {
	login := loginRequest {
		Username:"John Doe", 
		Password: 123456,
	}

	err := Validate(login)
	if err != nil {
		t.Errorf("Validate() returned an unexpected error: %v", err)
	}
}

func TestValidateFail(t *testing.T) {
	login := loginRequest {
		Username:"John Doe",
	}

	err := Validate(login)

	if err == nil {
		t.Errorf("Validate() returned an unexpected error: %v", err)
	}
}
