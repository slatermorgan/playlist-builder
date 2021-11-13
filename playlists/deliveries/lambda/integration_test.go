package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/slatermorgan/go-care/pkg/helpers"
	"github.com/slatermorgan/go-care/shifts"
	"github.com/stretchr/testify/assert"
)

var (
	id           = ""
	validShift   = `{ "name": "Test Shift", "email": "test@test.com", "age": 30 }`
	updatedShift = `{ "name": "Updated Shift", "email": "test@test.com", "age": 30 }`
)

func setup() *handler {
	os.Setenv("TABLE_NAME", "example-shifts-integration")
	usecase, err := shifts.Init(true)
	if err != nil {
		log.Panic(err)
	}

	h := &handler{usecase}
	return h
}

func clear() {
	os.Setenv("TABLE_NAME", "example-shifts-integration")
	usecase, err := shifts.Init(true)
	if err != nil {
		log.Panic(err)
	}

	ctx := context.Background()
	shifts, _ := usecase.GetAll(ctx)
	for _, shift := range shifts {
		go usecase.Delete(ctx, shift.ID)
	}
}

func TestCanCreate(t *testing.T) {
	ctx := context.Background()
	shift := &shifts.Shift{}
	clear()
	h := setup()
	req := helpers.Request{
		HTTPMethod: "POST",
		Body:       validShift,
	}
	res, err := helpers.Router(h)(ctx, req)
	assert.NoError(t, err)

	err = json.Unmarshal([]byte(res.Body), &shift)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotNil(t, shift.ID)
	id = shift.ID
}

func TestCanGetAllShifts(t *testing.T) {
	ctx := context.Background()
	u := []*shifts.Shift{}
	h := setup()
	req := helpers.Request{
		HTTPMethod: "GET",
	}
	res, err := helpers.Router(h)(ctx, req)
	assert.NoError(t, err)
	err = json.Unmarshal([]byte(res.Body), &u)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, "Test Shift", u[0].Name)
}

func TestCanGetShift(t *testing.T) {
	ctx := context.Background()
	u := &shifts.Shift{}
	h := setup()
	req := helpers.Request{
		HTTPMethod: "GET",
		PathParameters: map[string]string{
			"id": id,
		},
	}
	res, err := helpers.Router(h)(ctx, req)
	err = json.Unmarshal([]byte(res.Body), &u)
	assert.NoError(t, err)
	assert.Equal(t, "Test Shift", u.Name)
}

func TestCanUpdateShift(t *testing.T) {
	ctx := context.Background()
	r := map[string]interface{}{}
	h := setup()
	req := helpers.Request{
		HTTPMethod: "PUT",
		PathParameters: map[string]string{
			"id": id,
		},
		Body: updatedShift,
	}
	res, err := helpers.Router(h)(ctx, req)
	err = json.Unmarshal([]byte(res.Body), &r)
	assert.NoError(t, err)
	assert.Equal(t, true, r["success"])
}

func TestCanDeleteShift(t *testing.T) {
	ctx := context.Background()
	r := map[string]interface{}{}
	h := setup()
	req := helpers.Request{
		HTTPMethod: "DELETE",
		PathParameters: map[string]string{
			"id": id,
		},
	}
	res, err := helpers.Router(h)(ctx, req)
	err = json.Unmarshal([]byte(res.Body), &r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, res.StatusCode)
	expected := true
	assert.Equal(t, expected, r["success"])
}
