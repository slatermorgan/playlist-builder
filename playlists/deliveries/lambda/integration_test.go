package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/slatermorgan/playlist-builder/pkg/helpers"
	"github.com/slatermorgan/playlist-builder/playlists"
	"github.com/stretchr/testify/assert"
)

var (
	id              = ""
	validPlaylist   = `{ "name": "Test Playlist", "email": "test@test.com", "age": 30 }`
	updatedPlaylist = `{ "name": "Updated Playlist", "email": "test@test.com", "age": 30 }`
)

func setup() *handler {
	os.Setenv("TABLE_NAME", "example-playlists-integration")
	usecase, err := playlists.Init(true)
	if err != nil {
		log.Panic(err)
	}

	h := &handler{usecase}
	return h
}

func clear() {
	os.Setenv("TABLE_NAME", "example-playlists-integration")
	usecase, err := playlists.Init(true)
	if err != nil {
		log.Panic(err)
	}

	ctx := context.Background()
	playlists, _ := usecase.GetAll(ctx)
	for _, playlist := range playlists {
		go usecase.Delete(ctx, playlist.ID)
	}
}

func TestCanCreate(t *testing.T) {
	ctx := context.Background()
	playlist := &playlists.Playlist{}
	clear()
	h := setup()
	req := helpers.Request{
		HTTPMethod: "POST",
		Body:       validPlaylist,
	}
	res, err := helpers.Router(h)(ctx, req)
	assert.NoError(t, err)

	err = json.Unmarshal([]byte(res.Body), &playlist)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotNil(t, playlist.ID)
	id = playlist.ID
}

func TestCanGetAllPlaylists(t *testing.T) {
	ctx := context.Background()
	u := []*playlists.Playlist{}
	h := setup()
	req := helpers.Request{
		HTTPMethod: "GET",
	}
	res, err := helpers.Router(h)(ctx, req)
	assert.NoError(t, err)
	err = json.Unmarshal([]byte(res.Body), &u)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, "Test Playlist", u[0].Name)
}

func TestCanGetPlaylist(t *testing.T) {
	ctx := context.Background()
	u := &playlists.Playlist{}
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
	assert.Equal(t, "Test Playlist", u.Name)
}

func TestCanUpdatePlaylist(t *testing.T) {
	ctx := context.Background()
	r := map[string]interface{}{}
	h := setup()
	req := helpers.Request{
		HTTPMethod: "PUT",
		PathParameters: map[string]string{
			"id": id,
		},
		Body: updatedPlaylist,
	}
	res, err := helpers.Router(h)(ctx, req)
	err = json.Unmarshal([]byte(res.Body), &r)
	assert.NoError(t, err)
	assert.Equal(t, true, r["success"])
}

func TestCanDeletePlaylist(t *testing.T) {
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
