package playlists

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

var (
	validate *validator.Validate
)

type sdk interface {
	GetArtistsTopTrack(ctx context.Context, id string) (*Playlist, error)
	CreatePlaylist(ctx context.Context) ([]*Playlist, error)
	AddTrackToPlaylist(ctx context.Context, id string, playlist *UpdatePlaylist) error
}

// Usecase for interacting with playlists
type Usecase struct {
	SDK sdk
}

// Get a single playlist
func (u *Usecase) Get(ctx context.Context, id string) (*Playlist, error) {
	return nil, errors.Wrap(nil, "method not supported")
}

// GetAll gets all playlists
func (u *Usecase) GetAll(ctx context.Context) ([]*Playlist, error) {
	return nil, errors.Wrap(nil, "method not supported")
}

// Update a single playlist
func (u *Usecase) Update(ctx context.Context, id string, playlist *UpdatePlaylist) error {
	return errors.Wrap(nil, "method not supported")
}

// Create a single playlist
func (u *Usecase) Create(ctx context.Context, playlist *Playlist) (*Playlist, error) {
	validate = validator.New()
	if err := validate.Struct(*playlist); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		// Here we should create custom returns
		return nil, validationErrors
	}

	playlist.ID = u.newID()
	if playlist, err := u.SDK.CreatePlaylist(playlist)(ctx, playlist); err != nil {
		return nil, errors.Wrap(err, "error creating new playlist")
	}

	return playlist, nil
}

// Delete a single playlist
func (u *Usecase) Delete(ctx context.Context, id string) error {
	return errors.Wrap(nil, "method not supported")
}

func (u *Usecase) newID() string {
	uid := uuid.New()
	return uid.String()
}
