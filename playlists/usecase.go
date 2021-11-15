package playlists

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/slatermorgan/playlist-builder/pkg/spotify"
	"gopkg.in/go-playground/validator.v9"
)

var (
	validate *validator.Validate
)

type Client interface {
	GetArtistsTopTracks(artistID string) ([]*spotify.Track, error)
	SearchForArtist(q string) ([]*spotify.Artist, error)
	AddItemsToPlaylist(updatePlaylist spotify.UpdatePlaylist, userID string) error
	CreatePlaylist(playlistReq *spotify.Playlist, playlistID string) (*spotify.Playlist, error)
}

// Usecase for interacting with playlists
type Usecase struct {
	Client Client
}

// Get a single playlist
func (u *Usecase) Get(ctx context.Context, id string) (*spotify.Playlist, error) {
	return nil, errors.Wrap(nil, "method not supported")
}

// GetAll gets all playlists
func (u *Usecase) GetAll(ctx context.Context) ([]*spotify.Playlist, error) {
	return nil, errors.Wrap(nil, "method not supported")
}

// Update a single playlist
func (u *Usecase) Update(ctx context.Context, id string, playlist *spotify.UpdatePlaylist) error {
	return errors.Wrap(nil, "method not supported")
}

// Create a single playlist
func (u *Usecase) Create(
	ctx context.Context,
	lineupPlaylist *CreatePlaylist,
) (*LineupPlaylist, error) {
	playlistReq := &spotify.Playlist{}
	u.Client.CreatePlaylist(playlistReq, "")
	return &LineupPlaylist{}, nil
}

// Delete a single playlist
func (u *Usecase) Delete(ctx context.Context, id string) error {
	return errors.Wrap(nil, "method not supported")
}

func (u *Usecase) newID() string {
	uid := uuid.New()
	return uid.String()
}
