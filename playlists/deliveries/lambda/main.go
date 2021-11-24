package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/slatermorgan/playlist-builder/pkg/helpers"
	"github.com/slatermorgan/playlist-builder/playlists"
)

type handler struct {
	usecase playlists.PlaylistService
}

func returnInvalidMethod() (helpers.Response, error) {
	return helpers.Fail(
		errors.New("method not supported"),
		http.StatusMethodNotAllowed,
	)
}

// Get a single playlist
func (h *handler) Get(ctx context.Context, id string) (helpers.Response, error) {
	return returnInvalidMethod()
}

// GetAll playlists
func (h *handler) GetAll(ctx context.Context) (helpers.Response, error) {
	return returnInvalidMethod()
}

// Update a single playlist
func (h *handler) Update(ctx context.Context, id string, body []byte) (helpers.Response, error) {
	return returnInvalidMethod()
}

// Create a playlist
func (h *handler) Create(ctx context.Context, body []byte, token string) (helpers.Response, error) {
	playlist := &playlists.CreatePlaylist{}
	if err := json.Unmarshal(body, &playlist); err != nil {
		return helpers.Fail(err, http.StatusInternalServerError)
	}

	createdPlaylist, err := h.usecase.Create(ctx, playlist, token)
	if err != nil {
		return helpers.Fail(err, http.StatusInternalServerError)
	}

	return helpers.Success(createdPlaylist, http.StatusCreated)
}

// Delete a playlist
func (h *handler) Delete(ctx context.Context, id string) (helpers.Response, error) {
	return returnInvalidMethod()
}

func main() {
	usecase, err := playlists.Init(false)
	if err != nil {
		log.Panic(err)
	}

	h := &handler{usecase}
	lambda.Start(helpers.Router(h))
}
