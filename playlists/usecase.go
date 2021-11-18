package playlists

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/slatermorgan/playlist-builder/pkg/spotify"
	"gopkg.in/go-playground/validator.v9"
)

var (
	validate *validator.Validate
)

const tracksPerArtist = 5

type Client interface {
	GetArtistsTopTracks(artistID string) ([]*spotify.Track, error)
	SearchForArtist(q string) ([]*spotify.Artist, error)
	AddItemsToPlaylist(updatePlaylist spotify.UpdatePlaylist, userID string) error
	CreatePlaylist(playlistReq *spotify.Playlist, userID string) (*spotify.Playlist, error)
	GetPlaylist(playlistID string) (*spotify.Playlist, error)
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
	createPlaylist *CreatePlaylist,
) (*spotify.Playlist, error) {
	validate = validator.New()
	if err := validate.Struct(*createPlaylist); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		// Here we should create custom returns
		return nil, validationErrors
	}

	newPlaylist, err := u.Client.CreatePlaylist(
		&spotify.Playlist{
			Name:        createPlaylist.Name,
			IsPublic:    createPlaylist.IsPublic,
			Description: createPlaylist.Description,
		},
		"userID",
	)
	if err != nil {
		return nil, errors.Wrap(nil, "failed to create playlist")
	}

	var playlistTrackURIs string
	for _, name := range createPlaylist.ArtistNames {
		artists, err := u.Client.SearchForArtist(name)
		if err != nil {
			return nil, errors.Wrap(nil, "failed to search artist: "+name)
		}

		for _, artist := range artists {
			if artist.Name == name {
				topTracks, err := u.Client.GetArtistsTopTracks(artist.ID)
				if err != nil {
					return nil, errors.Wrap(nil, "failed to search artist: "+name)
				}

				for i := 0; i < tracksPerArtist; i++ {
					playlistTrackURIs += topTracks[i].URI + ","
				}
			}
		}
		playlistTrackURIs = strings.TrimRight(playlistTrackURIs, ",")
	}

	err = u.Client.AddItemsToPlaylist(
		spotify.UpdatePlaylist{
			URIs: playlistTrackURIs,
		},
		newPlaylist.ID,
	)
	if err != nil {
		return nil, errors.Wrap(nil, "failed to add items to playlist ID: "+newPlaylist.ID)
	}

	newPlaylist, err = u.Client.GetPlaylist(newPlaylist.ID)
	if err != nil {
		return nil, errors.Wrap(nil, "failed to get playlist ID: "+newPlaylist.ID)
	}

	return newPlaylist, nil
}

// Delete a single playlist
func (u *Usecase) Delete(ctx context.Context, id string) error {
	return errors.Wrap(nil, "method not supported")
}

func (u *Usecase) newID() string {
	uid := uuid.New()
	return uid.String()
}
