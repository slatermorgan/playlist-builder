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

type repository interface {
	Get(ctx context.Context, id string) (*Playlist, error)
	GetAll(ctx context.Context) ([]*Playlist, error)
	Update(ctx context.Context, id string, playlist *UpdatePlaylist) error
	Create(ctx context.Context, playlist *Playlist) error
	Delete(ctx context.Context, id string) error
}

// Usecase for interacting with playlists
type Usecase struct {
	Repository repository
}

// Get a single playlist
func (u *Usecase) Get(ctx context.Context, id string) (*Playlist, error) {
	playlist, err := u.Repository.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "error fetching a single playlist")
	}
	return playlist, nil
}

// GetAll gets all playlists
func (u *Usecase) GetAll(ctx context.Context) ([]*Playlist, error) {
	playlists, err := u.Repository.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error fetching all playlists")
	}
	return playlists, nil
}

// Update a single playlist
func (u *Usecase) Update(ctx context.Context, id string, playlist *UpdatePlaylist) error {
	validate = validator.New()
	if err := validate.Struct(playlist); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}

	if err := u.Repository.Update(ctx, id, playlist); err != nil {
		return errors.Wrap(err, "error updating playlist")
	}
	return nil
}

// Create a single playlist
func (u *Usecase) Create(ctx context.Context, playlist *Playlist) error {
	validate = validator.New()
	if err := validate.Struct(*playlist); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		// Here we should create custom returns
		return validationErrors
	}

	playlist.ID = u.newID()
	if err := u.Repository.Create(ctx, playlist); err != nil {
		return errors.Wrap(err, "error creating new playlist")
	}

	return nil
}

// Delete a single playlist
func (u *Usecase) Delete(ctx context.Context, id string) error {
	if err := u.Repository.Delete(ctx, id); err != nil {
		return errors.Wrap(err, "error deleting playlist")
	}
	return nil
}

func (u *Usecase) newID() string {
	uid := uuid.New()
	return uid.String()
}
