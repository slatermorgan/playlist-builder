package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/slatermorgan/playlist-builder/pkg/helpers"
	"github.com/slatermorgan/playlist-builder/pkg/spotify"
	"github.com/slatermorgan/playlist-builder/playlists"
)

type handler struct {
	usecase playlists.PlaylistService
}

// Get a single playlist
func (h *handler) Get(ctx context.Context, id string) (helpers.Response, error) {
	playlist, err := h.usecase.Get(ctx, id)
	if err != nil {
		return helpers.Fail(err, http.StatusInternalServerError)
	}

	return helpers.Success(playlist, http.StatusOK)
}

// GetAll playlists
func (h *handler) GetAll(ctx context.Context) (helpers.Response, error) {
	playlists, err := h.usecase.GetAll(ctx)
	if err != nil {
		return helpers.Fail(err, http.StatusInternalServerError)
	}

	return helpers.Success(playlists, http.StatusOK)
}

// Update a single playlist
func (h *handler) Update(ctx context.Context, id string, body []byte) (helpers.Response, error) {
	updatePlaylist := &spotify.UpdatePlaylist{}
	if err := json.Unmarshal(body, &updatePlaylist); err != nil {
		return helpers.Fail(err, http.StatusInternalServerError)
	}

	if err := h.usecase.Update(ctx, id, updatePlaylist); err != nil {
		return helpers.Fail(err, http.StatusInternalServerError)
	}

	return helpers.Success(map[string]interface{}{
		"success": true,
	}, http.StatusNoContent)
}

// Create a playlist
func (h *handler) Create(ctx context.Context, body []byte) (helpers.Response, error) {
	playlist := &spotify.Playlist{}
	if err := json.Unmarshal(body, &playlist); err != nil {
		return helpers.Fail(err, http.StatusInternalServerError)
	}

	createdPlaylist, err := h.usecase.Create(ctx, playlist)
	if err != nil {
		return helpers.Fail(err, http.StatusInternalServerError)
	}

	return helpers.Success(createdPlaylist, http.StatusCreated)
}

// Delete a playlist
func (h *handler) Delete(ctx context.Context, id string) (helpers.Response, error) {
	if err := h.usecase.Delete(ctx, id); err != nil {
		return helpers.Fail(err, http.StatusInternalServerError)
	}

	return helpers.Success(map[string]interface{}{
		"success": true,
	}, http.StatusNoContent)
}

func main() {
	usecase, err := playlists.Init(false)
	if err != nil {
		log.Panic(err)
	}

	h := &handler{usecase}
	lambda.Start(helpers.Router(h))
}
