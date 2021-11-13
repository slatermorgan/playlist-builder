package main

import (
	"context"
	"encoding/json"
	"go-care/shifts"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/slatermorgan/go-care/pkg/helpers"
	"github.com/slatermorgan/go-care/playlists"
)

type handler struct {
	usecase playlists.ShiftService
}

// Get a single shift
func (h *handler) Get(ctx context.Context, id string) (helpers.Response, error) {
	shift, err := h.usecase.Get(ctx, id)
	if err != nil {
		return helpers.Fail(err, http.StatusInternalServerError)
	}

	return helpers.Success(shift, http.StatusOK)
}

// GetAll shifts
func (h *handler) GetAll(ctx context.Context) (helpers.Response, error) {
	shifts, err := h.usecase.GetAll(ctx)
	if err != nil {
		return helpers.Fail(err, http.StatusInternalServerError)
	}

	return helpers.Success(shifts, http.StatusOK)
}

// Update a single shift
func (h *handler) Update(ctx context.Context, id string, body []byte) (helpers.Response, error) {
	updateShift := &shifts.UpdateShift{}
	if err := json.Unmarshal(body, &updateShift); err != nil {
		return helpers.Fail(err, http.StatusInternalServerError)
	}

	if err := h.usecase.Update(ctx, id, updateShift); err != nil {
		return helpers.Fail(err, http.StatusInternalServerError)
	}

	return helpers.Success(map[string]interface{}{
		"success": true,
	}, http.StatusNoContent)
}

// Create a shift
func (h *handler) Create(ctx context.Context, body []byte) (helpers.Response, error) {
	shift := &shifts.Shift{}
	if err := json.Unmarshal(body, &shift); err != nil {
		return helpers.Fail(err, http.StatusInternalServerError)
	}

	if err := h.usecase.Create(ctx, shift); err != nil {
		return helpers.Fail(err, http.StatusInternalServerError)
	}

	return helpers.Success(shift, http.StatusCreated)
}

// Delete a shift
func (h *handler) Delete(ctx context.Context, id string) (helpers.Response, error) {
	if err := h.usecase.Delete(ctx, id); err != nil {
		return helpers.Fail(err, http.StatusInternalServerError)
	}

	return helpers.Success(map[string]interface{}{
		"success": true,
	}, http.StatusNoContent)
}

func main() {
	usecase, err := shifts.Init(false)
	if err != nil {
		log.Panic(err)
	}

	h := &handler{usecase}
	lambda.Start(helpers.Router(h))
}
